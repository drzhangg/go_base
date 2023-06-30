package main

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Pipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

func main() {

	p := &Pipeline{}

	if p.Annotations == nil {
		p.Annotations = map[string]string{}
	}
	p.Annotations["test"] = "finish"
	fmt.Println(p.Annotations["test"])
}
