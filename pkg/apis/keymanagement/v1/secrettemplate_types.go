package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SecretTemplateValue defines an entry in the generated secret. Each entry can be defined as either a raw
// text string, a base 64 encoded string, or an id for a key in a key management/vault system.
type SecretTemplateValue struct {
	// The name or key for the entry in the secret data (e.g. the secret will contain <Name>: <Value>
	Name string `json:"name"`

	// Value represents a raw value that will be base 64 encoded and injected into the secret
	// +optional
	Value string `yaml:"value,omitempty" json:"value,omitempty"`

	// B64Value represents a base 64 encoded value that will be injected directly into the secret
	// +optional
	B64Value string `yaml:"b64value,omitempty" json:"b64value,omitempty"`

	// KeyId represents an identifier for a key value stored in a key management system. The operator will look up
	// the value of the key and inject it into the generated secret
	// +optional
	KeyId string `yaml:"keyId,omitempty" json:"keyId,omitempty"`
}

// SecretTemplateSpec defines the desired state of SecretTemplate
// +k8s:openapi-gen=true
type SecretTemplateSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and services.
	// More info: http://kubernetes.io/docs/user-guide/labels
	// +optional
	Labels map[string]string `json:"labels,omitempty" protobuf:"bytes,11,rep,name=labels"`

	// Annotations is an unstructured key value map stored with a resource that may be
	// set by external tools to store and retrieve arbitrary metadata. They are not
	// queryable and should be preserved when modifying objects.
	// More info: http://kubernetes.io/docs/user-guide/annotations
	// +optional
	Annotations map[string]string `json:"annotations,omitempty" protobuf:"bytes,12,rep,name=annotations"`

	// List of values that will be inserted into the generated secret
	Values []SecretTemplateValue  `json:"values"`
}

// SecretTemplateStatus defines the observed state of SecretTemplate
type SecretTemplateStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SecretTemplate is the Schema for the secrettemplates API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=secrettemplates,scope=Namespaced
type SecretTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecretTemplateSpec   `json:"spec,omitempty"`
	Status SecretTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SecretTemplateList contains a list of SecretTemplate
type SecretTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecretTemplate{}, &SecretTemplateList{})
}
