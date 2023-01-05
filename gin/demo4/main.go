package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/admission/v1"
	v13 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func main() {

	r := gin.Default()
	r.POST("/addAffinity",addAffinity)
}

func addAffinity(c *gin.Context) {
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
		fmt.Println("asaqweqwe:", err)
	}

	ar := v1.AdmissionReview{}
	err = c.ShouldBindJSON(&ar)
	if err !=nil {
		fmt.Println("bind json err:",err)
	}

	clientset = clientset

	newDeploy(&ar)

}

func newDeploy(ar *v1.AdmissionReview) *v13.Deployment {
	req := ar.Request


	var deploy v13.Deployment
	err := json.Unmarshal(req.Object.Raw,&deploy)
	if err !=nil{
		fmt.Println("unmarshal err:",err)
	}
	return nil
}
