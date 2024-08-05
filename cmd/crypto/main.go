package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "my_project/internal/crypto"
)

func main() {
    // Check if the keys directory exists
    if _, err := os.Stat("keys"); os.IsNotExist(err) {
        err := os.Mkdir("keys", 0755)
        if err != nil {
            log.Fatalf("Failed to create keys directory: %v", err)
        }
    }

    privateKeyPath := filepath.Join("keys", "private.pem")
    publicKeyPath := filepath.Join("keys", "public.pem")
    
    message := []byte("Hello, World!")

    // Create keys
    err := crypto.CreateECDSAKey(privateKeyPath, publicKeyPath)
    if err != nil {
        log.Fatalf("Failed to create keys: %v", err)
    }
    
    fmt.Println("ECDSA keys created successfully")
    
    // Sign the message
    signature, err := crypto.SignMessage(privateKeyPath, message)
    if err != nil {
        log.Fatalf("Failed to sign message: %v", err)
    }
    
    fmt.Printf("Signature: %s\n", signature)
    
    // Verify the signature
    valid, err := crypto.VerifySignature(publicKeyPath, message, signature)
    if err != nil {
        log.Fatalf("Failed to verify signature: %v", err)
    }
    
    if valid {
        fmt.Println("Signature is valid")
    } else {
        fmt.Println("Signature is invalid")
    }
}
