/*
* @Author: Administrator
* @Date:   2017-07-29 10:53:16
* @Last Modified by:   Administrator
* @Last Modified time: 2017-07-29 12:14:48
*/

/*

CREATE TABLE `pfm`.`userinfo` ( 
`uid` INT(10)  NOT NULL AUTO_INCREMENT , 
`username` VARCHAR(64)  NULL DEFAULT NULL ,
 `departname` VARCHAR(64)  NULL DEFAULT NULL ,
  `created` DATE  DEFAULT NULL , 
  PRIMARY KEY (`uid`)
 ) ENGINE = InnoDB CHARACTER SET utf8 COLLATE utf8_general_ci;



CREATE TABLE `pfm`.`userdetail` (
    
    `uid` INT(10) NOT NULL DEFAULT '0',
    `intro` TEXT NULL,
    `profile` TEXT NULL,
    PRIMARY KEY(`uid`)
)ENGINE = InnoDB CHARACTER SET utf8 COLLATE utf8_general_ci;
 */
package main

import (
    
    "fmt"
    
    "log"
    "os"
    "strconv"
    "strings"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"


    "encoding/json"
    "net/http"
    _ "github.com/gorilla/mux"
)

const (
    dbHost         string = "127.0.0.1"
    dbPort         string = "3306"
    dbName         string = "pfm"
    dbUserName     string = "root"
    dbUserPassword string = "mysql"
    RestPort string = "6699"
)



type Person struct {
    ID        string   `json:"id,omitempty"`
    Username string   `json:"username,omitempty"`
    Departname  string   `json:"departname,omitempty"`
    Created  string   `json:"created,omitempty"`
    UserDetail   *UserDetail `json:"userdetail,omitempty"`
}
 

 type UserDetail struct {
    Intro  string `json:"intro,omitempty"`
    Profile string `json:"proflie,omitempty"`
}
 
var people []Person


func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    SetHeader(w)
    json.NewEncoder(w).Encode(&Person{})
}
 
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
    SetHeader(w)
    json.NewEncoder(w).Encode(people)
}
 
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
    
   /* params := mux.Vars(req)*/
    var person Person
    _ = json.NewDecoder(req.Body).Decode(&person)
   

    /*数据库中创建*/
    id:= Insert( person.Username, person.Departname)
    person.ID = strconv.Itoa(id)
    /*  person.ID = aaa*/
    people = append(people, person)
    SetHeader(w)
    json.NewEncoder(w).Encode(people)
}
 
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {

    params := mux.Vars(req)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            id,_ := strconv.Atoi(params["id"])
            Delete(id)
            break
        }
    }
    SetHeader(w)
    json.NewEncoder(w).Encode(people)
}

func SetHeader(w http.ResponseWriter){

     w.Header().Set("Content-Type", "application/json")
}


func getPeople(){


    db, err := ConnectDB()

    checkErr(err)

    getTaskSQL := "select * from userinfo"

    rows, err := db.Query(getTaskSQL)
    var (
        uid        int
        username   string
        departname string
        created    string
    )
    defer rows.Close()
    for rows.Next() {
        err := rows.Scan(&uid, &username, &departname, &created)
        if err != nil {
            log.Println(err)
        }

        /*获取userDetail*/
        _,intro, profile := GetUserDetailByUid( uid)

        if len(intro)>1||len(profile)>1 {
            people = append(people, Person{ID: fmt.Sprintf("%d", uid), Username: username, Departname: departname,Created:created, UserDetail: &UserDetail{Profile: profile, Intro: intro}})

        }else{
            people = append(people, Person{ID: fmt.Sprintf("%d", uid), Username: username, Departname: departname,Created:created})
        }

    }


}
 


func GetUserDetailByUid(id int)( int, string, string ) {

    db, err := ConnectDB()

    defer db.Close()

    checkErr(err)

    row := db.QueryRow("SELECT * FROM userdetail WHERE uid=?", id)

    var (
        uid        int
        intro   string
        profile string

    )
    err = row.Scan(&uid, &intro, &profile)

    if nil != err{
        return uid ,intro, profile
    }
    return  uid ,intro, profile

}


func checkErr(err error) {

    if err != nil {

        log.Println(err)
        panic(err)
    }

}

func ConnectDB() (*sql.DB, error) {

    db, err := sql.Open("mysql", dbUserName+":"+dbUserPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset-utf8&sql_mode=TRADITIONAL")
    return db, err
}

func Insert(userName string, departName string) int {

    db, err := ConnectDB()

    if err != nil {
        log.Println("插入: 打开数据库出错")
    } else {

        log.Println("插入: 已打开数据库")
    }

    defer db.Close()
    checkErr(err)

    stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")

    checkErr(err)

    res, err := stmt.Exec(userName, departName, "2017-02-05")

    checkErr(err)

    id, err := res.LastInsertId()

    checkErr(err)

    return int(id)

}

func Query(id int) (int, string, string, string) {

    db, err := ConnectDB()
    if err != nil {
        log.Println("查询: 打开数据库出错")
    } else {

        log.Println("查询:已打开数据库")
    }

    defer db.Close()

    checkErr(err)

    row := db.QueryRow("SELECT * FROM userinfo WHERE uid=?", id)

    var (
        uid        int
        username   string
        departname string
        created    string
    )
    err = row.Scan(&uid, &username, &departname, &created)
    checkErr(err)

    return uid, username, departname, created

}

func GetAll() {

    db, err := ConnectDB()

    checkErr(err)

    getTaskSQL := "select * from userinfo"

    rows, err := db.Query(getTaskSQL)
    var (
        uid        int
        username   string
        departname string
        created    string
    )
    defer rows.Close()
    fmt.Printf(strings.Repeat("*", 68) + "\n")
    fmt.Printf("%-4s\t%-16s\t%-16s\t%-20s\n", "ID", "名称", "部门", "创建时间")
    fmt.Printf(strings.Repeat("_", 68) + "\n")
    for rows.Next() {
        err := rows.Scan(&uid, &username, &departname, &created)
        if err != nil {
            log.Println(err)
        }
        fmt.Printf("%-4s\t%-16s\t%-16s\t%-20s", fmt.Sprintf("%d", uid), username, departname, created)

        fmt.Printf("\n")
    }
    fmt.Printf("\n" + strings.Repeat("*", 68))
}

func Update(id int, uName string) {

    db, err := ConnectDB()
    if err != nil {
        log.Println("更新: 打开数据库出错")
    } else {

        log.Println("更新: 已打开数据库")
    }

    defer db.Close()

    checkErr(err)

    stmt, err := db.Prepare("UPDATE userinfo SET username=? WHERE uid=?")

    checkErr(err)


    res, err := stmt.Exec(uName, id)

    checkErr(err)

    affectedRows, err := res.RowsAffected()

    checkErr(err)
    /*

        int64 转换为 string

        fmt.Sprintf("%d", affectedRows)
        strconv.FormatInt(affectedRows,10)
    */

    var affectedRowsCount string = fmt.Sprintf("%d", affectedRows)

    log.Println("更新了" + affectedRowsCount + "行数据")

}

func Delete(id int)  {

    db, err := ConnectDB()
    if err != nil {
        log.Println("删除: 打开数据库出错")
    } else {

        log.Println("删除: 已打开数据库")
    }

    defer db.Close()

    checkErr(err)

    stmt, err := db.Prepare("DELETE FROM userinfo WHERE uid=?")

    checkErr(err)

    res, err := stmt.Exec(id)

    checkErr(err)

    affect, err := res.RowsAffected()

    checkErr(err)

    log.Println(affect)

/*  return true*/

}

/*

 go build -o a.exe main.go && a

*/
func main() {





    if len(os.Args) > 2 {

        switch os.Args[1] {

        case "a":
        case "i":
            fmt.Println("开始添加")
            var id int = Insert(os.Args[2], os.Args[3])
            var Id string = strconv.Itoa(id)
            fmt.Println("已添加为 " + Id)
            GetAll()

        case "u":
            fmt.Println("开始修改")
            /*id,_ := strconv.ParseInt(os.Args[2] ,10, 32)*/
            id, _ := strconv.Atoi(os.Args[2])
            Update(id, os.Args[3])
            GetAll()

        case "d":
            fmt.Println("开始删除")
            id, _ := strconv.Atoi(os.Args[2])
            Delete(id)
            GetAll()

        case "g":
            fmt.Println("开始查询")
            fmt.Println(os.Args)
            id, _ := strconv.Atoi(os.Args[2])
            uid, username, departname, created := Query(id)

            log.Println(uid)
            log.Println(username)
            log.Println(departname)
            log.Println(created)

        }

    } else {
        GetAll()

        router := mux.NewRouter()

        /*
        运行后对数据库的修改不能及时反馈,也就是说查询是一次性的。
         */
        getPeople()


        /*
        
        [{"id":"22","username":"李四光","departname":"销售部门"},{"id":"23","username":"张三","departname":"科技"},{"id":"24","username":"李四","departname":"销售"},{"id":"25","username":"马兵地","departname":"总办"},{"id":"27","username":"吕桂花","departname":"公关部"}]


         */
        router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
        router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")

        /*创建*/
        /*
        
        {
            "Username":"王雁平",
            "Departname":"总办"
        }
         */
        router.HandleFunc("/people/", CreatePersonEndpoint).Methods("POST")
        
        /*删除*/
        router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

        fmt.Println("\n")
        fmt.Println("浏览器打开 \nhttp://localhost:"+RestPort+"/people\n")
        fmt.Println("可用的链接 \nhttp://localhost:"+RestPort+"/people/uid\n")
        log.Fatal(http.ListenAndServe(":"+RestPort, router))


    }

}
