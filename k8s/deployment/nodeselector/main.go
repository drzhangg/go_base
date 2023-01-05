package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func main() {
	h := os.Getenv("HOME")

	f := filepath.Join(h, ".kube", "config")
	kubeconfig := flag.String("kubeconfig", f, "(optional) absolute path to the kubeconfig file")

	masterUrl := ""
	config, err := clientcmd.BuildConfigFromFlags(masterUrl, *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("kubernetes.NewForConfig err:", err)
	}
	deploy := newDeployment()

	dep, err := clientset.AppsV1().Deployments("default").Create(context.TODO(), deploy, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("create deploy err:", err)
	}
	//fmt.Println(dep)
	dep = dep
}

func newDeployment() *v1.Deployment {
	labels := map[string]string{
		"app": "test",
	}

	var replicas *int32
	var num int32 = 2
	replicas = &num

	return &v1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   "test-deploy",
			Labels: labels,
		},
		Spec: v1.DeploymentSpec{
			Replicas: replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: v12.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:   "test-deploy",
					Labels: labels,
				},
				Spec: v12.PodSpec{
					Containers: []v12.Container{
						{
							Name:  "test-deploy",
							Image: "nginx",
							Ports: []v12.ContainerPort{
								{
									ContainerPort: 80,
								},
							},
							//Resources: v12.ResourceRequirements{
							//	Limits: v12.ResourceList{
							//		v12.ResourceCPU: resource.Quantity{},
							//	},
							//	Requests: nil,
							//},
						},
					},
					NodeSelector: map[string]string{
						"app": "test",
					},
				},
			},
		},
	}
}
