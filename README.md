# Password-Generator-Vault

A simple, secure, and offline password manager written in Go.  
Generate strong passwords and store them in an encrypted vault file on your computer.  
**No internet required**—your secrets stay private.

---

## Features

- **Generate strong random passwords** with customizable length.
- **Store multiple passwords** for different services/accounts.
- **All data is encrypted** using AES-GCM with a key derived from your master password (Argon2).
- **No plaintext storage**—only you can decrypt your vault.
- **No cloud, no tracking, no dependencies** beyond Go standard library and `golang.org/x/crypto`.

---

## Usage

### 1. Build and Run

```sh
go run .
```
or
```sh
go build
./Password-Generator-Vault
```

### 2. Menu Options

- **1. Generate Password**  
  Enter a desired length and get a strong, random password.

- **2. Decrypt Vault**  
  Enter your master password to view all stored service/password pairs.

- **3. Save Password to Vault**  
  Enter your master password, service name, and password to add or update an entry in your encrypted vault.

---

## How It Works

- The vault is stored in `storage.json` and is always encrypted.
- The **master password** is never stored—it's used to derive the encryption key each time.
- Only someone with the correct master password can decrypt and view the vault contents.

---

## Security Notes

- **Keep your master password strong and secret.**
- **Do not share or upload your `storage.json` file.**
- If you lose your master password, you cannot recover your vault.
- **Disclaimer**: This is a personal project to explore local encryption and password management using Go. It is not intended as a production-ready security tool. Please use it for educational purposes and local testing only.

---

## Requirements

- Go 1.17 or newer

---

## License

Licensed under the [MIT License](LICENSE)

---

*Created for privacy-focused users who want simple, offline password management!*