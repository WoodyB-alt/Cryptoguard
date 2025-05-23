# Cryptoguard 🔐

**Cryptoguard** is a secure AES-256 encryption and decryption CLI tool written in Go.

Use it to encrypt and decrypt **text, files, and folders** using password-based encryption with modern key derivation. Ideal for local secrets management, practice, or lightweight secure storage.

---

## ✨ Features

- 🔐 AES-256 symmetric encryption (CFB mode)
- 🧂 PBKDF2 password-based key derivation with SHA-256
- 🔑 Per-encryption random salt generation
- 🔁 100,000 iteration key stretching for brute-force resistance
- 🧾 Base64 encoding for safe output and transport
- 💻 Easy-to-use CLI interface
- 📂 Secure file encryption and decryption support
- 📂 Recursive folder encryption and decryption with `--recursive`
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

### 📂 Encrypt an entire folder (recursive)
```bash
./cryptoguard encrypt-folder -p "mypassword" --recursive ./mydocs ./encrypted_docs
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

- 🔐 PBKDF2 is used for password-based key derivation with:

- Random 128-bit salt per encryption

- 100,000 iterations

- SHA-256 as the hashing algorithm

- 🔁 IV (Initialization Vector) is randomly generated for each file and stored alongside the ciphertext.

- 🧾 Encrypted output is always Base64 encoded to allow safe text/file handling.

- 🚨 Use strong passwords — the security of symmetric encryption depends on it.



## 🧑‍💻 Author
- Blake Wood — 2024 Cybersecurity CLI Project
