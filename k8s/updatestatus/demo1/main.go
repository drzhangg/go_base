package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)


var devopsProject = schema.GroupVersionResource{
	Group:    "ship.my.domain",
	Version:  "v1beta1",
	Resource: "frigates",
}

type Frigate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FrigateSpec   `json:"spec,omitempty"`
	Status FrigateStatus `json:"status,omitempty"`
}

type FrigateSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Frigate. Edit frigate_types.go to remove/update
	Foo         string `json:"foo,omitempty"`
	Name        string `json:"name,omitempty"`
	Addr        string `json:"addr,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image"`
}

// FrigateStatus defines the observed state of Frigate
type FrigateStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	AdminNamespace string `json:"adminNamespace,omitempty"`
}

func main() {
	h := os.Getenv("HOME")
	f := filepath.Join(h, ".kube", "config")
	kubeconfig := flag.String("kubeconfig", f, "(optional) absolute path to the kubeconfig file")

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("1111:", err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Println("222:", err)
	}

	//workspaceLabel := "kubesphere.io/workspace"
	//workspaceName := "devops-workspace"

	obj,err := dynamicClient.Resource(devopsProject).Namespace("default").Get(context.TODO(),"frigate-sample",metav1.GetOptions{})
	if err != nil {
		fmt.Println("333:", err)
	}

	dd,err := obj.MarshalJSON()
	if err != nil {
		fmt.Println("44:", err)
	}

	instanc := &Frigate{}
	err = json.Unmarshal(dd,instanc)
	if err != nil {
		fmt.Println("55:", err)
	}

	//fmt.Println(instanc.Name)

	instanc.Status.AdminNamespace = "1111"


	_,err = dynamicClient.Resource(devopsProject).Namespace("default").Update(context.TODO(),&unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind": instanc.Kind,
			"apiVersion":instanc.APIVersion,
			"metadata":map[string]interface{}{
				"name":instanc.Name,
				"resourceVersion": instanc.ResourceVersion,
				"labels":instanc.Labels,
				"managedFields":instanc.ManagedFields,
				"annotations":instanc.Annotations,
				"generateName": instanc.Name,
			},
			"spec":instanc.Spec,
		},
	},metav1.UpdateOptions{})
	if err != nil {
		fmt.Println("update status err:",err)
	}

	//_,err = dynamicClient.Resource(devopsProject).Namespace("default").UpdateStatus(context.Background(),&unstructured.Unstructured{
	//	Object: map[string]interface{}{
	//		"kind": instanc.Kind,
	//		"apiVersion":instanc.APIVersion,
	//		"metadata":map[string]interface{}{
	//			"name":instanc.Name,
	//			"resourceVersion": instanc.ResourceVersion,
	//			"labels":instanc.Labels,
	//			"managedFields":instanc.ManagedFields,
	//			"annotations":instanc.Annotations,
	//			"generateName": instanc.Name,
	//		},
	//		"spec":instanc.Spec,
	//		"status":instanc.Status,
	//	},
	//},metav1.UpdateOptions{})
	//if err != nil {
	//	fmt.Println("update status err:",err)
	//}


}
