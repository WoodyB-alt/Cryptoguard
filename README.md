# Cryptoguard 🔐

**Cryptoguard** is a secure AES-256 encryption and decryption CLI tool written in Go.

Use it to encrypt and decrypt **text, files, folders, or even hide data inside PNG images** using password-based encryption with modern key derivation and **authenticated encryption (AES-GCM)**. Ideal for local secrets management, practice, or lightweight secure storage.

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
- 📦 `--zip` flag to compress + encrypt folders
- 🔓 `decrypt-zip` support to auto-decrypt and extract encrypted zip archives
- 🧼 Optional `--delete-original` flag to securely remove plaintext
- 🖼️ **Steganography Mode**: hide encrypted data in PNGs (`steg-embed` / `steg-extract`)
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
- secret.txt: The input file to encrypt
- secret.enc: The AES-GCM + base64 encrypted output file
- Optional: Add --delete-original to remove secret.txt after encryption

### 🔓 Decrypt a file message
```bash
./cryptoguard decrypt-file -p "mypassword" secret.enc recovered.txt
```
- secret.enc: The encrypted file
- recovered.txt: The output file after decryption

### 📂 Encrypt an entire folder (recursive)
```bash
./cryptoguard encrypt-folder -p "mypassword" --recursive --delete-original ./mydocs ./encrypted_docs
```
- Recursively encrypts all files from ./mydocs into ./encrypted_docs
- Retains directory structure and appends .enc to each file
- Shows a progress bar while encrypting
- Add --delete-original to remove the source files after encryption

### 📁 Decrypt an entire folder (recursive)
```bash
./cryptoguard decrypt-folder -p "mypassword" --recursive ./encrypted_docs ./mydocs_copy
```
- Decrypts all .enc files in ./encrypted_docs into ./mydocs_copy
- Preserves original folder structure

### 📦 Zip and Encrypt a Folder
```bash
./cryptoguard encrypt-folder -p "mypassword" --zip --delete-original ./mydocs ./backups
```
- Compresses ./mydocs into a .zip file
- Encrypts the .zip to .zip.enc
- Add --delete-original to remove both the folder and .zip file after encryption

### 📦 Zip and Decrypt a Folder
```bash
./cryptoguard decrypt-zip -p "mypassword" --delete-original backups/mydocs.zip.enc ./restored_docs
```
- Decrypts the .zip.enc file
- Extracts the resulting zip archive into ./restored_docs
- Add --delete-original to remove both .zip and .zip.enc after extraction

### 🖼️ Embed Encrypted Text or File Inside a PNG
```bash
./cryptoguard steg-embed -p "mypassword" -in secret.txt -img carrier.png -out hidden.png
```
- Encrypts the file or text content
- Embeds it into carrier.png using least significant bits (LSBs)
- Saves the modified image as hidden.png

### 🖼️ Extract and Decrypt from a PNG
```bash
./cryptoguard steg-extract -p "mypassword" -img hidden.png --out decrypted.txt
```
- Extracts hidden encrypted message
- Decrypts it using your password
- Optionally writes result to decrypted.txt

## 📌 Notes

- AES-GCM provides authenticated encryption — tampered or incorrectly decrypted data will fail securely

- PBKDF2 is used for password-based key derivation with:

- Random 128-bit salt per encryption

- 100,000 SHA-256 iterations

- A 96-bit random nonce (IV) is generated for every encryption

- Encrypted output is Base64 encoded for safe storage and transport

- Use --delete-original to automatically wipe plaintext files/folders post-encryption

- PNG steganography is suitable for hidden, encrypted transmission

- Always use strong, unique passwords — encryption strength depends on it



## 🧑‍💻 Author
- Blake Wood — 2024 Cybersecurity CLI Project
