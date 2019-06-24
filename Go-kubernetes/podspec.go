package main

import (
    "fmt"
    "flag"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/kubernetes"
)

func main() {
    kubeconfig := flag.String("kubeconfig", "/Users/vigneshwarviswanathan/.kube/config", "path to the kubeconfig file")
    flag.Parse()
    config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
    clientset, err := kubernetes.NewForConfig(config)
    pod, err := clientset.CoreV1().Pods("default").Get("app-6c7597db55-mk4qn", metav1.GetOptions{})
    fmt.Println(pod)
    if err != nil {
      panic(err)
    }
}
