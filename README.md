# Cryptoguard 🔐

A simple AES-256 encryption and decryption CLI tool built in Go.

## Features
- AES-256 symmetric encryption using CFB mode
- SHA-256 derived password key
- Base64 encoding for output
- CLI interface

## Usage

### Encrypt a message:
```bash
cryptoguard encrypt "Hello World" -p "mypassword"
