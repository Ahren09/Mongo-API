package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
    //"strconv"
    // "encoding/gob"
    // "bytes"

	"encoding/json"
)

const url  = "mongodb://0.0.0.0:27017";

type Body struct {
    DB         string     `json:"db"`
    Collection string     `json:"collection"`
    Documents []Documents `json:"documents"`
}

type Documents struct {
    Id        string        `json:"id"`
    Name      string        `json:"name"`
    Owner     string        `json:"owner"`
    Creator   string        `json:"createdby"`
}

const mongoUrl string = "127.0.0.1:27017"

var collection *mgo.Collection
var session *mgo.Session

func connect(mongoUrl, dbName string, collectionName string) (*mgo.Database, *mgo.Collection) {
    session, _ = mgo.Dial(mongoUrl)
    db := session.DB(dbName)
    collection = db.C(collectionName)
    return db, collection
}

func Insert(mongoUrl string, dbName string, collectionName string, user []byte) {
    _, collection := connect(mongoUrl, dbName, collectionName)

    fmt.Println("=========== Insert ============")
    // var userMap map[string] interface{}
    // _ =json.Unmarshal(user, &userMap)
    // err := collection.Insert(userMap)
    // fmt.Println("[User]")
    // fmt.Println("db: ", userMap["db"])
    // fmt.Println("collection: ", userMap["collection"])
    // //fmt.Println("Id: ", userMap["id"]) //Output <nil>
    // documents := userMap["documents"]
    // fmt.Println("Document:", documents)
    // fmt.Printf("Document Type:%T\n", documents)
    
    var bodyMap Body
    err := json.Unmarshal(user, &bodyMap)
    fmt.Println("db: ", bodyMap.DB)
    fmt.Println("collection: ", bodyMap.Collection)
    documents := bodyMap.Documents[0]
    fmt.Println("Documents:", documents)
    fmt.Println("Id:", bodyMap.Documents[0].Id)
    fmt.Println("Name:", bodyMap.Documents[0].Name)

    err = collection.Insert(documents)
    


    if err == nil {
        fmt.Println("Successful Insertion")
    } else {
        fmt.Println(err.Error())
    }
    fmt.Println("=========== Insert End ============")
}


func Find(mongoUrl string, dbName string, collectionName string, name string) map[string]interface{} {
    _, collection := connect(mongoUrl, dbName, collectionName)
    var users map[string]interface{}
    if false {
        //collection.Find(json.M{"name": name}).All(&users)
    } else {
        collection.Find(nil).One(&users)
    }
    // fmt.Println("Number of users:", len(users))
    // // Print Id's of all returned users
    // for _,value := range users {
    //     fmt.Println(value.Id)
    // }
    return users
}

func FindByStruct(mongoUrl string, dbName string, collectionName string) []Documents {
    _, collection := connect(mongoUrl, dbName, collectionName)
    var users []Documents
    collection.Find(nil).All(&users)
    
    fmt.Println("Number of users:", len(users))
    
    for _,u := range users {
        fmt.Println(u.Id)
        fmt.Println(u.Owner)
    }
    return users
}

func ListCollection(mongoUrl string, dbName string, collectionName string) {
    fmt.Println("++++++++++ List Collection ++++++++++++")
    user := Find(mongoUrl, dbName, collectionName, "")
    //for _, user := range users{
        fmt.Println("[User]")
        fmt.Println("Id: ", user["id"])
        fmt.Println("Creator: ", user["createdby"])
    //}
    fmt.Println("++++++++++ End ++++++++++++")
}

func main() {
    // ListCollection(mongoUrl, "newBlog", "newMgotest")
    _ = FindByStruct(mongoUrl, "newBlog", "newMgotest")
    user := "{\"db\": \"k8srepo\",\"collection\": \"clusters\",\"documents\": [{\"id\": \"0\",\"name\": \"mycluster2\",\"owner\": \"ykunyk@cn.ibm.com\",\"createdby\": \"ykunyk@cn.ibm.com\"}]}"
    userByte := []byte(user)
    Insert(mongoUrl, "newBlog", "newMgotest", userByte)
    _ = FindByStruct(mongoUrl, "newBlog", "newMgotest")
    //ListCollection(mongoUrl, "newBlog", "newMgotest")
}