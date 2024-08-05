package crypto

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/x509"
    "encoding/pem"
    "os"
)

// CreateECDSAKey generates an ECDSA-P256 key and saves it in PEM format
func CreateECDSAKey(privateKeyPath, publicKeyPath string) error {
    privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        return err
    }

    // Save the private key
    privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
    if err != nil {
        return err
    }
    privateKeyPem := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: privateKeyBytes})
    if err := os.WriteFile(privateKeyPath, privateKeyPem, 0600); err != nil {
        return err
    }

    // Save the public key
    publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
    if err != nil {
        return err
    }
    publicKeyPem := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: publicKeyBytes})
    if err := os.WriteFile(publicKeyPath, publicKeyPem, 0644); err != nil {
        return err
    }

    return nil
}
