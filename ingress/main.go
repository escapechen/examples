/*
Copyright 2016 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	//"fmt"
	//"time"
	"github.com/davecgh/go-spew/spew"
        "io"
        "net/http"

	//"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		narf := func(w http.ResponseWriter, _ *http.Request) {
			//pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
			ingresses, err := clientset.ExtensionsV1beta1().Ingresses("").List(metav1.ListOptions{})
			if err != nil {
				panic(err.Error())
			}
			for _, ing := range ingresses.Items {
				//spew.Dump(ing)
				//needs major improvement, but hey output! 8)
				mystring := spew.Sprintf(":: %s/%s > %s\n", ing.Namespace, ing.Name, ing.Spec.Rules)
				io.WriteString(w, mystring)
			}
		}
		http.HandleFunc("/", narf)
	        http.ListenAndServe(":8000", nil)
	}
}

