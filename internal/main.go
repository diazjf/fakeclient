package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/diazjf/fakeclient/internal/secrets"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	home := homeDir()

	// Grab the KUBECONFIG using either 'kubeconfig' flag, or from 'home/.kube/config'
	if home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// Use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalf("Error Using Current Context: %v", err)
	}

	// Create the 'real' clientset
	// as seen in https://github.com/kubernetes/client-go/blob/master/kubernetes/clientset.go
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error Creating the Client: %v", err)
	}

	// Get Secrets from the 'default' namespace, and print them nicely
	secretMap := secrets.GetSecrets(clientset, "default")
	prettifyOutput(secretMap)
}

// homeDir just returns the home directory
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

// prettifyOutput outputs a secretMap in a nice readable format
func prettifyOutput(secretMap map[string]map[string]string) {
	for k, val := range secretMap {
		fmt.Println(k)

		b, err := json.MarshalIndent(val, "", "  ")
		if err != nil {
			log.Fatalf("Error Marshalling: %v", err)
		}
		fmt.Printf("%s\n\n", string(b))
	}
}
