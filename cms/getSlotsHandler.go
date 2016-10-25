package main

import (
  "fmt"
  "time"
  "github.com/gin-gonic/gin"
)

func Bod(t time.Time) time.Time {
    year, month, day := t.Date()
    fmt.Println(year, month, day)
    return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

var dayBucket int

type Appointments struct{
  AppointmentID string `json:"appointment_id"`
  AppointmentDay int64 `json:appointment_day`
  SlotID string `json:"slot_id"`
  Username string `json:"username"`
}
type Slot struct{
  SlotID string `json:"slot_id"`
  SlotName string `json:"slot_name"`
  NumSlots int `json:"num_slots"`
}
type Slots struct{
  Day int64 `json:"day"`
  Slots []Slot `json:"slots"`
}
type AllSlots struct{
  Data []Slots `json:"available"`
}
func getSlotsHandler(c *gin.Context)  {
  var data AllSlots
  data.Data = make([]Slots, 0)
  india, err := time.LoadLocation("Asia/Kolkata")
  if err != nil{
    checkErr(err)
  }
  now := time.Now().In(india)
  for i := 0; i<7; i++ {
    nextday := time.Hour * 24 * time.Duration(i)
    current_day := now.Add(nextday)
    day := Bod(current_day).Unix()
    fmt.Println("Current time:", day)
    Slots := getSlots(day)
    data.Data = append(data.Data, Slots)
  }
  c.JSON(200, gin.H{
    "status" : "success",
    "data" : data,
  })
}
func getNumberPilots()  int{
  var count int
  db.QueryRow("SELECT COUNT(*) FROM username_map").Scan(&count)
  return count
}
func getSlots(day int64) Slots{
  var Slots Slots
  numberofPilots := getNumberPilots()
  Slots.Day = day
  Slots.Slots = make([]Slot, 0)
  rows,_ := db.Query("SELECT * from slots")
  for rows.Next(){
    var slot Slot
    if err := rows.Scan(&slot.SlotID, &slot.SlotName); err !=nil{
      checkErr(err)
    }else{
      slotsRemaining := getSlotsRemaining(day, slot.SlotID)
      fmt.Println("Number of pilots:",numberofPilots)
      fmt.Println("Number of slots available:",slotsRemaining)
      slot.NumSlots = numberofPilots - slotsRemaining
      Slots.Slots = append(Slots.Slots, slot)
    }
  }
  return Slots
}
func getSlotsRemaining(day int64, slot_id string) int{
  fmt.Println("Getting:",day, slot_id)
  rows,err := db.Query("SELECT * FROM appointments WHERE slot_id=$1", slot_id)
  if err != nil{
    checkErr(err)
  }
  count := 0
  for rows.Next(){
    var appointment Appointments
    rows.Scan(&appointment.AppointmentID, &appointment.AppointmentDay, &appointment.SlotID, &appointment.Username)
    fmt.Println("Comparing:", appointment.AppointmentDay, day)
    if appointment.AppointmentDay == day{
      count++
    }
  }
  return count
}
