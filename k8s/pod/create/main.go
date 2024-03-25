package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func main() {
	h := os.Getenv("HOME")
	f := filepath.Join(h, ".kube", "config")
	kubeconfig := flag.String("kubeconfig", f, "(optional) absolute path to the kubeconfig file")

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("asaqweqwe:", err)
	}

	rc := clientset.RESTClient()

	//namespace := "default"
	//deployName := "test"

	//deployPath := "apis/apps/v1/namespaces/default/deployments"


	podPath := "api/v1/namespaces/default/pods"

	data1, err :=json.Marshal(newResource())
	if err != nil {
		fmt.Println("marshal pod failed,err:",err)
	}

	fmt.Println(string(data1))

	result, err := rc.Post().AbsPath(podPath).Body(data1).SetHeader("Accept", "application/json").DoRaw(context.Background())
	if err != nil {
		fmt.Println("post err:",err)
		panic(err.Error())
	}

	rc.Delete().AbsPath()

	fmt.Println("resss:",string(result))

	//flunder := `{"apiVersion":"` + "v1" + `","kind":"Pod","metadata":{"labels":{"sample-label":"true"},"name":"` + "test111" + `","namespace":"default"},"spec":{"containers":[{"image":"nginx:latest","imagePullPolicy":"Always","name":"nginx11"}]}}`
	//result1 := rc.Post().AbsPath("/api/v1/namespaces/default/pods").Body([]byte(flunder)).SetHeader("Accept", "application/json").Do(context.Background())
	//if result1.Error() != nil {
	//	fmt.Println("err:::", result1.Error().Error())
	//}
	//var statusCode int
	//result1.StatusCode(&statusCode)
	//
	//fmt.Println("code::", statusCode)

	//r := rc.Patch(types.StrategicMergePatchType).AbsPath(deployPath).Body(data)

}

type K8SMetadata struct {
	Name        string            `json:"name"`
	Namespace   string            `json:"namespace"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`

	ResourceVersion   string                   `json:"ResourceVersion,omitempty"`
	CreationTimestamp string                   `json:"creationTimestamp,omitempty"`
	DeletionTimestamp string                   `json:"deletionTimestamp,omitempty"`
	OwnerReferences   []map[string]interface{} `json:"ownerReferences,omitempty"`
	Finalizers        []string                 `json:"finalizers,omitempty"`
	ClusterName       string                   `json:"clusterName,omitempty"`
	GenerateName      string                   `json:"generateName,omitempty"`
	Generation        int32                    `json:"generation,omitempty"`
	SelfLink          string                   `json:"selfLink,omitempty"`
	Uid               string                   `json:"uid,omitempty"`
}
type K8sResource struct {
	Kind       string                 `json:"kind"`
	APIVersion string                 `json:"apiVersion"`
	Metadata   K8SMetadata            `json:"metadata"`
	Spec       map[string]interface{} `json:"spec"`
	Status     map[string]interface{} `json:"status,omitempty"`
}

func newResource() K8sResource {
	spec := make(map[string]interface{})
	containers := []map[string]string{}
	containers = append(containers, map[string]string{
		"image":           "nginx:latest",
		"name":            "nginx1",
		"imagePullPolicy": "Always",
	})

	spec["containers"] = containers

	return K8sResource{
		Kind:       "Pod",
		APIVersion: "v1",
		Metadata: K8SMetadata{
			Namespace: "default",
			Name:      "test132",
		},
		Spec:   spec,
		Status: nil,
	}
}
