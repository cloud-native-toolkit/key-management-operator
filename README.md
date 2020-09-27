# Key Management operator

## Prerequisites

- Operator SDK v1.0.1

## Initial setup

### Initialize the operator

```shell script
operator-sdk init --domain=ibm --repo=github.com/ibmgaragecloud/key-management-operator
```

### Generate the skeleton CRD and controller

```shell script
operator-sdk create api --group keymanagement --version v1 --kind SecretTemplate --resource=true --controller=true
```
