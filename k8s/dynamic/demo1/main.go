package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

var podYaml = `
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: my-statefulset
  namespace: default
spec:
  serviceName: my-statefulset
  replicas: 3
  selector:
    matchLabels:
      app: my-app
  template:
    metadata:
      labels:
        app: my-app
    spec:
      containers:
      - name: my-container
        image: nginx
        ports:
        - containerPort: 80
`

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
		fmt.Println("new client set config err:", err)
		return
	}

	dynamicClient,err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Println("new dynamic client set config err:", err)
		return
	}

	decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewBufferString(podYaml),100)

	dc := clientset.Discovery()

	restMapper,err := restmapper.GetAPIGroupResources(dc)
	if err != nil {
		fmt.Println("get api group resource err:",err)
		return
	}

	restMap := restmapper.NewDiscoveryRESTMapper(restMapper)

	for {
		ext := runtime.RawExtension{}

		if err := decoder.Decode(&ext);err != nil {
			if err == io.EOF{
				break
			}
			fmt.Println("decode err:",err)
		}

		obj,gvk,err := unstructured.UnstructuredJSONScheme.Decode(ext.Raw,nil,nil)
		if err != nil {
			fmt.Println("unstructured decode err:",err)
			return
		}

		mapping,err := restMap.RESTMapping(gvk.GroupKind(),gvk.Version)
		if err != nil {
			fmt.Println("unstructured decode err:",err)
			return
		}

		fmt.Printf("mapping:+%v\n",mapping)

		unObj,err :=  runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			fmt.Println("unstructured ToUnstructured err:",err)
			return
		}

		//fmt.Printf("ToUnstructured:+%v\n",unObj)

		var unstruct unstructured.Unstructured
		unstruct.Object = unObj


		res,err := dynamicClient.Resource(mapping.Resource).Namespace(unstruct.GetNamespace()).Create(context.Background(),&unstruct,v1.CreateOptions{})
		if err != nil {
			fmt.Println("dynamic create error:",err)
			return
		}

		fmt.Println("res:::",res)

	}

	//decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewBufferString(podYaml),100)
	//runtimeScheme := runtime.NewScheme()
	//codecFactory := serializer.NewCodecFactory(runtimeScheme)
	//decoder.Decode(&runtimeScheme)
	//
	//var obj runtime.Unstructured
	//err = decoder.Decode(&obj)
	//if err != nil {
	//	fmt.Println("decode yaml error:",err)
	//	return
	//}

}
