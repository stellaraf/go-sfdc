package _auth

import (
	"strings"
)

func formatPrivateKey(rawKey string) (key string) {
	lines := []string{}
	for _, line := range strings.Split(rawKey, "\n") {
		newLine := strings.TrimSpace(line)
		lines = append(lines, newLine)
	}
	key = strings.Join(lines, "\n")
	return
}

// func parseToken(token string, privateKey string) (parsed *jwt.Token, err error) {
// 	parsed, err = jwt.Parse(token, func(t *jwt.Token) (key any, err error) {
// 		k := formatPrivateKey(privateKey)
// 		key = []byte(k)
// 		return
// 	})
// 	return
// }
