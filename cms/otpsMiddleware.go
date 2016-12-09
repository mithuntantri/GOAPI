package main

import (
  "bytes"
  "fmt"
  "encoding/json"
  "io/ioutil"
  "net/http"
)
func callnewOTP(email_id, mobileno, reqType string) bool{
    url := "http://127.0.0.1:2000/otp/create"
    request := fmt.Sprintf("{\"email_id\":\"%s\",\"mobileno\":\"%s\", \"request_type\":\"%s\"}",
        email_id, mobileno, reqType)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(request)))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    var response struct{
      Status string `json:"status"`
      IsExists bool `json:"is_exists"`
      ValidRequest bool `json:"valid_request"`
    }
    body, _ := ioutil.ReadAll(resp.Body)
    json.Unmarshal([]byte(string(body)), &response)
    return response.Status == "success"
}
func callverOTP(mobileno, otp, reqType string) (bool,bool){
  url := "http://127.0.0.1:2000/otp/verify"
  request := fmt.Sprintf("{\"mobileno\":\"%s\", \"otp\":\"%s\", \"request_type\":\"%s\"}",
      mobileno, otp, reqType)

  req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(request)))
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      panic(err)
  }
  defer resp.Body.Close()

  var response struct{
    Status string `json:"status"`
    Verified bool `json:"verified"`
    Blocked bool `json:"is_blocked"`
    ValidNumber bool `json:"valid_number"`
    ValidRequest bool `json:"valid_request"`
  }
  body, _ := ioutil.ReadAll(resp.Body)
  json.Unmarshal([]byte(string(body)), &response)
  return response.Blocked, response.Verified
}
func callresOTP(mobileno, reqType string) (bool){
  url := "http://127.0.0.1:2000/otp/resend"
  request := fmt.Sprintf("{\"mobileno\":\"%s\", \"request_type\":\"%s\"}",
      mobileno, reqType)

  req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(request)))
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      panic(err)
  }
  defer resp.Body.Close()

  var response struct{
    Status string `json:"status"`
    Blocked bool `json:"is_blocked"`
    ValidNumber bool `json:"valid_number"`
    ValidRequest bool `json:"valid_request"`
  }
  body, _ := ioutil.ReadAll(resp.Body)
  json.Unmarshal([]byte(string(body)), &response)
  return response.Blocked
}
