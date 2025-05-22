# Cryptoguard 🔐

**Cryptoguard** is a simple AES-256 encryption and decryption CLI tool written in Go.

Use it to securely encrypt and decrypt both **text and files** with password-based encryption. Perfect for quick local secrets management or practicing secure CLI tooling.

---

## ✨ Features

- 🔐 AES-256 symmetric encryption (CFB mode)
- 🧂 SHA-256 derived password key
- 🔑 Password-based key derivation
- 🧾 Base64 encoding for safe output
- 💻 Easy-to-use CLI interface
- 📂 File encryption and decryption support
- 📦 Clean and modular Go code structure

---

## 🛠️ Build

```bash
go build -o cryptoguard
```
## 📦 Usage

### 🔐 Encrypt a text message
```bash
-./cryptoguard encrypt -p "mypassword" "Hello World"
```

### 🔓 Decrypt a text message
```bash
./cryptoguard decrypt -p "mypassword" "base64_ciphertext_here"
```

### 📁 Encrypt a file
```bash
./cryptoguard encrypt-file -p "mypassword" secret.txt secret.enc
```
--secret.txt: The input file to encrypt.
--secret.enc: The encrypted output file (IV + base64 encoded).

### 📁 Decrypt a file
```bash
./cryptoguard decrypt-file -p "mypassword" secret.enc recovered.txt
```
--secret.enc: Encrypted input file.
--recovered.txt: Output file with decrypted content.

## 📌 Notes

- Each file encryption uses a secure, random IV (initialization vector).

- IV is prepended to the encrypted file automatically.

- All encrypted output is Base64 encoded for safe transport/storage.

- Use strong passwords for best results — SHA-256 is fast but not slow-hash protected (upgradeable in future).

## 🧑‍💻 Author
- Blake Wood — 2024 Cybersecurity CLI Project
