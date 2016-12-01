package main

import (
  "strconv"
  "github.com/gin-gonic/gin"
)
type Options struct{
    Name string `json:"name"`
    Code string `json:"code"`
    Key string `json:"key"`
    Enabled bool `json:"enabled"`
    Selected bool `json:"selected"`
    Price string `json:"price"`
}
type Set struct{
    Key string `json:"key"`
    Name string `json:"name"`
    Options []Options `json:"options"`
}
type DesignShirt struct{
    BackDetails []Set `json:"back_details"`
    BottomCut []Set `json:"bottom_cut"`
    Button []Set `json:"button"`
    ButtonPlacketContrast []Set `json:"button_placket_contrast"`
    Collar []Set `json:"collar"`
    CollarContrast []Set `json:"collar_contrast"`
    ContrastButtonHoleThread []Set `json:"contrast_button_hole_thread"`
    ContrastButtonThread []Set `json:"contrast_button_thread"`
    Cuff []Set `json:"cuff"`
    ElbowPatchContrast []Set `json:"elbow_patch_contrast"`
    EmbroideryFont []Set `json:"embroidery_font"`
    EmbroideryThreadColor []Set `json:"embroidery_thread_color"`
    Fabric []Set `json:"fabric"`
    FasteningSpin []Set `json:"fastening_spin"`
    Fit []Set `json:"fit"`
    Handkerchief []Set `json:"handkerchief"`
    InnerCollarContrast []Set `json:"inner_collar_contrast"`
    InnerCuffContrast []Set `json:"inner_cuff_contrast"`
    InnerFasteningContrast []Set `json:"inner_fastening_contrast"`
    OuterCuffContrast []Set `json:"outer_cuff_contrast"`
    OuterFasteningContrast []Set `json:"outer_fastening_contrast"`
    Placket []Set `json:"placket"`
    PocketContrast []Set `json:"pocket_contrast"`
    PocketLid []Set `json:"pocket_lid"`
    PocketPlacement []Set `json:"pocket_placement"`
    PocketType []Set `json:"pocket_type"`
    Sleeve []Set `json:"sleeve"`
    SleevePlacketContrast []Set `json:"sleeve_placket_contrast"`
    TieFix []Set `json:"tie_fix"`
    UnderCollarContrast []Set `json:"under_collar_contrast"`
}
type initData struct{
  Hash string `json:"hash"`
  TotalPrice string `json:"total_price"`
  Favorites bool `json:"favorites"`
  Gender string `json:"gender"`
  CheckedOut bool `json:"cheked_out"`
  Data []Set `json:"all_options"`
}
func makeOptionsList() ([]string,[]int){
  var optionsCount []int = []int{3, 3, 17, 10, 7, 3, 3, 2, 4, 3}
  var optionsList []string = []string{
    "Fit",
    "Sleeve",
    "Collar",
    "Cuff",
    "Placket",
    "Pocket",
    "Pocket Type",
    "Pocket Lid",
    "Back Details",
    "Bottom Cut",
  }
  return optionsList, optionsCount
}
func initProductHandler(c *gin.Context)  {
  var request struct{
    Mobileno string `form:"mobileno"`
  }
  if c.Bind(&request) == nil{
    var optionsList, optionsCount = makeOptionsList()
    var common_set Set
    var initdata initData
    initdata.Hash, _ = Generate(`[a-Z]{20}`)
    initdata.TotalPrice = "699.00"
    initdata.Favorites = false
    initdata.CheckedOut = false
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
        common_set.Options = append(common_set.Options, option)
      }
      initdata.Data = append(initdata.Data, common_set)
    }
    insertNewHash(initdata.Hash, request.Mobileno)
    c.JSON(200, gin.H{
      "status": "success",
      "data": initdata,
    })
  }
}
