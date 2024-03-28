package main

import (
	"context"
	"flag"
	"fmt"
	v12 "k8s.io/api/apps/v1"
	v13 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

type AppService struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppServiceSpec   `json:"spec,omitempty"`
	Status AppServiceStatus `json:"status,omitempty"`
}
type AppServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of AppService. Edit appservice_types.go to remove/updatestatus
	Size     *int32                  `json:"size"`
	Image    string                  `json:"image"`
	Resource v13.ResourceRequirements `json:"resource,omitempty"`
	Envs     []v13.EnvVar             `json:"envs,omitempty"`
	Ports    []v13.ServicePort        `json:"ports,omitempty"`
}

type AppServiceStatus struct {
	v12.DeploymentStatus `json:",inline"`
}

func main() {
	h := os.Getenv("HOME")

	f := filepath.Join(h, ".kube", "config.conf")
	kubeconfig := flag.String("kubeconfig", f, "(optional) absolute path to the kubeconfig file")

	masterUrl := ""
	config, err := clientcmd.BuildConfigFromFlags(masterUrl, *kubeconfig)
	if err != nil {

	}

	client, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Println("new config.conf err:", err)
		return
	}

	gvr := schema.GroupVersionResource{
		Group:    "app.drzhangg.io",
		Version:  "v1beta1",
		Resource: "appservices",
	}

	ns := "default"
	name := "test1112"

	un := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "app.drzhangg.io/v1beta1",
			"kind": "AppService",
			"metadata": map[string]interface{}{
				"name":name,
				"namespace": ns,
			},
			"status": map[string]interface{}{
				"conditions": map[string]interface{}{
					"message":"hello hhh",
				},
			},
		},
	}

	_,err = client.Resource(gvr).Namespace(ns).Create(context.Background(),un,v1.CreateOptions{})
	if err !=nil{
		fmt.Println("create err:",err)
	}


}
