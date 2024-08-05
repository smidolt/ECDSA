# ECDSA Key Generation and Message Signing

## Overview

This project demonstrates how to generate ECDSA-P256 keys, sign messages, and verify signatures using the Go programming language. The implementation includes:

1. Generating ECDSA-P256 keys and saving them in PEM format.
2. Signing messages with the private key and returning the signature in Base64URL format.
3. Verifying the signature of messages using the public key.

## Project Structure

- `main.go`: The entry point of the application. This file tests the generation of keys, signing of a message, and verification of the signature.
- `internal/crypto/keys.go`: Contains the function to generate ECDSA-P256 keys and save them in PEM format.
- `internal/crypto/sign.go`: Contains the function to sign a message using the ECDSA private key.
- `internal/crypto/verify.go`: Contains the function to verify the signature of a message using the ECDSA public key.

## Usage

1. Ensure you have Go installed on your machine.
2. Create a directory for the project and navigate into it.
3. Create the necessary directory structure and files as described above.
4. Run `go mod init my_project` to initialize the Go module.
5. Run `go run main.go` to execute the application and see the results.

## Functions

### `CreateECDSAKey`

Generates an ECDSA-P256 key pair and saves the private and public keys in PEM format.

### `SignMessage`

Signs a message using the ECDSA private key and returns the signature in Base64URL format.

### `VerifySignature`

Verifies the signature of a message using the ECDSA public key.

## Example

The example in `main.go` demonstrates the following steps:

1. Check for the existence of the `keys` directory and create it if it does not exist.
2. Generate ECDSA-P256 keys and save them as `private.pem` and `public.pem`.
3. Sign a message with the private key.
4. Verify the signature with the public key.

The output will indicate whether the keys were created successfully, display the generated signature, and confirm whether the signature is valid.
