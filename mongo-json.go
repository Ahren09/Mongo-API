package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
    "strconv"

	"encoding/json"
)

const url  = "mongodb://0.0.0.0:27017";

type person struct {
	AGE    int  `json:"age"`
	NAME   string `json:"name"`
	HEIGHT int    `json:"height"`
}

type User struct {
    Id        json.ObjectId `json:"_id"`
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

func createUser(name string, owner string, createdby string) (*User) {
    newUser := new(User)
    newUser.Id = json.NewObjectId()
    newUser.Name = name
    newUser.Owner = owner
    newUser.Creator = createdby

    return newUser
}

func Insert(mongoUrl string, dbName string, collectionName string, user *User) {
    _, collection := connect(mongoUrl, dbName, collectionName)

    err := collection.Insert(user)
    if err == nil {
        fmt.Println("Successful Insertion")
    } else {
        fmt.Println(err.Error())
    }
}


func Find(mongoUrl string, dbName string, collectionName string, name string) []User{
    _, collection := connect(mongoUrl, dbName, collectionName)
    var users []User
    if name != ""{
        collection.Find(json.M{"name": name}).All(&users)
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

func main() {
	ListCollection(mongoUrl, "newBlog", "newMgotest")
	user1 := createUser("Cassie", "123@yahoo.com", "N/A")
	Insert(mongoUrl, "newBlog", "newMgotest", user1)
	ListCollection(mongoUrl, "newBlog", "newMgotest")


	// p:=person{
	// 	33,
	// 	"周杰伦",
	// 	175,
	// }
	// err=op.insert(p)
	// if err != nil {
	// 	fmt.Println("插入出错",err)
	// }
	// op.update()
	// op.query();

	// count,err:=op.count()
	// if err !=nil {
	// 	fmt.Println("统计出错",err)
	// 	return
	// }

	// err=op.delete(&json.M{"height": 0})
	// if err!=nil {
	// 	fmt.Println("删除错误",err)
	// }else{
	// 	fmt.Println("删除成功")
	// }

	// fmt.Println("共有数据",count)
	// op.mogSession.Close()
}
//连接数据库
// func (operater *Operater) connect() error {
// 	mogsession, err := mgo.Dial(url)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	operater.mogSession=mogsession
// 	return nil
// }
// //插入
// func (operater *Operater) insert( p person) error {
// 	collcetion:=operater.mogSession.DB(operater.dbname).C(operater.document)
// 	err:=collcetion.Insert(p)
// 	return err
// }
// //查询所有
// func (operater *Operater) queryAll() ([]person,error) {
// 	collcetion:=operater.mogSession.DB(operater.dbname).C(operater.document)
// 	p:=new(person)
// 	p.AGE=33
// 	query:=collcetion.Find(nil)
// 	ps:=[]person{}
// 	query.All(&ps)
// 	iter:=collcetion.Find(nil).Iter()
// 	//
// 	result:=new(person)
// 	for iter.Next(&result) {
// 		fmt.Println("一个一个输出：", result)
// 	}
// 	return ps,nil
// }
// //条件查询
// func (operater *Operater) query() ([]person,error) {
// 	collcetion:=operater.mogSession.DB(operater.dbname).C(operater.document)
// 	p:=new(person)
// 	p.AGE=33
// 	query:=collcetion.Find(json.M{"age":json.M{"$eq":21}})
// 	ps:=[]person{}
// 	query.All(&ps)
// 	fmt.Println(ps)
// 	return ps,nil
// }

// //更新一行
// func (operater *Operater) update() (error) {
// 	collcetion:=operater.mogSession.DB(operater.dbname).C(operater.document)
// 	update:=person{
// 		33,
// 		"詹姆斯",
// 		201,
// 	}
// 	err:=collcetion.Update(json.M{"name": "周杰伦"},update)
// 	if err !=nil {
// 	    fmt.Println(err)
// 	}
// 	return err
// }
// //更新所有数据
// func (operater *Operater) updateAll() (error) {
// 	collcetion:=operater.mogSession.DB(operater.dbname).C(operater.document)
// 	update:=person{
// 		33,
// 		"詹姆斯",
// 		201,
// 	}
// 	changeinfo,err:=collcetion.UpdateAll(json.M{"name": "周杰伦"},update)
// 	if err !=nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("共有多少行",changeinfo.Matched,"影响")
// 	return nil
// }


// //单行删除
// func (operater *Operater) delete(seletor interface{}) (error) {
// 	collcetion:=operater.mogSession.DB(operater.dbname).C(operater.document)
// 	return collcetion.Remove(seletor)
// }

// //统计文档中数据的个数
// func (operater *Operater) count() (int,error) {
// 	collcetion:=operater.mogSession.DB(operater.dbname).C(operater.document)
// 	i,err:=collcetion.Count()
// 	return i,err
// }








