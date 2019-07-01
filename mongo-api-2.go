// mgotest project main.go
package main

import (
    "fmt"
    //"time"

    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type User struct {
    Id        bson.ObjectId `bson:"_id"`
    Name      string        `bson:"name"`
    Owner     string        `bson:"owner"`
    Creator   string         `bson:"createdby"`
}

const mongoUrl string = "127.0.0.1:27017"

var collection *mgo.Collection
var session *mgo.Session

func (user User) ToString() string {
    return fmt.Sprintf("%#v", user)
}

func connect(mongoUrl, dbName string, collectionName string) (*mgo.Database, *mgo.Collection) {
    session, _ = mgo.Dial(mongoUrl)
    db := session.DB(dbName)
    collection = db.C(collectionName)
    return db, collection
}

func createUser(name string, owner string, createdby string) (*User) {
    newUser := new(User)
    newUser.Id = bson.NewObjectId()
    newUser.Name = name
    newUser.Owner = owner
    newUser.Creator = createdby

    return newUser
}

func createInterest(name string, owner string, createdby string) (string) {
    if(name == ""){
        fmt.Println("Empty")
    }
    return ""

}


// func init() {
//     session, _ = mgo.Dial(mongoUrl)
//     db := session.DB("newBlog")
//     collection = db.C("newMgotest")
// }

func Add(mongoUrl string, dbName string, collectionName string, user *User) {
    _, collection := connect(mongoUrl, dbName, collectionName)

    err := collection.Insert(user)
    if err == nil {
        fmt.Println("Inserted!")
    } else {
        fmt.Println(err.Error())
    }
}

// func Delete(mongoUrl string, dbName string, collectionName string, interest string) {
//     err := collection.Remove(bson.M{})
// }

func Find(mongoUrl string, dbName string, collectionName string, name string) []User{
    _, collection := connect(mongoUrl, dbName, collectionName)
    var users []User
    collection.Find(bson.M{"name": name}).All(&users)
    // Showing Id's of all returned users
    for _,value := range users {
        fmt.Println(value.Id)
    }
    return users
}

// func Update {

// }




func main() {
    createInterest("Here!", "", "")
    user := createUser("Ahren", "newBlog", "newMgotest")
    Add(mongoUrl, "newBlog", "newMgotest", user)
    Find(mongoUrl, "newBlog", "newMgotest", "Ahren")
}