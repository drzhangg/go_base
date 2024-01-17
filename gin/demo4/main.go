package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/admission/v1"
	v13 "k8s.io/api/apps/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func main() {
	client := NewClient()

	r := gin.Default()
	r.POST("/addAffinity",addAffinity)
	r.GET("/api/v1/namespaces/:namespace/configmap/:name",client.getCm)
	r.GET("/api/v1/namespaces/:namespace/configmaps",client.getCms)

	r.Run(":8881")
}

type Client struct {
	clientset *kubernetes.Clientset
}

func NewClient() *Client{
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

	return &Client{clientset: clientset}
}

func (cl *Client)getCms(c *gin.Context)  {
	ns := c.Param("namespace")
	cm,err := cl.clientset.CoreV1().ConfigMaps(ns).List(context.TODO(),v12.ListOptions{})
	if err != nil {
		fmt.Println("errrr:",err)
	}

	data,err := json.Marshal(cm)
	if err != nil {
		fmt.Println("sda:",err)
	}
	re := make(map[string]interface{})

	json.Unmarshal(data,&re)

	c.JSON(200,re)

}

func (cl *Client)getCm(c *gin.Context)  {
	ns := c.Param("namespace")
	name := c.Param("name")
	cm,err := cl.clientset.CoreV1().ConfigMaps(ns).Get(context.TODO(),name,v12.GetOptions{})
	if err != nil {
		fmt.Println("errrr:",err)
	}

	c.JSON(200,cm)

}

func addAffinity(c *gin.Context) {
	//newDeploy(&ar)
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
