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
    session, _ := mgo.Dial(mongoUrl)
    db := session.DB(dbName)
    collection := db.C(collectionName)
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
    interest := ""
    if(name != ""){
        interest += "\"name\": \"" + name + "\", "
    }
    if(owner != ""){
        interest += "\"owner\": \"" + owner + "\", "
    }
    if(createdby != ""){
        interest += "\"createdby\": \"" + createdby + "\", "
    }
    fmt.Println(interest)
    return interest
}



// func init() {
//     session, _ = mgo.Dial(mongoUrl)
//     db := session.DB("newBlog")
//     collection = db.C("newMgotest")
// }

func Insert(mongoUrl string, dbName string, collectionName string, user *User) {
    _, collection := connect(mongoUrl, dbName, collectionName)

    err := collection.Insert(user)
    if err == nil {
        fmt.Println("Successful Insertion")
    } else {
        fmt.Println(err.Error())
    }
}

func Update (mongoUrl string, dbName string, collectionName string, newUser *User) {
    _, collection := connect(mongoUrl, dbName, collectionName)
    // Create object
    err := collection.Update(bson.M{"name":newUser.Id}, bson.M{"$set": *newUser})
    if err == nil {
        fmt.Println("Successful Update!")
    } else {
        fmt.Println(err.Error())
    }
}

func Upsert (mongoUrl string, dbName string, collectionName string, newUser *User) {
    _, collection := connect(mongoUrl, dbName, collectionName)
    // Create object
    fmt.Println("New User Id: ", newUser.Id)
    _, err := collection.UpsertId(newUser.Id, bson.M{"$set": *newUser})
    if err == nil {
        fmt.Println("Successful Upsert!")
    } else {
        fmt.Println(err.Error())
    }
}

func Delete(mongoUrl string, dbName string, collectionName string, id bson.ObjectId) {
    _, collection := connect(mongoUrl, dbName, collectionName)
    //fmt.Println(interest)
    _, err := collection.RemoveAll(bson.M{
        "_id" : id,
    })
    if err == nil {
        fmt.Println("Successful Deletion!")
        var users []User
        var name string
        fmt.Println("\nList current state of database:")
        collection.Find(bson.M{"name": name}).All(&users)
        for _,value := range users {
            fmt.Println(value.Id)
        }
    } else {
        fmt.Println(err.Error())
    }

}

func Count(mongoUrl string, dbName string, collectionName string, name string) {
    _, collection := connect(mongoUrl, dbName, collectionName)
    num, err := collection.Count() //bson.M{"name": "Ahren"})
    if err == nil {
        fmt.Println("Number of records:")
        fmt.Println(num)
    } else {
        fmt.Println(err.Error())
    }
}

func Find(mongoUrl string, dbName string, collectionName string, name string) []User{
    _, collection := connect(mongoUrl, dbName, collectionName)
    var users []User
    if name != ""{
        collection.Find(bson.M{"name": name}).All(&users)
    } else {
        collection.Find(nil).All(&users)
    }
    fmt.Println("Number of users:", len(users))
    // Print Id's of all returned users
    for _,value := range users {
        fmt.Println(value.Id)
    }
    return users
}

func ListCollection(mongoUrl string, dbName string, collectionName string) {
    fmt.Println("++++++++++ List Collection ++++++++++++")
    users := Find(mongoUrl, dbName, collectionName, "")
    for _, u := range users{
        fmt.Println("[User]")
        fmt.Println("Id: ", u.Id)
        fmt.Println("Name: ", u.Name)
        fmt.Println("Owner: ", u.Owner)
        fmt.Println("Creator: ", u.Creator)
    }
    fmt.Println("++++++++++ End ++++++++++++")
}


func TestInsert() {
    fmt.Println("======= Test Insert ======")
    user1 := createUser("Cassie", "123@yahoo.com", "N/A")
    Insert(mongoUrl, "newBlog", "newMgotest", user1)
    Find(mongoUrl, "newBlog", "newMgotest", "Cassie")
    fmt.Println("========= END =========")
}

func TestDelete() {
    fmt.Println("======= Test Delete ======")
    users := Find(mongoUrl, "newBlog", "newMgotest", "Cassie")
    for _, value := range users{
        Delete(mongoUrl, "newBlog", "newMgotest", value.Id)
    }
    fmt.Println("========= END =========")


}

func TestUpsert() {
    fmt.Println("======= Test Upsert ======")
    users := Find(mongoUrl, "newBlog", "newMgotest", "Ahren")
    for _,value := range users {
        value.Creator = "New Creator"
        Upsert(mongoUrl, "newBlog", "newMgotest", &value)
    }

    fmt.Println("========= END =========")
}



func main() {
    TestInsert()
    ListCollection(mongoUrl, "newBlog", "newMgotest")
    TestDelete()
    ListCollection(mongoUrl, "newBlog", "newMgotest")

    TestUpsert()
    ListCollection(mongoUrl, "newBlog", "newMgotest")
    //Find(mongoUrl, "newBlog", "newMgotest", "Ahren")
    // Count(mongoUrl, "newBlog", "newMgotest", "Ahren")
    // user1 := createUser("Bob", "newBlog", "newMgotest")
    // Update(mongoUrl, "newBlog", "newMgotest", "Ahren", *user)

    // // Insert(mongoUrl, "newBlog", "newMgotest", user)
    // Find(mongoUrl, "newBlog", "newMgotest", "Bob")
    // createInterest("Here!", "", "")
   

    // Find(mongoUrl, "newBlog", "newMgotest", "Ahren")

    // 
    // user1 := createUser("Cassie", "123@yahoo.com", "N/A")
    // interest_1 := createInterest("Cassie", "123@yahoo.com", "N/A")
    // // fmt.Println(interest_1)
    // Delete(mongoUrl, "newBlog", "newMgotest", interest_1)
    // Find(mongoUrl, "newBlog", "newMgotest", "Cassie")



}