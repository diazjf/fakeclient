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
	clientSet := fake.NewSimpleClientset()
	namespace := "default"

	data1 := make(map[string][]byte)
	data1["password"] = []byte("p4ssw0rd1")
	s1 := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "secret1",
			Namespace: namespace,
		},
	}
	_, err := clientSet.CoreV1().Secrets(namespace).Create(s1)
	if err != nil {
		assert.Equal(t, nil, err)
	}

	data2 := make(map[string][]byte)
	data2["password"] = []byte("p4ssw0rd2")
	s2 := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "secret2",
			Namespace: namespace,
		},
	}
	_, err = clientSet.CoreV1().Secrets(namespace).Create(s2)
	if err != nil {
		assert.Equal(t, nil, err)
	}

	decodedMap := secrets.GetSecrets(clientSet, "default")
	assert.Equal(t, 2, len(decodedMap))
}
