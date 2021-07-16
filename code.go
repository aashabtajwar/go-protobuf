package main

import (
	"fmt"
	"gol/model"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
    fmt.Println("now starting")
    car := &model.Car {
        Brand: "toyota",
        Model: "corolla",
        Release: 1999,
    }
    // serialize data
    serialized, err := proto.Marshal(car)
    if err != nil {
        log.Fatal("Marshalling Error: ", err)
    }
    ioutil.WriteFile("car.data", serialized, 0644)
    
    // deserialize
    car2 := &model.Car{}
    
    err = proto.Unmarshal(serialized, car2)
    if err != nil {
        log.Fatal("Unmarshal error: ", err)

    }
    fmt.Println(car2.Brand)
    fmt.Println(car2.Model)
    fmt.Println(car2.Release)
}
