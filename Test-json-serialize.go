package main

import (
	"fmt"
	// "net/url"
	// "strings"
	// "gopkg.in/mgo.v2"
	"encoding/json"
    //"strconv"
    //"gopkg.in/mgo.v2/bson"
)

type User struct {
    Id        byte `json:"id"`
    Name      string        `json:"name"`
    Owner     string        `json:"owner"`
    Creator   string         `json:"createdby"`
}

func main(){
   
    u1:=`{
     "db": "DB",
     "collection": "Col",
     "documents": [
         {
         "id": "0",
         "name": "myName",
         "owner": "123@yahoo.com",
         "createdby": "123@yahoo.com"
         }
        ]
     }`
    fmt.Println("--------\nType of u1:")
    fmt.Printf("%T\n", u1)

    var obj interface{}
    err := json.Unmarshal([]byte(u1), &obj)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("--------\n", obj)
    }
    data, err := json.Marshal(obj)
    str:=string(data)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("--------\n", str)
    }

    // Test map
    mp:=map[string]string{}
    err = json.Unmarshal([]byte(u1), &mp)
    fmt.Println("--------\n", mp["db"])

    // Print type of mp["document"]
    fmt.Println("--------\nType of mp[\"document\"]:")
    fmt.Printf("%T\n", mp["document"])
    fmt.Println(mp["document"])

    // mp2:=map[string]string{}
    // err = json.Unmarshal([]byte(mp), &mp2)
    // fmt.Println("--------\n", mp2["db"])

}
