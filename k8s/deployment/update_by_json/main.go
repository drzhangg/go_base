package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

var deployJson = `
{
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
        "annotations": {
            "deployment.kubernetes.io/revision": "1"
        },
        "generation": 1,
        "name": "my-deployment",
        "namespace": "default"
    },
    "spec": {
        "progressDeadlineSeconds": 600,
        "replicas": 2,
        "revisionHistoryLimit": 10,
        "selector": {
            "matchLabels": {
                "app": "my-app"
            }
        },
        "strategy": {
            "rollingUpdate": {
                "maxSurge": "25%",
                "maxUnavailable": "25%"
            },
            "type": "RollingUpdate"
        },
        "template": {
            "metadata": {
                "creationTimestamp": null,
                "labels": {
                    "app": "my-app"
                }
            },
            "spec": {
                "containers": [
                    {
                        "image": "nginx",
                        "imagePullPolicy": "Always",
                        "name": "my-container",
                        "ports": [
                            {
                                "containerPort": 80,
                                "protocol": "TCP"
                            }
                        ],
						"env": [
							{
								"name": "DEBUG",
								"value": "true"
							},
							{
								"name": "DEBUG1",
								"value": "false"
							}
						],
                        "resources": {},
                        "terminationMessagePath": "/dev/termination-log",
                        "terminationMessagePolicy": "File"
                    }
                ],
                "dnsPolicy": "ClusterFirst",
                "restartPolicy": "Always",
                "schedulerName": "default-scheduler",
                "securityContext": {},
                "terminationGracePeriodSeconds": 30
            }
        }
    }
}

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

	obj,gvk,err := unstructured.UnstructuredJSONScheme.Decode([]byte(deployJson),nil,nil)
	if err != nil {
		fmt.Println("unstructured decode err:",err)
		return
	}

	unObj,err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		fmt.Println("unstructured ToUnstructured err:",err)
		return
	}

	fmt.Println("gvk1::",gvk)
	fmt.Println("un::",unObj)


	dc := clientset.Discovery()

	restMapper,err := restmapper.GetAPIGroupResources(dc)
	if err != nil {
		fmt.Println("get api group resource err:",err)
		return
	}

	restMap := restmapper.NewDiscoveryRESTMapper(restMapper)

	mapping,err := restMap.RESTMapping(gvk.GroupKind(),gvk.Version)
	if err != nil {
		fmt.Println("unstructured decode err:",err)
		return
	}

	var unstruct unstructured.Unstructured
	unstruct.Object = unObj

	fmt.Println("mapping::",mapping)

	_,err = dynamicClient.Resource(mapping.Resource).Namespace(unstruct.GetNamespace()).Update(context.Background(),&unstruct,v1.UpdateOptions{})
	if err != nil {
		fmt.Println("dynamic update error:",err)
		return
	}
}
