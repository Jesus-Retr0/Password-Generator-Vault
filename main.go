package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Password Vault")
    fmt.Println("1. Generate Password")
    fmt.Println("2. Decrypt Vault")
    fmt.Println("3. Save Password to Vault")
    fmt.Print("Choose an option: ")
    choice, _ := reader.ReadString('\n')
    choice = strings.TrimSpace(choice)

    switch choice {
    case "1":
        fmt.Print("Enter password length: ")
        lengthStr, _ := reader.ReadString('\n')
        lengthStr = strings.TrimSpace(lengthStr)
        length, err := strconv.Atoi(lengthStr)
        if err != nil || length <= 0 {
            fmt.Println("Invalid length.")
            return
        }
        password, err := GeneratePassword(length)
        if err != nil {
            fmt.Println("Error generating password:", err)
            return
        }
        fmt.Println("Generated password:", password)
    case "2":
        fmt.Print("Enter master password: ")
        masterPassword, _ := reader.ReadString('\n')
        masterPassword = strings.TrimSpace(masterPassword)
        decrypted, err := decryptVault("storage.json", masterPassword)
        if err != nil {
            fmt.Println("Error decrypting vault:", err)
            return
        }
        var vault VaultData
        if len(decrypted) > 0 {
            json.Unmarshal(decrypted, &vault)
        }
        if len(vault) == 0 {
            fmt.Println("Vault is empty.")
            return
        }
        fmt.Println("Vault contents:")
        for service, password := range vault {
            fmt.Printf("%s: %s\n", service, password)
        }
    case "3":
        fmt.Print("Enter master password: ")
        masterPassword, _ := reader.ReadString('\n')
        masterPassword = strings.TrimSpace(masterPassword)

        // Try to decrypt existing vault
        var vault VaultData
        decrypted, err := decryptVault("storage.json", masterPassword)
        if err == nil && len(decrypted) > 0 {
            json.Unmarshal(decrypted, &vault)
        } else {
            vault = make(VaultData)
        }

        fmt.Print("Enter service name: ")
        service, _ := reader.ReadString('\n')
        service = strings.TrimSpace(service)
        fmt.Print("Enter password to save: ")
        passwordToSave, _ := reader.ReadString('\n')
        passwordToSave = strings.TrimSpace(passwordToSave)

        vault[service] = passwordToSave
        vaultBytes, _ := json.Marshal(vault)
        err = encryptVault("storage.json", masterPassword, vaultBytes)
        if err != nil {
            fmt.Println("Error saving to vault:", err)
            return
        }
        fmt.Println("Password saved to vault.")
    default:
        fmt.Println("Invalid option.")
    }
}