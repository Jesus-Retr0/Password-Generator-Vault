package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "encoding/json"
    "golang.org/x/crypto/argon2"
    "os"
)

type EncryptedVault struct {
    Salt       string `json:"salt"`
    Nonce      string `json:"nonce"`
    Tag        string `json:"tag"`
    Ciphertext string `json:"ciphertext"`
}

// VaultData is a map of service name to password
type VaultData map[string]string

// deriveKey generates a key from the password and salt using Argon2
func deriveKey(password, salt []byte) []byte {
    return argon2.IDKey(password, salt, 1, 64*1024, 4, 32)
}

// decryptVault decrypts the vault data using AES-GCM and returns the plain data
func decryptVault(filename, masterPassword string) ([]byte, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    var vault EncryptedVault
    if err := json.Unmarshal(data, &vault); err != nil {
        return nil, err
    }

    salt, err := base64.StdEncoding.DecodeString(vault.Salt)
    if err != nil {
        return nil, err
    }
    nonce, err := base64.StdEncoding.DecodeString(vault.Nonce)
    if err != nil {
        return nil, err
    }
    tag, err := base64.StdEncoding.DecodeString(vault.Tag)
    if err != nil {
        return nil, err
    }
    ciphertext, err := base64.StdEncoding.DecodeString(vault.Ciphertext)
    if err != nil {
        return nil, err
    }

    key := deriveKey([]byte(masterPassword), salt)

    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    fullCiphertext := append(ciphertext, tag...)
    plaintext, err := aesgcm.Open(nil, nonce, fullCiphertext, nil)
    if err != nil {
        return nil, err
    }
    return plaintext, nil
}

// encryptVault encrypts the vault data using AES-GCM and saves it to file
func encryptVault(filename, masterPassword string, plaintext []byte) error {
    salt := make([]byte, 16)
    if _, err := rand.Read(salt); err != nil {
        return err
    }
    key := deriveKey([]byte(masterPassword), salt)

    block, err := aes.NewCipher(key)
    if err != nil {
        return err
    }
    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
        return err
    }
    nonce := make([]byte, aesgcm.NonceSize())
    if _, err := rand.Read(nonce); err != nil {
        return err
    }

    ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
    tag := ciphertext[len(ciphertext)-aesgcm.Overhead():]
    ciphertext = ciphertext[:len(ciphertext)-aesgcm.Overhead()]

    encrypted := EncryptedVault{
        Salt:       base64.StdEncoding.EncodeToString(salt),
        Nonce:      base64.StdEncoding.EncodeToString(nonce),
        Tag:        base64.StdEncoding.EncodeToString(tag),
        Ciphertext: base64.StdEncoding.EncodeToString(ciphertext),
    }
    data, err := json.MarshalIndent(encrypted, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(filename, data, 0600)
}