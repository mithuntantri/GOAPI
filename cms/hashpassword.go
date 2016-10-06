package main

import (
    "fmt"
    "io"
    "encoding/base64"
    "crypto/sha1"
)

const saltSize = 16

func generateSalt(secret []byte) []byte {
    buf := make([]byte, saltSize, saltSize+sha1.Size)
    hash := sha1.New()
    hash.Write(buf)
    hash.Write(secret)
    return hash.Sum(buf)
}
func getHashedPassword(password string) string{
    salt := generateSalt([]byte(password))
    fmt.Println("salt",salt)
    fmt.Println("salt",password)
    combination := string(salt) + string(password)
    passwordHash := sha1.New()
    io.WriteString(passwordHash, combination)
    pass := base64.URLEncoding.EncodeToString(passwordHash.Sum(nil))
    fmt.Printf(pass)
    return pass
  }
