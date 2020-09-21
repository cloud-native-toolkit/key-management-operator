module github.com/ibm-garage-cloud/key-management-operator

go 1.13

require (
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32
	github.com/imdario/mergo v0.3.7
	github.com/operator-framework/operator-sdk v0.17.2
	github.com/spf13/pflag v1.0.5
	gopkg.in/yaml.v2 v2.2.8
	k8s.io/api v0.17.4
	k8s.io/apimachinery v0.17.4
	k8s.io/client-go v12.0.0+incompatible
	sigs.k8s.io/controller-runtime v0.5.2
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible // Required by OLM
	github.com/ibm-garage-cloud/key-management-operator => ./
	k8s.io/client-go => k8s.io/client-go v0.17.4 // Required by prometheus-operator
)
