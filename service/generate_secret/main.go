package generate_secret

import (
	"encoding/base64"
	"fmt"
	keymanagementv1 "github.com/ibmgaragecloud/key-management-operator/api/v1"
	"github.com/ibmgaragecloud/key-management-operator/service/key_management"
	"github.com/ibmgaragecloud/key-management-operator/service/key_management/factory"
	"github.com/imdario/mergo"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func newMetadata(name string, labels *map[string]string, annotations *map[string]string) *metav1.ObjectMeta {
	return &metav1.ObjectMeta{
		Name:        name,
		Labels:      *labels,
		Annotations: *annotations,
	}
}

func newSecret(metadata *metav1.ObjectMeta, data *map[string][]byte) *corev1.Secret {
	return &corev1.Secret{
		TypeMeta:   metav1.TypeMeta{Kind: "Secret", APIVersion: "v1"},
		ObjectMeta: *metadata,
		Data:       *data,
	}
}

func GenerateSecret(secretTemplate *keymanagementv1.SecretTemplate) *corev1.Secret {
	kp := *secretTemplate

	keyManager := factory.LoadKeyManager(kp.ObjectMeta.Annotations)

	return generateSecret(keyManager, secretTemplate)
}

func generateSecret(keyManager *key_management.KeyManager, secretTemplate *keymanagementv1.SecretTemplate) *corev1.Secret {
	var values map[string][]byte
	var annotations map[string]string

	values = make(map[string][]byte)
	annotations = make(map[string]string)

	kp := *secretTemplate

	(*keyManager).PopulateMetadata(&annotations)

	specValues := kp.Spec.Values
	for _, kps := range specValues {
		values[*stringValue(&kps.Key, &kps.Name)] = *convertValue(keyManager, &kps, &annotations)
	}

	mergo.Merge(&annotations, kp.Spec.Annotations)

	return newSecret(newMetadata(kp.ObjectMeta.Name, &kp.Spec.Labels, &annotations), &values)
}

func stringValue(value *string, defaultValue *string) *string {
	val := *value
	if val == "" {
		val = *defaultValue
	}

	return &val
}

func convertValue(keyManager *key_management.KeyManager, keyValue *keymanagementv1.SecretTemplateValue, annotations *map[string]string) *[]byte {
	kp := *keyValue

	if kp.KeyId != "" {
		km := *keyManager
		result := processBase64StringValue(km.GetKey(&kp.KeyId))

		a := *annotations

		a[fmt.Sprintf("%s.keyId/%s", km.Id(), kp.Name)] = kp.KeyId

		return result
	} else if kp.Value != "" || kp.StringData != "" {
		return processStringValue(stringValue(&kp.StringData, &kp.Value))
	} else {
		return processBase64StringValue(stringValue(&kp.Data, &kp.B64Value))
	}
}

func processBase64StringValue(b64value *string) *[]byte {
	encodedBytes := []byte(*b64value)

	var unencodedBytes []byte
	unencodedBytes = make([]byte, len(encodedBytes))

	base64.StdEncoding.Decode(unencodedBytes, encodedBytes)

	return &unencodedBytes
}

func processStringValue(value *string) *[]byte {
	val := []byte(*value)

	return &val
}
