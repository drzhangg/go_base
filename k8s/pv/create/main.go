package main

import (
	"context"
	"flag"
	"fmt"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	pv := v12.PersistentVolume{
		ObjectMeta: v1.ObjectMeta{
			Name:                       "test-pv",
		},
		Spec:       v12.PersistentVolumeSpec{
			Capacity:                      v12.ResourceList{
				"storage": resource.Quantity{},
			},
			PersistentVolumeSource:        v12.PersistentVolumeSource{},
			AccessModes:                   []v12.PersistentVolumeAccessMode{
				"ReadWriteOnce",
			},
			ClaimRef:                      nil,
			PersistentVolumeReclaimPolicy: "",
			StorageClassName:              "",
			MountOptions:                  nil,
			VolumeMode:                    nil,
			NodeAffinity:                  nil,
		},
		Status:     v12.PersistentVolumeStatus{},
	}


	_,err = clientset.CoreV1().PersistentVolumes().Create(context.TODO(),&pv,v1.CreateOptions{})
	if err !=nil{
		fmt.Println("err:",err)
	}

}
