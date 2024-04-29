package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"gopkg.in/yaml.v3"
	"io/ioutil"
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

	masterUrl := ""
	config1, err := clientcmd.BuildConfigFromFlags(masterUrl, *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config1)
	if err != nil {
		fmt.Println("asaqweqwe:", err)
	}

	ns := "default"
	name := "test-role"

	rolecm,err := clientset.CoreV1().ConfigMaps(ns).Get(context.Background(),name,v1.GetOptions{})
	if err != nil {
		fmt.Println("get cm failed:",err)
	}

	err = ioutil.WriteFile("yaml/demo1/config.yaml",[]byte(rolecm.Data["role.conf"]),0644)
	if err != nil {
		fmt.Println("get cm failed:",err)
	}

	c := config.New(
		config.WithSource(
			file.NewSource("/Users/drzhang/demo/go/go_base/yaml/demo1/config.yaml"),
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			configData := kv.Value
			//configData = uConf.ReplaceEnv(configData)
			return yaml.Unmarshal(configData, v)
		}),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	type Role struct {
		Role map[string]map[string][]string `json:"role"`
	}

	var role Role
	if err := c.Scan(&role); err != nil {
		panic(err)
	}

	fmt.Println("role::", role.Role["admin"])

	admin := role.Role["admin"]

	deleteValueFromSlice(admin,"api.clickhouse.v1.ClickhouseService","update")

	fmt.Println("role",role)

	yamldata,err := yaml.Marshal(&role)
	if err!=nil{
		fmt.Println("yaml marshal failed:",err)
	}

	//fmt.Println("yaml::",string(yamldata))

	rolecm.Data["role.conf"] = string(yamldata)
	//
	_,err = clientset.CoreV1().ConfigMaps(ns).Update(context.Background(),rolecm,v1.UpdateOptions{})
	if err!=nil{
		fmt.Println("configmap update failed:",err)
	}

}

func deleteValueFromSlice(m map[string][]string, key, value string) {
	slice := m[key]

	var index int  = -1
	for i,v := range slice{
		if v  == value{
			index = i
			break
		}
	}

	if index != -1{
		m[key] = append(slice[:index], slice[index+1:]...)
	}
}
