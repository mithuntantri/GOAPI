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
  Data []Set `json:"data"`
}
func fetchOptions(choice int, option_number int)  Options{
  var option = Options{
    Name : "Long Sleeve",
    Code : "long_sleeve",
    Key : "202",
    Enabled : true,
    Selected : true,
    Price : "",
  }
  return option
}
func getOptionsCount() []string{
  var optionsCount []string = []string{3, 4, 3, 17, 10}
}
func makeOptionsList() []string{
  var optionsList []string = []string{
    "Fit",
    "Fabric",
    "Sleeve",
    "Collar",
    "Cuff",
    "Back Details",
    "Bottom Cut",
    "Button",
    "Button Placket Contrast",
    "Collar Contrast",
    "Contrast Button Hole Thread",
    "Contrast Button Thread",
    "Elbow Patch Contrast",
    "Embroidery Font",
    "Embroidery Thread Color",
    "Fastening Spin",
    "Handkerchief",
    "Inner Cuff Contrast",
    "Inner Collar Contrast",
    "Inner Fastening Contrast",
    "Outer Cuff Contrast",
    "Outer Fastening Contrast",
    "Placket",
    "Pocket Lid",
    "Pocket Type",
    "Pocket Contrast",
    "Pocket Placement",
    "Sleeve Placket Contrast",
    "Tie Fix",
    "Under Collar Contrast"
  }
  return optionsList
}
func initProductHandler(c *gin.Context)  {
    var optionsList = makeOptionsList()
    var optionsCount = getOptionsCount()

    var common_set Set
    var initdata initData
    initdata.Hash, _ = Generate(`[a-Z]{20}`)

    initdata.Data = make([]Set, 0)
    for i:=0;i<=29;i++{
      common_set.Key = strconv.Itoa(i+1)
      common_set.Name = optionsList[i]
      common_set.Options = make([]Options, 0)
      optionsCount := optionsCount[i]
      for j:=1; j<=optionsCount; j++{
        var option Options
        option = fetchOptions(i)
        common_set.Options = append(common_set.Options, option)
      }
      initdata.Data = append(initdata.Data, common_set)
    }

    c.JSON(200, gin.H{
      "status": "success",
      "data": initdata,
    })
}
