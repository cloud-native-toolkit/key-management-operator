package generate_secret

import (
	"bytes"
	"encoding/base64"
	"fmt"
	keymanagementv1 "github.com/ibm-garage-cloud/key-management-operator/api/v1"
	"github.com/ibm-garage-cloud/key-management-operator/service/key_management"
	"github.com/ibm-garage-cloud/key-management-operator/service/key_management/factory"
	"github.com/imdario/mergo"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func newMetadata(objectMetadata *metav1.ObjectMeta, labels *map[string]string, annotations *map[string]string) *metav1.ObjectMeta {
	om := *objectMetadata

	result := metav1.ObjectMeta{
		Name:        om.Name,
		Labels:      *labels,
		Annotations: *annotations,
	}

	if om.Namespace != "" {
		result.Namespace = om.Namespace
	}

	return &result
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
		key := *stringValue(&kps.Key, &kps.Name)
		value := *convertValue(keyManager, &kps, &annotations)

		values[key] = value
	}

	mergo.Merge(&annotations, kp.Spec.Annotations)

	return newSecret(newMetadata(&kp.ObjectMeta, &kp.Spec.Labels, &annotations), &values)
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

		annotationName := *stringValue(&kp.Key, &kp.Name)

		a[fmt.Sprintf("%s.keyId/%s", km.Id(), annotationName)] = kp.KeyId

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

	trimmedBytes := bytes.Trim(unencodedBytes, "\x00")

	return &trimmedBytes
}

func processStringValue(value *string) *[]byte {
	val := []byte(*value)

	return &val
}
