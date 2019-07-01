// mgotest project main.go
package main

import (
    "fmt"
    "time"

    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type User struct {
    Id        bson.ObjectId `bson:"_id"`
    Username  string        `bson:"name"`
    Pass      string        `bson:"pass"`
    Regtime   int64         `bson:"regtime"`
    Interests []string      `bson:"interests"`
}

const URL string = "127.0.0.1:27017"

var c *mgo.Collection
var session *mgo.Session

func (user User) ToString() string {
    return fmt.Sprintf("%#v", user)
}

func init() {
    session, _ = mgo.Dial(URL)
    //切换到数据库
    db := session.DB("blog")
    //切换到collection
    c = db.C("mgotest")
}

func add() {
    //    defer session.Close()
    stu1 := new(User)
    stu1.Id = bson.NewObjectId()
    stu1.Username = "Ahren"
    stu1.Pass = "pass"
    stu1.Regtime = time.Now().Unix()
    stu1.Interests = []string{"C++", "Python", "ML"}
    err := c.Insert(stu1)
    if err == nil {
        fmt.Println("Inserted!")
    } else {
        fmt.Println(err.Error())
        defer panic(err)
    }
}

func update(){
    interests := "Quick pass"
    err := c.Update(bson.M{"name":"Ahren"}, bson.M{"$set":bson.M{
        "name": "Jin",
        "pass": interests,
        "regtime": time.Now().Unix(),
        "interests": interests,
    }})
    if err != nil{
        fmt.Println("Fail")
        fmt.Println(err.Error())
    } else {
        fmt.Println("Success!")
    }
}

func main() {
    update()
}