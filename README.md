# Cryptoguard 🔐

**Cryptoguard** is a secure AES-256 encryption and decryption CLI tool written in Go.

Use it to encrypt and decrypt **text, files, and folders** using password-based encryption with modern key derivation and **authenticated encryption (AES-GCM)**. Ideal for local secrets management, practice, or lightweight secure storage.

---

## ✨ Features

- 🔐 AES-256 authenticated encryption using AES-GCM mode
- 🧂 PBKDF2 password-based key derivation with SHA-256
- 🔑 Per-encryption random salt generation (128-bit)
- 🔁 100,000 iteration key stretching for brute-force resistance
- 🧾 Base64 encoding for safe output and transport
- 💻 Easy-to-use CLI interface
- 📁 Secure file encryption and decryption
- 📂 Recursive folder encryption and decryption with `--recursive`
- 🧼 Optional `--delete-original` flag to securely remove plaintext
- 📊 Progress bar during folder encryption
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
- secret.txt: The input file to encrypt.
- secret.enc: The encrypted output file (IV + base64 encoded).

### 🔓 Decrypt a file message
```bash
./cryptoguard encrypt-file -p "mypassword" secret.enc secret.txt
```
- secret.txt: The input file to Decrypt.
- secret.enc: The Decrypted output file (IV + base64 encoded).

### 📂 Encrypt an entire folder (recursive)
```bash
./cryptoguard encrypt-folder -p "mypassword" --recursive --delete-original ./mydocs ./encrypted_docs
```
- Recursively encrypts all files from ./mydocs to ./encrypted_docs
- Retains directory structure and appends .enc to each file

### 📁 Decrypt an entire folder (recursive)
```bash
./cryptoguard decrypt-folder -p "mypassword" --recursive ./encrypted_docs ./mydocs_copy
```
- Decrypts all .enc files in ./encrypted_docs into ./mydocs_copy
- Preserves original folder structure


## 📌 Notes

- AES-GCM provides authenticated encryption — tampered or incorrectly decrypted data will fail securely

- PBKDF2 is used for password-based key derivation with:

- Random 128-bit salt per encryption

- 100,000 SHA-256 iterations

- A 96-bit random nonce (IV) is generated for every encryption

- Encrypted output is Base64 encoded for safe storage and transport

- Use --delete-original to automatically wipe plaintext files/folders post-encryption

- Always use strong, unique passwords — encryption strength depends on it



## 🧑‍💻 Author
- Blake Wood — 2024 Cybersecurity CLI Project
