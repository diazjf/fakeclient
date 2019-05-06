package secrets

import (
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetSecrets returns a Map containing all the Data within a Secret in String Format
// It accepts a kubernetes.Interface, in which a "real" or "fake" ClientSet can be passed
func GetSecrets(clientSet kubernetes.Interface, namespace string) map[string]map[string]string {
	list, err := clientSet.CoreV1().Secrets(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error Listing Secrets: %v", err)
	}

	secretMap := make(map[string]map[string]string)
	for _, s := range list.Items {
		secretData := s.Data
		secretDataMap := make(map[string]string)
		for k, v := range secretData {
			secretDataMap[k] = string(v)
		}
		secretMap[s.Name] = secretDataMap
	}

	return secretMap
}
