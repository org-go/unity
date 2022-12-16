package coreSystem

//
//  SignReq
//  @Description:
//
type SignReq struct {
    Sign string `json:"sign"`
}

func (s *SignReq) Init(key string) SignReq {
    s.Sign = key
    return *s
}

//
//  StateResp
//  @Description:
//
type StateResp struct {
    Result  string `json:"result"`
    Message string `json:"message"`
}
