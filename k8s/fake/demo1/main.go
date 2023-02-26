package main

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
)

func main() {
	
	svc := &v1.Service{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "v2",
			APIVersion: "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:                       "test svc",
			Namespace:                  "test",
			Labels: map[string]string{
				"name":"jerry",
				"address":"wuhan",
			},
			Annotations: map[string]string{
				"tag":"true",
			},
		},
		Spec:       v1.ServiceSpec{},
		Status:     v1.ServiceStatus{},
	}

	mockClient := fake.NewSimpleClientset(svc)
	s,err := getService(mockClient,"test svc","test")
	if err != nil {
		fmt.Println("err:",err)
	}
	fmt.Println(s)
	
}

func getService(client kubernetes.Interface, name, namespace string) (*v1.Service,error) {

	svc,err := client.CoreV1().Services(namespace).Get(context.Background(),name,metav1.GetOptions{})
	if err != nil {
		return nil,fmt.Errorf(fmt.Sprintf("get service error:",err))
	}
	return svc,nil
}
