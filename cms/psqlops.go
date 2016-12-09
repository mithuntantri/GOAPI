package main

import (
  "fmt"
  "strings"
  "database/sql"
  _ "github.com/lib/pq"
)

var db *sql.DB

func connectPSQL() {
    db, _ = sql.Open("postgres", "user=postgres dbname=mithun sslmode=disable password=postgres")
}

func QuoteIdentifier(name string) string {
    end := strings.IndexRune(name, 0)
    if end > -1 {
        name = name[:end]
    }
    return `"` + strings.Replace(name, `"`, `""`, -1) + `"`
}
func getMobileNumberFromEmail(email_id string) string{
  var mobileno string
  db.QueryRow("SELECT mobileno FROM credentials WHERE email_id=$1",email_id).Scan(&mobileno)
  return mobileno
}
func getMobileNumber(id, table_name, field_name string) string{
  var mobileno string
  db.QueryRow(fmt.Sprintf("SELECT mobileno FROM %s WHERE %s=$1",
  QuoteIdentifier(table_name),QuoteIdentifier(field_name)),id).Scan(&mobileno)
  return mobileno
}
func getwalletIDFromMobile(mobileno string) string{
  var wallet_id string
  db.QueryRow("SELECT wallet_id FROM profile WHERE mobileno=$1",mobileno).Scan(&wallet_id)
  return wallet_id
}
