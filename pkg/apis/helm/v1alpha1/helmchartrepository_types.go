package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HelmChartRepositorySpec defines the desired state of HelmChartRepository
type HelmChartRepositorySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// HelmChartRepositoryStatus defines the observed state of HelmChartRepository
type HelmChartRepositoryStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HelmChartRepository is the Schema for the helmchartrepositories API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=helmchartrepositories,scope=Namespaced
type HelmChartRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HelmChartRepositorySpec   `json:"spec,omitempty"`
	Status HelmChartRepositoryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HelmChartRepositoryList contains a list of
// +kubebuilder:resource:path=helmchartrepositories,scope=Cluster
type HelmChartRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HelmChartRepository `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HelmChartRepository{}, &HelmChartRepositoryList{})
}
