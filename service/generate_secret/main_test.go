package generate_secret

import (
	"encoding/base64"
	keymanagementv1 "github.com/ibm-garage-cloud/key-management-operator/api/v1"
	"github.com/ibm-garage-cloud/key-management-operator/service/key_management"
	"github.com/ibm-garage-cloud/key-management-operator/util/test_support"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

type MockKeyManager struct {
	id    string
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

	Data       map[string]string `json:"data,omitempty" yaml:"data,omitempty"`
	StringData map[string]string `json:"stringdata,omitempty" yaml:"stringdata,omitempty"`
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

func Test_processBase64StringValue(t *testing.T) {

	b64String := "VGhpcyBpcyBhIHRlc3Qgc3RyaW5n"

	result := processBase64StringValue(&b64String)

	test_support.ExpectEqual(t, "This is a test string", string(*result))
}

type TestKeyManager struct {
	Value *string
}

func (k TestKeyManager) GetKey(keyId *string) *string {
	return k.Value
}

func (k TestKeyManager) PopulateMetadata(annotations *map[string]string) {

}

func (k TestKeyManager) Id() string {
	return "id"
}

func Test_convertValue_keyId(t *testing.T) {

	b64String := "VGhpcyBpcyBhIHRlc3Qgc3RyaW5n"

	var km key_management.KeyManager
	km = TestKeyManager{Value: &b64String}

	keyValue := keymanagementv1.SecretTemplateValue{Key: "Key", KeyId: "keyid"}

	annotations := make(map[string]string)

	result := convertValue(&km, &keyValue, &annotations)

	test_support.ExpectEqual(t, "This is a test string", string(*result))
}
