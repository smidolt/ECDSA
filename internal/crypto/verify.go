package crypto

import (
    "crypto/ecdsa"
    "crypto/sha256"
    "crypto/x509"
    "encoding/base64"
    "encoding/pem"
    "errors"
    "math/big"
    "os"
)

// VerifySignature verifies the signature of a message using the ECDSA public key
func VerifySignature(publicKeyPath string, message []byte, signatureBase64URL string) (bool, error) {
    // Read the public key from the PEM file
    publicKeyBytes, err := os.ReadFile(publicKeyPath)
    if err != nil {
        return false, err
    }
    
    block, _ := pem.Decode(publicKeyBytes)
    if block == nil || block.Type != "PUBLIC KEY" {
        return false, errors.New("failed to decode PEM block containing public key")
    }
    
    pub, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return false, err
    }
    
    publicKey, ok := pub.(*ecdsa.PublicKey)
    if !ok {
        return false, errors.New("not ECDSA public key")
    }
    
    // Hash the message
    hash := sha256.Sum256(message)
    
    // Decode the signature from Base64URL
    signature, err := base64.URLEncoding.DecodeString(signatureBase64URL)
    if err != nil {
        return false, err
    }
    
    // Split the signature into r and s
    r := big.NewInt(0).SetBytes(signature[:len(signature)/2])
    s := big.NewInt(0).SetBytes(signature[len(signature)/2:])
    
    // Verify the signature
    valid := ecdsa.Verify(publicKey, hash[:], r, s)
    
    return valid, nil
}
