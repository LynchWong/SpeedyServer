package database

import (
	"database/sql"
	"fmt"
	"SpeedyServer/types"
	"SpeedyServer/models"
)

type Mysql struct  {
	db *sql.DB
}

func InitMysql() (*Mysql, error) {
	mysql := new(Mysql)
	db, err := sql.Open("mysql", "root:lynch@/database")
	if err != nil {
		fmt.Println("database initialize error : ", err.Error())
		return nil, err
	}
	mysql.db = db
	return mysql, nil
}
func (mysql *Mysql)Create(name string, age int) error {
	if mysql.db == nil {
		return types.CRUDError{ErrorMessage:"数据库不存在!"}
	}
	stmt, err := mysql.db.Prepare("INSERT INTO User(name, age) VALUES( ?, ? )")
	if err != nil {
		fmt.Println(err.Error())
		return types.CRUDError{ErrorMessage:err.Error()}
	}
	defer stmt.Close()

	if result, err := stmt.Exec(name, age); err == nil {
		if id, err := result.LastInsertId(); err == nil {
			fmt.Println("Insert Id: ", id)
		}
	}
	return nil
}

func (mysql *Mysql)Read() ([]models.User, error) {
	if mysql.db == nil {
		return nil, types.CRUDError{ErrorMessage:"数据库不存在!"}
	}
	//stmt, err := mysql.db.Prepare("SELECT id, name, age FROM User limit 0,5")
	rows, err := mysql.db.Query("SELECT * FROM User limit 0,5")
	if err != nil {
		fmt.Println(err.Error())
		return nil, types.CRUDError{ErrorMessage:err.Error()}
	}
	defer rows.Close()

	cols, _ := rows.Columns()
	for i := range cols {
		fmt.Println(cols[i])
	}

	var users []models.User

	var (
		id int
		name string
		age int
	)
	for rows.Next() {
		if err := rows.Scan(&id, &name, &age); err == nil {
			fmt.Printf("id: %d, name: %s, age: %d", id, name, age)
			//user := new(User)
			//user.Id = id
			//user.Name = name
			//user.Age = age
			//users = append(users, user)
			users = append(users, models.User{Id:id, Name:name, Age:age})

		}
	}
	return users, nil
}

func (mysql *Mysql)Update() error {
	if mysql.db == nil {
		return types.CRUDError{ErrorMessage:"数据库不存在!"}
	}
	stmt, err := mysql.db.Prepare("UPDATE User SET name = ?, age = ? WHERE id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return types.CRUDError{ErrorMessage:err.Error()}
	}
	defer stmt.Close()

	if result, err := stmt.Exec("Lynch Wong", "27", 19); err == nil {
		if c, err := result.RowsAffected(); err == nil {
			fmt.Println("Update Count : ", c)
		}
	}
	return nil
}

func (mysql *Mysql)Delete(id int) error {
	if mysql.db == nil {
		return types.CRUDError{ErrorMessage:"数据库不存在!"}
	}
	stmt, err := mysql.db.Prepare("DELETE FROM User WHERE id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return types.CRUDError{ErrorMessage:err.Error()}
	}
	defer stmt.Close()

	if result, err := stmt.Exec(id); err == nil {
		if c, err := result.RowsAffected(); err == nil {
			fmt.Println("Remove Count :", c)
		}
	}
	return nil
}