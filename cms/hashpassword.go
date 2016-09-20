package main

import (
        "fmt"
        "io"
        "os"
        "crypto/rand"
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
  fmt.Printf("Password Hash : %x \n", passwordHash.Sum(nil))
  return string(passwordHash.Sum(nil))
}
// func createHashPass(password string) {
//         // during registration
//         fmt.Println("Password : ", password)
//
//         // generate salt from given password
//         salt := generateSalt(password)
//         fmt.Printf("Salt : %x \n", salt)
//
//         // generate password + salt hash to store into database
//         combination := string(salt) + string(password)
//         passwordHash := sha1.New()
//         io.WriteString(passwordHash, combination)
//         fmt.Printf("Password Hash : %x \n", passwordHash.Sum(nil))
//
//         // later on ...
//         // during login, retrieve passwordHash and salt from database
//
//         // test wrong password
//         wrongPassword := []byte("bye") // this is the password from login page
//         wrongCombination := string(salt) + string(wrongPassword)
//         wrongHash := sha1.New()
//         io.WriteString(wrongHash, wrongCombination)
//         fmt.Printf("%x \n", wrongHash.Sum(nil))
//         fmt.Printf("%x \n", passwordHash.Sum(nil))
//
//         match := bytes.Equal(wrongHash.Sum(nil), passwordHash.Sum(nil))
//         fmt.Printf("Login successful ? : %v\n", match)
//
//         // test correct password
//         correctPassword := []byte("hello") // this is the password from login page
//         correctCombination := string(salt) + string(correctPassword)
//         correctHash := sha1.New()
//         io.WriteString(correctHash, correctCombination)
//         fmt.Printf("%x \n", correctHash.Sum(nil))
//         fmt.Printf("%x \n", passwordHash.Sum(nil))
//
//         match = bytes.Equal(correctHash.Sum(nil), passwordHash.Sum(nil))
//         fmt.Printf("Login successful ? : %v\n", match)
// }
