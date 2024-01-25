package main

import (
	"context"
	"flag"
	"fmt"
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


	getcj,err := clientset.BatchV1().CronJobs("default").Get(context.TODO(),"hello",v1.GetOptions{})
	if err != nil {
		fmt.Println("patch cronjob err:", err)
	}



	var su = new(bool)
	*su = true
	getcj.Spec.Suspend = su
	//
	//cj := v12.CronJob{
	//	TypeMeta:   v1.TypeMeta{},
	//	ObjectMeta: v1.ObjectMeta{},
	//	Spec:       v12.CronJobSpec{
	//		Suspend: su,
	//	},
	//	Status:     v12.CronJobStatus{},
	//}

	//data,err := json.Marshal(&getcj)

	cjj,err := clientset.BatchV1().CronJobs("default").Update(context.TODO(),getcj,v1.UpdateOptions{})

	//cjj,err := clientset.BatchV1().CronJobs("default").Patch(context.TODO(),"hello",types.MergePatchType,data,v1.PatchOptions{})
	if err != nil {
		fmt.Println("patch cronjob err:", err)
	}

	fmt.Println("cjjjj:",cjj)

	cmList, err := clientset.CoreV1().ConfigMaps("").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		fmt.Println("get cm list err:", err)
	}

	var ma = make(map[string]int)
	for _, v := range cmList.Items {
		if v.Name == "cm-demo" {
			for key, _ := range v.Data {
				ma[key] = 0
			}
			//fmt.Println("data:",v.Data)
		}
	}
	fmt.Println("ma::", ma)

}
