package main

import (
  "github.com/gin-gonic/gin"
)
type Categories struct{
  Category string `json:"category_name"`
  Descriptions []Description  `json:"description"`
}
type Description struct{
  Header string `json:"heading"`
  Description string `json:"description"`
  Link string `json:"link"`
  Img string `json:"img"`
}
func getCategoriesHandler(c *gin.Context)  {
    var response struct{
      Category []Categories `json:"categories"`
    }
    var category1, category2 Categories

    var desc1 = Description{
      Header : "TRENDING NOW",
      Description : "Custom Shirts @ Rs. 699 Only",
      Link : "SHOP",
      Img : "/img/custom_shirts.jpg",
    }
    var desc2 = Description{
      Header : "EXCLUSIVE OFFER",
      Description : "Refer & Earn Zigfo Credits",
      Link : "SIGN UP",
      Img : "/img/credits.png",
    }
    var desc3 = Description{
      Header : "SHOP ON THE GO",
      Description : "Additional 10% off on app",
      Link : "DOWNLOAD NOW",
      Img : "/img/download_app.png",
    }
    var desc4 = Description{
      Header : "TOP COUPONS",
      Description : "Flat Rs.100 Off on First Order",
      Link : "BUY NOW",
      Img : "/img/top_coupons.jpg",
    }
    var desc5 = Description{
      Header : "TAILOR MADE CUSTOM SHIRTS",
      Description : "Customize & Get Home Delivered Men's Formal",
      Link : "ORDER NOW",
      Img : "/img/slider1.jpg",
    }
    var desc6 = Description{
      Header : "TAILOR MADE CUSTOM SHIRTS",
      Description : "Customize & Get Home Delivered Men's Formal",
      Link : "ORDER NOW",
      Img : "/img/slider2.jpg",
    }
    var desc7 = Description{
      Header : "TAILOR MADE CUSTOM SHIRTS",
      Description : "Customize & Get Home Delivered Men's Formal",
      Link : "ORDER NOW",
      Img : "/img/slider3.jpg",
    }
    var desc8 = Description{
      Header : "TAILOR MADE CUSTOM SHIRTS",
      Description : "Customize & Get Home Delivered Men's Formal",
      Link : "ORDER NOW",
      Img : "/img/slider4.jpg",
    }

    category1.Descriptions = make([]Description, 0)
    category2.Descriptions = make([]Description, 0)
    category1.Descriptions = append(category1.Descriptions, desc1, desc2, desc3, desc4)
    category2.Descriptions = append(category2.Descriptions, desc5, desc6, desc7, desc8)

    category1.Category = "Top Categories"
    category2.Category = "All Categories"

    response.Category = make([]Categories, 0)
    response.Category = append(response.Category, category1, category2)

    c.JSON(200, gin.H{
      "status" :"success",
      "data" : response,
    })
}
