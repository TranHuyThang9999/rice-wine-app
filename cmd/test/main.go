package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// Hàm để tạo và lưu private key
func generateRSAKeys() error {
	// Tạo private key RSA 1024-bit
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return fmt.Errorf("lỗi tạo private key: %v", err)
	}

	// Lưu private key vào file
	privateKeyFile, err := os.Create("private_key.pem")
	if err != nil {
		return fmt.Errorf("lỗi tạo file private_key.pem: %v", err)
	}
	defer privateKeyFile.Close()

	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
		return fmt.Errorf("lỗi ghi private key vào file: %v", err)
	}
	fmt.Println("Private key được lưu tại private_key.pem")

	// Lưu public key vào file
	publicKey := &privateKey.PublicKey
	publicKeyFile, err := os.Create("public_key.pem")
	if err != nil {
		return fmt.Errorf("lỗi tạo file public_key.pem: %v", err)
	}
	defer publicKeyFile.Close()

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return fmt.Errorf("lỗi mã hóa public key: %v", err)
	}
	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	if err := pem.Encode(publicKeyFile, publicKeyPEM); err != nil {
		return fmt.Errorf("lỗi ghi public key vào file: %v", err)
	}
	fmt.Println("Public key được lưu tại public_key.pem")

	return nil
}

func main() {
	if err := generateRSAKeys(); err != nil {
		fmt.Printf("Lỗi: %v\n", err)
	} else {
		fmt.Println("Cặp khóa RSA đã được tạo thành công!")
	}
}
