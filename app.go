package main

import (
	"fmt"
	"github.com/RHEcosystemAppEng/dbaas-e2e-test-harness/pkg/metadata"
	"github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func main() {
	fmt.Println("Running")

	defer ginkgo.GinkgoRecover()
	config, err := rest.InClusterConfig()
	fmt.Println("Config:", config)
	fmt.Println("Error:", err)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	apiextensions, err := clientset.NewForConfig(config)
	fmt.Println("API extensions:", apiextensions)
	fmt.Println("Error:", err)
	//Expect(err).NotTo(HaveOccurred())

	// Make sure the CRD exists
	obj, err := apiextensions.ApiextensionsV1().CustomResourceDefinitions().Get("dbaasplatforms.dbaas.redhat.com", v1.GetOptions{})
	fmt.Println("obj:", obj)
	fmt.Println("Error:", err)

	if err != nil {
		metadata.Instance.FoundCRD = false
		fmt.Println(err)
	} else {
		metadata.Instance.FoundCRD = true
		fmt.Println(obj)
	}

	Expect(err).NotTo(HaveOccurred())
}
