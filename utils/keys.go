package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"golang.org/x/crypto/ssh"
)

// This helper method will create a private and public
// keypair under the following directory:
// 		~/.ssh/ssh-helper/{{name}}/
// If neither ssh-helper nor the directory exists, they
// will be created for you.
//
// Returns the path to the private key and an error.
func NewKey(name string) (string, error) {
	if name == "" {
		return "", errors.New("error: key name is empty.")
	}

	// default path is ~/.ssh/ssh-helper/<name>/
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// ensure ssh-helper directory exists
	sshpath := path.Join(homedir, ".ssh", "ssh-helper")
	if exists := FileExists(sshpath); !exists {
		err := MkDir(sshpath)
		if err != nil {
			return "", err
		}
	}

	// ensure config directory for host in ssh-helper
	// directory exists
	bpath := path.Join(sshpath, name)
	exists := FileExists(bpath)
	if !exists {
		fmt.Println("Making directory... ", bpath)
		err := MkDirP(bpath)
		if err != nil {
			return "", err
		}
	}

	// destination paths for keys
	privPath := path.Join(bpath, name)
	pubPath := path.Join(bpath, name+".pub")
	bits := 4096

	// skip if the keys already exist
	if FileExists(privPath) && FileExists(pubPath) {
		fmt.Println("Skipping: keys already exist.")
		return bpath, nil
	}

	// first: create the private key
	privkey, err := mkPrivKey(bits)
	if err != nil {
		log.Fatal(err.Error())
	}

	// second: generate public key from the private one
	pubkeyBytes, err := mkPubKey(&privkey.PublicKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	// encodes privkey: RSA -> PEM format
	privkeyBytes := pem.EncodeToMemory(&pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PrivateKey(privkey),
	})

	// save private key to file
	fmt.Println("Writing private key to: ", privPath)
	err = saveKey(privkeyBytes, privPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Done.")

	// save public key to file
	fmt.Println("Writing public key to: ", pubPath)
	err = saveKey([]byte(pubkeyBytes), pubPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Done.")

	return privPath, nil
}

// generate a private key given a number of bits
func mkPrivKey(bits int) (*rsa.PrivateKey, error) {
	// try to gen key
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}

	// make sure it is correct
	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// take a rsa.PublicKey and return bytes suitable for writing to .pub file
// returns in the format "ssh-rsa ..."
func mkPubKey(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)

	return pubKeyBytes, nil
}

// save key to file path
func saveKey(keyBytes []byte, saveFileTo string) error {
	err := ioutil.WriteFile(saveFileTo, keyBytes, 0600)
	if err != nil {
		return err
	}
	return nil
}
