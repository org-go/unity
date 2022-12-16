package resp

//  OptionsResp
type OptionsResp []*OptionResp
type OptionResp struct {
    Text  string `json:"text"`  // 文本
    Value int32    `json:"value"` // 值
}
