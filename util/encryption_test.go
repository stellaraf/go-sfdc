package util

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func Test_encrypt(t *testing.T) {
	t.Run("encrypts and decrypts value", func(t *testing.T) {
		key := randomString(32)
		value := randomString(128)
		encrypted, err := Encrypt(key, value)
		assert.NoError(t, err)
		decrypted, err := Decrypt(key, encrypted)
		assert.NoError(t, err)
		t.Logf("key=%s, value=%s, encrypted=%s, decrypted=%s", key, value, encrypted, decrypted)
		t.Logf("equal %t", decrypted == value)
		assert.Equal(t, value, decrypted)
	})
	// t.Run("one off", func(t *testing.T) {
	// 	key := "dont4get"
	// 	value := "566c4ec80db3f301-3616a2cd55fd673319b9cfa0-f3b36a7e4b01ae1b85b72c30d9a2f4cae76f9b68ea10d54f"
	// 	decrypted, err := Decrypt(key, value)
	// 	t.Logf("key=%s, value=%s, decrypted=%s", key, value, decrypted)
	// 	t.Logf("equal %t", decrypted == value)
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, value, decrypted)
	// })
}
