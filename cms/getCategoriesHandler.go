package main

import (
  "github.com/gin-gonic/gin"
)
type Categories struct{
  Category string `json:"category_name"`
  Descriptions []Description  `json:"description"`
}
type Description struct{
  Name string `json:"name"`
  Img string `json:"img"`
}
func getCategoriesHandler(c *gin.Context)  {
    var response struct{
      Category []Categories `json:"categories"`
    }
    var category1, category2 Categories

    var desc1 = Description{
      Name : "Custom Shirts",
      Img : "/img/custom_shirts.jpg",
    }
    var desc2 = Description{
      Name : "Womenswear",
      Img : "/img/womenswear.jpg",
    }
    var desc3 = Description{
      Name : "Fabrics",
      Img : "/img/fabrics.jpg",
    }
    var desc4 = Description{
      Name : "Men's Fashion",
      Img : "",
    }
    var desc5 = Description{
      Name : "Women's Fashion",
      Img : "",
    }
    var desc6 = Description{
      Name : "Suits",
      Img : "",
    }
    var desc7 = Description{
      Name : "Accessories",
      Img : "",
    }
    category1.Descriptions = make([]Description, 0)
    category2.Descriptions = make([]Description, 0)
    category1.Descriptions = append(category1.Descriptions, desc1, desc2, desc3)
    category2.Descriptions = append(category2.Descriptions, desc4, desc5, desc6, desc7)

    category1.Category = "Top Categories"
    category2.Category = "All Categories"

    response.Category = make([]Categories, 0)
    response.Category = append(response.Category, category1, category2)

    c.JSON(200, gin.H{
      "status" :"success",
      "data" : response,
    })
}
