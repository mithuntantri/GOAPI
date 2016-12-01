package main

import (
  "strconv"
  "github.com/gin-gonic/gin"
)
func getOptionsHandler(c *gin.Context)  {
  var request struct{
    Hash string `form:"hash" binding:"required"`
  }
  if c.Bind(&request) == nil{
    var optionsList, optionsCount = makeOptionsList()

    var common_set Set
    var initdata initData
    initdata.Hash = request.Hash
    initdata.TotalPrice = "699.00"
    initdata.Gender = "M"
    initdata.Data = make([]Set, 0)
    for i:=0;i<=9;i++{
      common_set.Key = strconv.Itoa(i+1)
      common_set.Name = optionsList[i]
      common_set.Options = make([]Options, 0)
      var optionsCount = optionsCount[i]
      for j:=1; j<=optionsCount; j++{
        var option Options
        option = fetchOptions(i + 1, j)
        option.Selected, initdata.Favorites = getDesignHash(initdata.Hash, i+1, j)
        common_set.Options = append(common_set.Options, option)
      }
      initdata.Data = append(initdata.Data, common_set)
    }
    // insertNewHash(initdata.Hash, request.Mobileno)
    c.JSON(200, gin.H{
      "status": "success",
      "data": initdata,
    })
  }
}
