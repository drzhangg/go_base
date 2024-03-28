package main

import (
	"context"
	"encoding/json"
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

func main() {
	h := os.Getenv("HOME")

	f := filepath.Join(h, ".kube", "config.conf")
	kubeconfig := flag.String("kubeconfig", f, "(optional) absolute path to the kubeconfig file")

	masterUrl := ""
	config, err := clientcmd.BuildConfigFromFlags(masterUrl, *kubeconfig)
	if err != nil {
		//panic(err)
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

	list, err := client.Resource(gvr).List(context.Background(), v1.ListOptions{})
	if err != nil {
		fmt.Println("get crd list err:", err)
		return
	}

	var crlist,crnslist []string

	for _,v := range list.Items{
		crlist = append(crlist, v.GetName())
		crnslist = append(crnslist, v.GetNamespace())
	}

	fmt.Println("crd list is:",crlist)
	fmt.Println("crd namespace list is:",crnslist)

	namespace := "default"

	crName := "appservice-nginx"

	apps, err := client.Resource(gvr).Namespace(namespace).Get(context.Background(), crName, v1.GetOptions{})
	if err != nil {
		fmt.Println("get crd err:", err)
		return
	}

	data,err := apps.MarshalJSON()
	if err != nil {
		fmt.Println("apps marshal err:", err)
	}

	as := AppService{}

	json.Unmarshal(data,&as)

	//as.Status.Conditions[0].Message = "this is message"

	//fmt.Println("as ::",as.Status)
	
	as.Status = AppServiceStatus{
		v12.DeploymentStatus{
			Conditions:          []v12.DeploymentCondition{
				{
					Message: "this is message",
				},
			},
		},
	}


	//result := apps.UnstructuredContent()
	//fmt.Println("result::", result)
	//
	_,err = client.Resource(gvr).Namespace(namespace).UpdateStatus(context.Background(),&unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind": as.Kind,
			"apiVersion":as.APIVersion,
			"metadata":map[string]interface{}{
				"name":as.Name,
				"resourceVersion": as.ResourceVersion,
				"labels":as.Labels,
				"managedFields":as.ManagedFields,
				"annotations":as.Annotations,
			},
			"spec":as.Spec,
			"status":as.Status,
		},
	},v1.UpdateOptions{})
	if err != nil {
		fmt.Println("updatestatus status err:",err)
	}

}

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
