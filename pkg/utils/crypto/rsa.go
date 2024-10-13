package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func GenerateRSAKeyPair(bitSize int) (publicKeyStr string, privateKey *rsa.PrivateKey, err error) {
	// Private Key
	privateKey, err = rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return publicKeyStr, privateKey, err
	}

	// Public Key
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return publicKeyStr, privateKey, err
	}

	publicKeyPem := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	publicKeyStr = string(pem.EncodeToMemory(publicKeyPem))
	
	return publicKeyStr, privateKey, nil
}