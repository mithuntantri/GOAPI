package main

import (
  "fmt"
  _ "github.com/lib/pq"
)

func checkMeasurementID(measurement_id string) bool {
  fmt.Println("Checking for valid measurements id", measurement_id)
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM measurements WHERE measurement_id=$1",measurement_id).Scan(&count)
  if count == 0 {
    return false
  }else{
    return true
  }
}
func createMeasurementsID(measurement_id, mobileno string)  bool{
  is_default := true
  db.QueryRow("INSERT INTO measurements(measurement_id, mobileno, name, units, neck, chest, waist, hip, length, shoulder, sleeve, is_default) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12);", measurement_id, mobileno, "Default Measurements", 0, 0 , 0, 0, 0, 0, 0, 0, is_default)
  return true
}
func getMeasurementsID(mobileno string) string{
  var measurement_id string
  db.QueryRow("SELECT measurement_id FROM profile WHERE mobileno=$1", mobileno).Scan(&measurement_id)
  return measurement_id
}
func getMeasurements(measurement_id string) measurements{
  var m measurements
  db.QueryRow("SELECT measurement_id, mobileno, name, units, neck, chest, waist, hip, length, shoulder, sleeve, is_default FROM measurements WHERE measurement_id=$1",measurement_id).Scan(&m.MeasurementID, &m.Mobileno, &m.Name, &m.Units, &m.Neck, &m.Chest, &m.Waist, &m.Hip, &m.Length, &m.Shoulder, &m.Sleeve, &m.Default)
  return m
}
func updateMeasurements(m measurements) bool{
  stmt, err := db.Prepare("UPDATE measurements SET name=$1,units=$2,neck=$3,chest=$4,waist=$5,hip=$6,length=$7,shoulder=$8,sleeve=$9,is_default=$10 where measurement_id=$11")
  checkErr(err)
  _,err = stmt.Exec(m.Name, m.Units, m.Neck, m.Chest, m.Waist, m.Hip, m.Length, m.Shoulder, m.Sleeve, m.Default, m.MeasurementID)
  if err != nil{
    return false
  }
  return true
}
func deleteMeasurements(measurement_id string) bool {
  stmt, err := db.Prepare("DELETE FROM measurements WHERE measurement_id=$1")
  checkErr(err)
  _,err = stmt.Exec(measurement_id)
  if err != nil{
    return false
  }
  return true
}
func fetchMeasurements(mobileno string) Measurements {
  var all_measurements Measurements
  all_measurements.AllMeasurements = make([]measurements,0)
  rows,err := db.Query("SELECT * from measurements WHERE mobileno=$1", mobileno)
  if err != nil{
    checkErr(err)
  }
  for rows.Next(){
    var m measurements
    if err := rows.Scan(&m.MeasurementID, &m.Mobileno, &m.Name, &m.Units, &m.Neck, &m.Chest, &m.Waist, &m.Hip, &m.Length, &m.Shoulder, &m.Sleeve, &m.Default); err != nil{
      checkErr(err)
    }else{
      all_measurements.AllMeasurements = append(all_measurements.AllMeasurements, m)
    }
  }
  return all_measurements
}
