package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
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

	pod := newPod()
	pod, err = clientset.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("err:", err)
	}


	bytes,err := json.Marshal(pod)
	if err !=nil{
		fmt.Println("json marshal failed,err:",err)
	}
	fmt.Println("pod:",string(bytes))
}

func newPod() *v12.Pod {

	// gpuNodeSelector,gradeNodeSelector := v12.NodeSelectorRequirement{},v12.NodeSelectorRequirement{}
	//
	// reflect.DeepEqual(gpuNodeSelector,v12.NodeSelectorRequirement{})
	//
	//if gpuNodeSelector != v12.NodeSelectorRequirement{} {
	//
	//}

	var matchExpression  []v12.NodeSelectorRequirement
	matchExpression = append(matchExpression, )

	return &v12.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "with-node-affinity1",
		},

		Spec: v12.PodSpec{
			Containers: []v12.Container{
				{
					Name:  "with-node-affinity",
					Image: "k8s.gcr.io/pause:2.0",
				},
			},
			Affinity: &v12.Affinity{ // if label == nil , NodeAffinity == nil
				NodeAffinity: &v12.NodeAffinity{
					RequiredDuringSchedulingIgnoredDuringExecution: &v12.NodeSelector{ // 硬策略
						NodeSelectorTerms: []v12.NodeSelectorTerm{
							{
								MatchExpressions: []v12.NodeSelectorRequirement{
									{
										Key:      "app",
										Operator: v12.NodeSelectorOpIn,
										Values:   []string{""},
									},
								},
							},
						},
					},
					PreferredDuringSchedulingIgnoredDuringExecution: []v12.PreferredSchedulingTerm{ // 软策略
						{
							Weight: 1,
							Preference: v12.NodeSelectorTerm{
								MatchExpressions: []v12.NodeSelectorRequirement{
									{
										Key:      "another-node-label-key",
										Operator: v12.NodeSelectorOpIn,
										Values:   []string{"another-node-label-value"},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
