package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// 根据CRD定义Crontab结构体
type CronTab struct {
	metav1.TypeMeta
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CronTabSpec `json:"spec"`
}

func (c CronTab) DeepCopyObject() runtime.Object {
	panic("implement me")
}

// +k8s:deepcopy-gen=false
type CronTabSpec struct {
	CronSpec string `json:"cronSpec"`
	Image    string `json:"image"`
	Replicas int    `json:"replicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CronTabList struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []CronTab `json:"items"`
}

func (c CronTabList) DeepCopyObject() runtime.Object {
	panic("implement me")
}

