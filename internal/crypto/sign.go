package crypto

import (
    "crypto/ecdsa"
    "crypto/sha256"
    "crypto/x509"
    "encoding/base64"
    "encoding/pem"
    "errors"
    "os"
    "crypto/rand"
)

// SignMessage signs a message using the ECDSA private key and returns the signature in Base64URL format
func SignMessage(privateKeyPath string, message []byte) (string, error) {
    // Read the private key from the PEM file
    privateKeyBytes, err := os.ReadFile(privateKeyPath)
    if err != nil {
        return "", err
    }
    
    block, _ := pem.Decode(privateKeyBytes)
    if block == nil || block.Type != "EC PRIVATE KEY" {
        return "", errors.New("failed to decode PEM block containing private key")
    }
    
    privateKey, err := x509.ParseECPrivateKey(block.Bytes)
    if err != nil {
        return "", err
    }
    
    // Hash the message
    hash := sha256.Sum256(message)
    
    // Sign the message
    r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
    if err != nil {
        return "", err
    }
    
    // Encode the signature in Base64URL
    signature := append(r.Bytes(), s.Bytes()...)
    signatureBase64URL := base64.URLEncoding.EncodeToString(signature)
    
    return signatureBase64URL, nil
}
