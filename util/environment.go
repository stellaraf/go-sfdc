package util

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/stellaraf/go-sfdc/types"
)

func findGoMod(start string) (dir string, err error) {
	err = filepath.Walk(start, func(path string, file fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(file.Name(), "go.mod") {
			dir = filepath.Dir(path)
			return nil
		}
		return nil
	})
	return
}

func findProjectRoot() (root string, err error) {
	start := "."
	for {
		if len(start) > 4 {
			err = fmt.Errorf("failed to locate project root, exceeded depth of 4")
			return
		}
		dir, err := findGoMod(start)
		if err != nil {
			return "", err
		}
		if dir == "" {
			start += "."
			continue
		} else {
			root, err = filepath.Abs(dir)
			if err != nil {
				return "", err
			}
			break
		}
	}
	return
}

func loadDotEnv() (err error) {
	projectRoot, err := findProjectRoot()
	if err != nil {
		return
	}
	envFile := filepath.Join(projectRoot, ".env")
	if _, err := os.Stat(envFile); err == nil {
		err = godotenv.Load(envFile)
		if err != nil {
			return err
		}
	}
	return
}

func LoadEnv() (env types.Environment, err error) {
	err = loadDotEnv()
	if err != nil {
		return
	}
	clientId := os.Getenv("SFDC_CLIENT_ID")
	privateKey := os.Getenv("SFDC_PRIVATE_KEY")
	encryptionPassphrase := os.Getenv("SFDC_ENCRYPTION_PASSPHRASE")
	authURL := os.Getenv("SFDC_AUTH_URL")
	authUsername := os.Getenv("SFDC_AUTH_USERNAME")
	testDataRaw := os.Getenv("SFDC_TEST_DATA")
	var testData types.TestData
	err = json.Unmarshal([]byte(testDataRaw), &testData)
	if err != nil {
		return
	}
	env = types.Environment{
		ClientID:             clientId,
		EncryptionPassphrase: encryptionPassphrase,
		PrivateKey:           privateKey,
		AuthURL:              authURL,
		AuthUsername:         authUsername,
		TestData:             testData,
	}
	return
}
