package coreSystem

import (
	"context"
	"fmt"
	"unity.service/pkg/go/clientx"
)


/**
 * $(Payment)
 * @Description:
 * @receiver this
 * @param ctx
 * @param opt
 * @return pay
 * @return err
 */
func (this core) Payment(ctx context.Context) (pay string, err error) {

	option := this.option().Payment
	var tp string

	fmt.Println(tp,`iss.core.dial:`, option.Uri, string(`msr`))
	header := make(map[string]string)
	header[`Authorization-Key`] = `3PLATZ-ID`
	data, err := clientx.RequestWithRetry(option.Uri, option.Request, string(`msr`), header, 1)

	if err != nil {
		return pay, err
	}
	return data, nil

}



