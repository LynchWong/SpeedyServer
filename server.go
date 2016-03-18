package main

import (
	"fmt"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"SpeedyServer/database"
	"SpeedyServer/types"
	"html/template"
	"os"
	"io"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()//解析url传递的参数，对于POST则解析响应包的主体（request body）
	name := r.Form["name"][0]
	age, _ := strconv.Atoi(r.Form["age"][0])
	var httpResult  types.HttpResult
	if err := mysql.Create(name, age); err == nil {
		httpResult = types.HttpResult{Code:200, Message:"添加成功!", Data:nil}
	} else {
		httpResult = types.HttpResult{Code:1001, Message:err.Error(), Data:nil}
	}
	fmt.Fprint(w, httpResult.JsonString())
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var httpResult types.HttpResult
	if userList, err := mysql.Read(); err == nil {
		httpResult = types.HttpResult{Code:200, Message:"获取成功!", Data:userList}
	} else {
		httpResult = types.HttpResult{Code:1001, Message:err.Error(), Data:nil}
	}
	fmt.Fprint(w, httpResult.JsonString())
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var httpResult types.HttpResult
	if err := mysql.Update(); err == nil {
		httpResult = types.HttpResult{Code:200, Message:"更新成功!", Data:nil}
	} else {
		httpResult = types.HttpResult{Code:1001, Message:err.Error(), Data:nil}
	}
	fmt.Fprint(w, httpResult.JsonString())
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()//解析url传递的参数，对于POST则解析响应包的主体（request body）
	id, _ := strconv.Atoi(r.Form["id"][0])
	var httpResult types.HttpResult
	if err := mysql.Delete(id); err == nil {
		httpResult = types.HttpResult{Code:200, Message:"删除成功!", Data:nil}
	} else {
		httpResult = types.HttpResult{Code:1001, Message:err.Error(), Data:nil}
	}
	fmt.Fprint(w, httpResult.JsonString())
}

func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, err := template.ParseFiles("./views/upload.gptl")
		if err != nil {
			err.Error()
		}
		t.Execute(w, nil)
	} else {
		file, handle, err := r.FormFile("file")
		if err != nil {
			err.Error()
		}
		//如果当前目录下没有upload文件夹，不会自动创建。需要手动创建
		f, err := os.OpenFile("./upload/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		io.Copy(f, file)
		if err != nil {
			err.Error()
		}
		defer f.Close()
		defer file.Close()

		file1, handle1, err1 := r.FormFile("file1")
		if err1 != nil {
			err1.Error()
		}
		//如果当前目录下没有upload文件夹，不会自动创建。需要手动创建
		f1, err1 := os.OpenFile("./upload/"+handle1.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		io.Copy(f1, file1)
		if err1 != nil {
			err1.Error()
		}
		defer f1.Close()
		defer file1.Close()

		fmt.Fprint(w, "upload success")
	}
}

var mysql *database.Mysql

func main() {

	var err error
	if mysql, err = database.InitMysql(); err == nil {
		fmt.Println("数据库初始化成功!")
	}

	http.Handle("/upload/", http.StripPrefix("/upload/", http.FileServer(http.Dir("./upload"))))

	http.HandleFunc("/add", AddUser)
	http.HandleFunc("/get", GetUser)
	http.HandleFunc("/update", UpdateUser)
	http.HandleFunc("/delete", DeleteUser)
	http.HandleFunc("/upload", Upload)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe:8080 :", err)
	}
}

