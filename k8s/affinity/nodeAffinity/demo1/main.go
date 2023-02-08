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

	f := filepath.Join(h, ".kube", "config.conf")
	kubeconfig := flag.String("kubeconfig", f, "(optional) absolute path to the kubeconfig file")

	masterUrl := ""
	config, err := clientcmd.BuildConfigFromFlags(masterUrl, *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("asaqweqwe:", err)
	}

	deploy := newDeployment()

	_, err = clientset.AppsV1().Deployments("default").Create(context.TODO(), deploy, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("err:", err)
	}
}

func newDeployment() *v1.Deployment {
	var replicas int32 = 2
	return &v1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-deployment",
			Labels: map[string]string{
				"app": "test-deploy",
			},
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "test-deploy",
				},
			},
			Template: v12.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: "nginx",
					Labels: map[string]string{
						"app": "test-deploy",
					},
				},
				Spec: v12.PodSpec{
					Affinity: &v12.Affinity{ // if label == nil , NodeAffinity == nil
						NodeAffinity: &v12.NodeAffinity{
							RequiredDuringSchedulingIgnoredDuringExecution: &v12.NodeSelector{ // 硬策略
								NodeSelectorTerms: []v12.NodeSelectorTerm{
									{
										MatchExpressions: []v12.NodeSelectorRequirement{
											{
												Key:      "app",
												Operator: "In",
												Values:   []string{"test"},
											},
										},
									},
								},
							},
							//PreferredDuringSchedulingIgnoredDuringExecution: []v12.PreferredSchedulingTerm{ // 软策略
							//	{
							//		Weight: 1,
							//		Preference: v12.NodeSelectorTerm{
							//			MatchExpressions: []v12.NodeSelectorRequirement{
							//				{
							//					Key:      "app",
							//					Operator: v12.NodeSelectorOpIn,
							//					Values:   []string{"test"},
							//				},
							//			},
							//		},
							//	},
							//},
						},
					},
					Containers: []v12.Container{
						{
							Name:  "nginx",
							Image: "nginx",
							Ports: []v12.ContainerPort{
								{
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
}
