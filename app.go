package main

import (
	"flag"
	"fmt"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Running")
	var err error
	var config *rest.Config
	if os.Getenv("KUBERNETES_SERVICE_HOST") == "" {
		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		// use the current context in kubeconfig
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	} else {
		config, err = rest.InClusterConfig()
	}
	if err != nil {
		panic(err.Error())
	}

	apiextensions, err := clientset.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Make sure the CRD exists
	_, err = apiextensions.ApiextensionsV1().CustomResourceDefinitions().Get("dbaasplatforms.dbaas.redhat.com", v1.GetOptions{})

	if err != nil {
		fmt.Println("Error retrieving CRD", err)
	} else {
		fmt.Println("CRD found")
	}
}
