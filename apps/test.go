package main

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws/awserr"
)

func main()  {

    e := awserr.New(`code`, `msg`, fmt.Errorf(`error`))

    fmt.Println(e)
}
