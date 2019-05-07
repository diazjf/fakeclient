package secrets_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/diazjf/fakeclient/internal/secrets"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes/fake"
)

func TestGetSecrets(t *testing.T) {

	// create the 'fake' clientSet where clientset.Interface = &Clientset{}, setting all the 'fake' methods
	// as seen in https://github.com/kubernetes/client-go/blob/master/kubernetes/fake/clientset_generated.go
	clientSet := fake.NewSimpleClientset()
	namespace := "default"

	// Generate Secret 1
	data1 := make(map[string][]byte)
	data1["username"] = []byte("us3r1")
	data1["password"] = []byte("p4ssw0rd2")
	s1 := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-credentials",
			Namespace: namespace,
		},
		Data: data1,
	}
	_, err := clientSet.CoreV1().Secrets(namespace).Create(s1)
	if err != nil {
		assert.Equal(t, nil, err)
	}

	// Generate Secret 2
	data2 := make(map[string][]byte)
	data2["apikey"] = []byte("ABCDEFGHIJKLMNOP")
	s2 := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-api-key",
			Namespace: namespace,
		},
		Data: data2,
	}
	_, err = clientSet.CoreV1().Secrets(namespace).Create(s2)
	if err != nil {
		assert.Equal(t, nil, err)
	}

	decodedMap := secrets.GetSecrets(clientSet, "default")
	assert.Equal(t, 2, len(decodedMap))
	assert.Equal(t, "us3r1", string(decodedMap["my-credentials"]["username"]))
	assert.Equal(t, "p4ssw0rd2", string(decodedMap["my-credentials"]["password"]))
	assert.Equal(t, "ABCDEFGHIJKLMNOP", string(decodedMap["my-api-key"]["apikey"]))
}
