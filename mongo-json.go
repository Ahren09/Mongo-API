package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
    //"strconv"

	"encoding/json"
)

const url  = "mongodb://0.0.0.0:27017";

type User struct {
    Id        string `json:"id"`
    Name      string        `json:"name"`
    Owner     string        `json:"owner"`
    Creator   string         `json:"createdby"`
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
    var userMap map[string] interface{}
    _ =json.Unmarshal(user, &userMap)
    err := collection.Insert(userMap)
    if err == nil {
        fmt.Println("Successful Insertion")
    } else {
        fmt.Println(err.Error())
    }
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

func ListCollection(mongoUrl string, dbName string, collectionName string) {
    fmt.Println("++++++++++ List Collection ++++++++++++")
    users := Find(mongoUrl, dbName, collectionName, "")
    //for _, u := range users{
        fmt.Println("[User]")
        fmt.Println("Id: ", users["id"])
        fmt.Println("Name: ", users["name"])
        fmt.Println("Owner: ", users["owner"])
        fmt.Println("Creator: ", users["createdby"])
    //}
    fmt.Println("++++++++++ End ++++++++++++")
}

func main() {
    ListCollection(mongoUrl, "newBlog", "newMgotest")
    user := "{\"id\": \"0\",\"name\": \"mycluster2\",\"owner\": \"ykunyk@cn.ibm.com\",\"createdby\": \"ykunyk@cn.ibm.com\"}"
    userByte := []byte(user)
    Insert(mongoUrl, "newBlog", "newMgotest", userByte)
    ListCollection(mongoUrl, "newBlog", "newMgotest")
}