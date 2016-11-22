package main

import (
  "fmt"
  _ "github.com/lib/pq"
)

func checkNewMobileno(mobileno string) bool{
  fmt.Println("Checking if New Mobile Number :", mobileno)
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM credentials WHERE mobileno=$1",mobileno).Scan(&count)
  if count == 0 {
    return true
  }else{
    return false
  }
}

func checkNewEmailID(email_id string) bool{
  fmt.Println("Checking if New Email ID :", email_id)
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM credentials WHERE email_id=$1",email_id).Scan(&count)
  if count == 0 {
    return true
  }else{
    return false
  }
}

func addtoCredentials(mobileno , email_id string, ver1 bool, ver2 bool, client_id , password string) bool{
  fmt.Println("Adding to Credentials", mobileno, email_id, ver1, ver2, client_id, password)
  db.QueryRow("INSERT INTO credentials (mobileno, email_id, verified_mobile, verified_email, client_id, password) VALUES($1,$2,$3,$4,$5,$6);", mobileno, email_id, ver1, ver2, client_id, password)
  return true
}

func updateVerifiedEmail(mobileno string)  bool{
  stmt, err := db.Prepare("UPDATE credentials SET verified_email=$1 WHERE mobileno=$2;")
  checkErr(err)
  _,err = stmt.Exec(true, mobileno)
  if err != nil{
    return false
  }
  return true
}

func checkCredentials(key, client_id, password string, is_mobileno bool) bool{
  var hashedPass string
  var db_clientid string
  var count int8
  if is_mobileno{
    db.QueryRow("SELECT COUNT(*) FROM credentials WHERE mobileno=$1", key).Scan(&count)
    if count == 0{
      return false
    }
    err := db.QueryRow("SELECT password, client_id FROM credentials WHERE mobileno=$1", key).Scan(&hashedPass, &db_clientid)
    checkErr(err)
  }else{
    db.QueryRow("SELECT COUNT(*) FROM credentials WHERE email_id=$1", key).Scan(&count)
    if count == 0{
      return false
    }
    err := db.QueryRow("SELECT password, client_id FROM credentials WHERE email_id=$1", key).Scan(&hashedPass, &db_clientid)
    checkErr(err)
  }
  fmt.Println(client_id, db_clientid)
  fmt.Println(hashedPass, password)
  password_verified := verifyBcrypt(hashedPass, password)
  if password == ""{
    return true
  }
  if client_id == db_clientid && password_verified{
    return true
  }else{
    return false
  }
}

func verifyMobileno(mobileno, client_id string) (bool, bool){
  var db_clientid string
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM credentials WHERE mobileno=$1",mobileno).Scan(&count)
  if count == 0{
    return false, false
  }
  err := db.QueryRow("SELECT client_id FROM credentials WHERE mobileno=$1",
      mobileno).Scan(&db_clientid)
  checkErr(err)
  if client_id == db_clientid{
    return true, true
  }else{
    return false , true
  }
}

// package main
//
// import (
//   "fmt"
//   _ "github.com/lib/pq"
// )
// func checkNewUser(mobileno string) bool{
//   fmt.Println("Checking if New user :", mobileno)
//   var count int8
//   db.QueryRow("SELECT COUNT(*) FROM credentials WHERE mobileno=$1",mobileno).Scan(&count)
//   if count == 0 {
//     return true
//   }else{
//     return false
//   }
// }
//
// func addtoCredentials(mobileno , client_id , password string) bool{
//   fmt.Println("Adding to Credentials", mobileno, client_id, password)
//   db.QueryRow("INSERT INTO credentials (mobileno, client_id, password) VALUES($1,$2,$3);", mobileno, client_id, password)
//   return true
// }
//
// func checkCredentials(mobileno, client_id, password string) bool{
//   var hashedPass string
//   var db_clientid string
//   var count int8
//   db.QueryRow("SELECT COUNT(*) FROM credentials WHERE mobileno=$1", mobileno).Scan(&count)
//   if count == 0{
//     return false
//   }
//   err := db.QueryRow("SELECT password, client_id FROM credentials WHERE mobileno=$1",
//       mobileno).Scan(&hashedPass, &db_clientid)
//   checkErr(err)
//   fmt.Println(client_id, db_clientid)
//   fmt.Println(hashedPass, password)
//   password_verified := verifyBcrypt(hashedPass, password)
//   if password == ""{
//     return true
//   }
//   if client_id == db_clientid && password_verified{
//     return true
//   }else{
//     return false
//   }
// }
//
// func verifyMobileno(mobileno, client_id string) (bool, bool){
//   var db_clientid string
//   var count int8
//   db.QueryRow("SELECT COUNT(*) FROM credentials WHERE mobileno=$1",mobileno).Scan(&count)
//   if count == 0{
//     return false, false
//   }
//   err := db.QueryRow("SELECT client_id FROM credentials WHERE mobileno=$1",
//       mobileno).Scan(&db_clientid)
//   checkErr(err)
//   if client_id == db_clientid{
//     return true, true
//   }else{
//     return false , true
//   }
// }
