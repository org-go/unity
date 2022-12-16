package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "unity.service/apps/v1/dto/common"
)

//
//  Response
//  @Description:
//
type Response struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

type pageResponse struct {
    List interface{}
    common.PaginationResult
}

/**
 * $(success)
 * @Description:
 * @param c
 * @param data
 */
func success(c *gin.Context, data interface{}, page ...interface{}) {

	response := replyWithPage(http.StatusOK, ``, data, page)

	c.JSON(http.StatusOK, response)
}


/**
 * $(failure)
 * @Description:
 * @param c
 * @param code
 * @param message
 */
func failure(c *gin.Context, code int, message string) {

	response := reply(code, message, nil)
	c.JSON(http.StatusOK, response)
}
/**
 * $(logError)
 * @Description:
 * @param c
 */
func logError(c *gin.Context) {

    if err := recover(); err != nil {
        c.JSON(http.StatusOK, err)
        return
    }
}

func reply(code int, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func replyWithPage(code int, message string, data interface{}, page interface{}) Response {

    if page != nil {
        data = pageResponse{
            List:             data,
            PaginationResult: page.(common.PaginationResult),
        }
    }
    return reply(code, message, data)

}
