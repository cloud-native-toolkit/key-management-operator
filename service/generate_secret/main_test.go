package generate_secret

import (
	"encoding/base64"
	"github.com/ibm-garage-cloud/key-management-operator/pkg/util/test_support"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

type MockKeyManager struct {
	id string
	value string
}

func (t MockKeyManager) GetKey(keyId string) string {
	return t.value
}

func (t MockKeyManager) PopulateMetadata(annotations *map[string]string) {
}

func (t MockKeyManager) Id() string {
	return t.id
}

type MySecret struct {
	metav1.TypeMeta   `json:",inline" yaml:"inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Data         map[string]string   `json:"data,omitempty" yaml:"data,omitempty"`
	StringData   map[string]string   `json:"stringdata,omitempty" yaml:"stringdata,omitempty"`
}

func Test_newSecret(t *testing.T) {

	name := "secret"
	value := "value1"
	expectedValue := "dmFsdWUx"

	metadata := metav1.ObjectMeta{Name: name}
	var data map[string][]byte
	data = make(map[string][]byte)

	data["test"] = []byte(base64.StdEncoding.EncodeToString([]byte(value)))

	secret := newSecret(&metadata, &data)

	test_support.ExpectEqual(t, name, secret.ObjectMeta.Name)
	test_support.ExpectEqual(t, expectedValue, string(secret.Data["test"]))
}

//func Test_generateSecret(t *testing.T) {
//
//	value := "value"
//	keyValue := "dmFsdWU="
//	b64encodedvalue := "YnY2NHZhbHVl="
//	b64value := "b64value"
//
//	var km key_management.KeyManager
//	km = MockKeyManager{id: "id", value: keyValue}
//
//	var values []v1.SecretTemplateValue
//	values = []v1.SecretTemplateValue{
//		v1.SecretTemplateValue{Name: "name1", KeyId: "id"},
//		v1.SecretTemplateValue{Name: "name2", Value: value},
//		v1.SecretTemplateValue{Name: "name3", B64Value: b64encodedvalue},
//	}
//
//	secretTemplate := v1.SecretTemplate{
//		ObjectMeta: metav1.ObjectMeta{Name: "mysecret"},
//		Spec: v1.SecretTemplateSpec{
//			Values: values,
//		},
//	}
//
//	actualValue := *generateSecret(&km, &secretTemplate)
//
//	test_support.ExpectEqual(t, value, string(actualValue.Data["name1"]))
//	test_support.ExpectEqual(t, value, string(actualValue.Data["name2"]))
//	test_support.ExpectEqual(t, b64value, string(actualValue.Data["name3"]))
//}

//func Test_generateSecret_marshal(t *testing.T) {
//
//	value := "value"
//	keyValue := "dmFsdWU="
//	b64value := "YnY2NHZhbHVl="
//
//	var km key_management.KeyManager
//	km = MockKeyManager{id: "id", value: keyValue}
//
//	var values []v1.SecretTemplateValue
//	values = []v1.SecretTemplateValue{
//		v1.SecretTemplateValue{Name: "name1", KeyId: "id"},
//		v1.SecretTemplateValue{Name: "name2", Value: value},
//		v1.SecretTemplateValue{Name: "name3", B64Value: b64value},
//	}
//
//	secretTemplate := v1.SecretTemplate{
//		ObjectMeta: metav1.ObjectMeta{Name: "mysecret"},
//		Spec: v1.SecretTemplateSpec{
//			Values: values,
//		},
//	}
//
//	secret := *generateSecret(&km, &secretTemplate)
//
//	jsonSecret, err := json.Marshal(&secret)
//	if err != nil {
//		panic(err)
//	}
//
//	yamlSecret, err := yaml2.JSONToYAML(jsonSecret)
//
//	actualValue := MySecret{}
//
//	err = yaml.Unmarshal(yamlSecret, &actualValue)
//	if err != nil{
//		t.Error(err)
//	}
//
//	test_support.ExpectEqualBase64(t, keyValue, string(actualValue.Data["name1"]))
//	test_support.ExpectEqualBase64(t, "dmFsdWU=", string(actualValue.Data["name2"]))
//	test_support.ExpectEqualBase64(t, b64value, string(actualValue.Data["name3"]))
//}
