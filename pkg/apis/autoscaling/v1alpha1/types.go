package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ClusterAutoscaler `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterAutoscaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ClusterAutoscalerSpec   `json:"spec"`
	Status            ClusterAutoscalerStatus `json:"status,omitempty"`
}

type ClusterAutoscalerSpec struct {
	ResourceLimits       *ResourceLimits  `json:"resourceLimits,omitempty"`
	ScaleDown            *ScaleDownConfig `json:"scaleDown,omitempty"`
	MaxPodGracePeriod    *int32           `json:"maxPodGracePeriod,omitempty"`
	PodPriorityThreshold *int32           `json:"PodPriorityThreshold,omitempty"`
}

type ClusterAutoscalerStatus struct {
	// Fill me
}

type ResourceLimits struct {
	MaxNodesTotal *int32         `json:"maxNodesTotal,omitempty"`
	Cores         *ResourceRange `json:"cores,omitempty"`
	Memory        *ResourceRange `json:"memory,omitempty"`
	GPUS          []GPULimit     `json:"gpus,omitempty"`
}

type GPULimit struct {
	Type string `json:"type"`
	ResourceRange
}

type ResourceRange struct {
	Min int32 `json:"min"`
	Max int32 `json:"max"`
}

type ScaleDownConfig struct {
	Enabled           bool   `json:"enabled"`
	DelayAfterAdd     string `json:"delayAfterAdd"`
	DelayAfterDelete  string `json:"delayAfterDelete"`
	DelayAfterFailure string `json:"delayAfterFailure"`
}