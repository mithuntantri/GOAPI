package main

import (
  _ "github.com/lib/pq"
)

func fetchCoupons() Coupon{
  var all_coupons Coupon
  all_coupons.AllCoupons = make([]coupons, 0)
  rows,err := db.Query("SELECT coupon_id, description, expiry, only_new, only_first, only_app  FROM coupons")
  if err != nil{
    checkErr(err)
  }
  for rows.Next(){
    var row coupons
    if err := rows.Scan(&row.CouponID, &row.Description, &row.Expiry, &row.OnlyNew, &row.OnlyFirst, &row.OnlyApp); err != nil{
      checkErr(nil)
    }else{
      all_coupons.AllCoupons = append(all_coupons.AllCoupons, row)
    }
  }
  return all_coupons
}
