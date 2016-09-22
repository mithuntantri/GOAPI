package main

import (
    "fmt"
    "io"
    "os"
    "crypto/rand"
    "encoding/base64"
    "crypto/sha1"
)

const saltSize = 16

func generateSalt(secret []byte) []byte {
    buf := make([]byte, saltSize, saltSize+sha1.Size)
    _, err := io.ReadFull(rand.Reader, buf)
    if err != nil {
            fmt.Printf("random read failed: %v", err)
            os.Exit(1)
    }
    hash := sha1.New()
    hash.Write(buf)
    hash.Write(secret)
    return hash.Sum(buf)
}
func getHashedPassword(password string) string{
    salt := generateSalt([]byte(password))
    combination := string(salt) + string(password)
    passwordHash := sha1.New()
    io.WriteString(passwordHash, combination)
    pass := base64.URLEncoding.EncodeToString(passwordHash.Sum(nil))
    fmt.Printf(pass)
    return pass
}
