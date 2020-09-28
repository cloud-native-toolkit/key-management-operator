# Key Management operator

Operator that reads secret configuration in a SecretTemplate custom resource, looks up secret values from 
a key manager, and generates a Secret.

## Why?

A common issue when doing GitOps is dealing with sensitive information that should not be stored in the
git repository (e.g. passwords, keys, etc). There are two different approaches to how to handle this issue:

1. Inject the values from another source into kubernetes Secret(s) at deployment time
2. Inject the values from another source in the pod at startup time via an InitContainer

The "other source" in this case would be a key management system that centralizes the storage and management
of sensitive information.

This operator addresses the first approach listed above by pulling sensitive values from **Key Protect** at 
deployment time to generate the appropriate kubernetes Secret(s).

## How it works

The operator takes a SecretTemplate custom resource(s) as input, looks up the values of any sensitive information
for the provided keyIds from Key Protect, and generates a kubernetes Secret for each input template.

### Secret template

The input to the operator is one or more "secret templates". In this case the "secret template" provides the 
structure of the desired template with placeholders for the values that will be pulled from the key management system. 
The following provides the structure of the template:

```yaml
apiVersion: keymanagement.ibm/v1
kind: SecretTemplate
metadata:
  name: mysecret
  annotations:
    key-manager: key-protect
    key-protect/instanceId: instance-id
    key-protect/region: us-east
spec:
  labels: {}
  annotations: {}
  values:
    - name: url
      value: https://ibm.com
    - name: username
      b64value: dGVhbS1jYXA=
    - name: password
      keyId: 36397b07-d98d-4c0b-bd7a-d6c290163684
``` 

- The `metadata.annotations` value is optional. 

    - `key-manager` - the only value supported currently is `key-protect`
    - `key-protect/instanceId` - the instance id of the key protect instance. If not provided then the `instance-id` value from the `key-protect-access` secret will be used.
    - `key-protect/region` - the region where the key protect instance has been provisioned. If not provided then the `region` value from the `key-protect-access` secret will be used.
    
- The `metadata.name` value given will be used as the name for the Secret that will be generated.
- The information in `spec.labels` and `spec.annotations` will be copied over as the `labels` and `annotations` in the Secret that is generated
- The `spec.values` section contains the information that should be provided in the `data` section of the generated Secret. There are three prossible ways the values can be provided:

    - `value` - the actual value can be provided directly as clear text. This would be appropriate for information that is not sensitive but is required in the secret
    - `b64value` - a base64 encoded value can be provided to the secret. This can be used for large values that might present formatting issues or for information that is not sensitive but that might be obfuscated a bit (like a username)
    - `keyId` - the id (not the name) of the Standard Key that has been stored in Key Protect. The value stored in Key Protect can be anything

### Managing keys in Key Protect

Key Protect manages two different types of keys: `root keys` and `standard keys`. `Standard keys` are used to store any
kind of protected information. The Key Protect plugin reads the contents of a standard key, identified by a given key id, and
stores the key value into a secret in the cluster.

The following steps describe how to create a standard key:

1. Open the IBM Cloud console and navigate to the Key Protect service

2. Within Key Protect, select the **Manage Keys** tab

3. Press the `Add key` button to open the "Add a new key" dialog

4. Select the `Import your own key` radio button and `Standard key` from the drop down

5. Provide a descriptive name for the key and paste the base-64 encoded value of the key into the `Key material` field

    **Note:** A value can be encoded as base-64 from the terminal with the following command:
    
    ```shell script
    echo -n "{VALUE}" | base64
    ```
   
    If you need to encode a larger value, create the value in a file and encode the entire contents of the file with:
    
    ```shell script
    cat {file} | base64
    ```

6. Click **Import key** to create the key

7. Copy the value of the **ID**. This will be used later by the plugin

## Setting up the operator

### Key Protect credentials

In order to connect with Key Protect you will need three pieces of information:

- `IBM Cloud API Key` - an API Key that has `Reader` and `ReaderPlus` access to the Key Protect instance
- `Key Protect region` - the region where the Key Protect instance has been deployed
- `Key Protect instance id` - the GUID of the Key Protect instance where the secrets are stored

The three values can be provided in a secret named `key-protect-access` in the same namespace where ArgoCD has been 
deployed. As shown above, the region and instance id values can be provided in the SecretTemplate configuration. Additionally, 
the default API Key will be used from the `cloud-access` secret if the value is not available from the `key-protect-access` secret.

#### Get the instance id for Key Protect

1. Set the target resource group and region for the Key Protect instance.
    
    ```shell script
    ibmcloud target -g {RESOURCE_GROUP} -r {REGION}
    ```
  
2. List the available resources and find the name of the Key Protect instance.

    ```shell script
    ibmcloud resource service-instances
    ```
   
3. List the details for the Key Protect instance. The `Key Protect instance id` is listed as `GUID`.

    ```shell script
    ibmcloud resource service-instance {INSTANCE_NAME} 
    ```

#### Create the secret with the information needed to access Key Protect

Armed with the information from the previous step, run the following to create a secret:

```shell script
NAMESPACE="tools"
kubectl create secret generic -n ${NAMESPACE} key-protect-access \
  --from-literal=api-key=${API_KEY} \ 
  --from-literal=region=${REGION} \
  --from-literal=instance-id=${KP_INSTANCE_ID}
```

where:
- `NAMESPACE` should be the namespace where ArgoCD has been deployed

## Development

### Prerequisites

- Operator SDK v1.0.1

### Run the tests

```shell script
make test
```

### Build and push the image

```shell script
make container-build container-push
```

### Run the operator

```shell script
make install
make deploy
```

## Initial setup

These are the steps that were performed to initialize the operator.

### Initialize the operator

```shell script
operator-sdk init --domain=ibm --repo=github.com/ibmgaragecloud/key-management-operator
```

### Generate the skeleton CRD and controller

```shell script
operator-sdk create api --group keymanagement --version v1 --kind SecretTemplate --resource=true --controller=true
```
