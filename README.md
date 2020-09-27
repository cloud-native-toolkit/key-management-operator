# Key Management operator

## Prerequisites

- Operator SDK v1.0.1

## Development

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
