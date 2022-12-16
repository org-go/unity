package coreSystem

import (
    "encoding/json"
    "fmt"
)


/**
 * $(decode)
 * @Description:
 * @receiver self
 * @param str
 * @param body
 * @return help
 */
func (self help) decode(str string, body interface{}) help {
    _ = json.Unmarshal([]byte(str), &body)
    return self
}


/**
 * $(error)
 * @Description:
 * @receiver self
 * @param rsp
 * @return err
 */
func (self help) error (rsp StateResp) (err error) {

    if rsp.Result != `200` {
        marshal, _ := json.Marshal(rsp)
        err = fmt.Errorf(`%s`, string(marshal))
    }
    return err

}

