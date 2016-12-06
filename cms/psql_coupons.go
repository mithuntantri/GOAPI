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
func checkCouponValidity(coupon string, mobile, firstorder bool, amount float64)  (bool, float64){
  var count int8
  var only_first, only_app bool
  var min_amount, discount float64
  db.QueryRow("SELECT count(*) FROM coupons WHERE coupon_id=$1",coupon).Scan(&count)
  if count==0{
    return false, 0
  }
  db.QueryRow("SELECT only_first, only_app, min_amount, discount FROM coupons WHERE coupon_id=$1",coupon).Scan(&only_first, &only_app, &min_amount, &discount)
  if only_app && mobile{
    //Only App
    if only_first && firstorder{
      //Only App and First Order
      if amount<min_amount{
        return false, 0
      }else if discount == 10{
        return true,amount * 10 / 100
      }else if discount == 100{
        return true, 100
      }
    }else if !only_first{
      //Only App and need not to be firstorder
      if amount<min_amount{
        return false, 0
      }else if discount == 10{
        return true,amount * 10 /100
      }else if discount == 100{
        return true, 100
      }
    }
  }else if !only_app{
    if only_first && firstorder{
      //Not only app and First order only
      if amount<min_amount{
        return false, 0
      }else if discount == 10{
        return true,amount * 10 /100
      }else if discount == 100{
        return true, 100
      }
    }else if !only_first{
      //Not only App and need not to be first order
      if amount<min_amount{
        return false, 0
      }else if discount == 10{
        return true,amount * 10 /100
      }else if discount == 100{
        return true, 100
      }
    }
  }
  return false, 0
}
