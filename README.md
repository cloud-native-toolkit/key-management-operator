# Key Management operator

## Prerequisites

- Operator SDK v1.0.1

## Development

### Run the tests

```shell script
make test
```

### Run the operator

```shell script
make install
make deploy IMG=<some-registry>/<project-name>:<tag>
```

## Initial setup

### Initialize the operator

```shell script
operator-sdk init --domain=ibm --repo=github.com/ibmgaragecloud/key-management-operator
```

### Generate the skeleton CRD and controller

```shell script
operator-sdk create api --group keymanagement --version v1 --kind SecretTemplate --resource=true --controller=true
```
