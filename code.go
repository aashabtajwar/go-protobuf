package main

import (
	"fmt"
	"gol/model"

	"github.com/golang/protobuf/proto"
)

func main() {
    fmt.Println("now starting")
    car := &model.Car {
        Brand: "toyota",
        Model: "corolla",
        Release: 1999,
    }
    serialize, err := proto.Marshal(car)
    if err == nil {
        fmt.Println(serialize)
    } 

}
