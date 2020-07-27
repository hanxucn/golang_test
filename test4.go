package main

import (
	"database/sql"
	"fmt"
	"log"
	"encoding/json"
	_ "github.com/lib/pq"
)

var MAX_POOL_SIZE = 20
var dbPoll chan *sql.DB

const (
	host     = "localhost"
	port     = 5432
	username     = "hanxu"
	password = "123"
	dbname   = "pg_test"
)

type userInfo struct {
	userID int
	userName string
	deviceID int
}

type Response struct {
	Userid int  `json:"Userid"`
	Result bool `json:"result"`
}

type ResponseExist struct {
	Userid int	 `json:"Userid"`
	Exist  bool  `json:"Exist"`
}

func putDB(db *sql.DB) {
	if dbPoll == nil {
		dbPoll = make(chan *sql.DB, MAX_POOL_SIZE)
	}
	if len(dbPoll) >= MAX_POOL_SIZE {
		db.Close()
		return
	}
	dbPoll <- db
}

func connectDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		host, port, username, password, dbname)
	if len(dbPoll) == 0 {
		dbPoll = make(chan *sql.DB, MAX_POOL_SIZE)
		go func() {
			for i:=0; i<MAX_POOL_SIZE/2; i++ {
				db, err := sql.Open("postgres", psqlInfo)
				if err != nil {
					panic(err)
				}
				putDB(db)
			}
		}()
	}
}

func GetDB() *sql.DB {
	if dbPoll == nil || len(dbPoll) == 0 {
		connectDB()
	}
	return <- dbPoll
}

func querySingleName(username string, db *sql.DB) bool {
	var user userInfo
	row := db.QueryRow(" select userID, userName from user where userName=$1",username)
	row.Scan(&user.userID, &user.userName)
	if user == (userInfo{}) {
		return false
	} else {
		return true
	}
}

func Login(usernames map[string]string) []byte{
	var user userInfo
	db := GetDB()
	tx, err := db.Begin()
	defer tx.Commit()
	if tx == nil {
		panic(err)
	}
	var resp []Response
	for user_name, password := range usernames {
		err := tx.QueryRow(" select userID, userName from user where userName=? and password=?",
			user_name, password).Scan(&user.userID, &user.userName)
		var queryRes bool = true
		if err != nil {
			queryRes = false
			log.Println(err)
		}
		jsonType := Response{Userid: user.userID, Result: queryRes}
		resp = append(resp, jsonType)
	}

	responseJSON, _ := json.Marshal(resp)
	return responseJSON
}

func Register(userinfos map[int32]string, device_name string) []byte{
	db := GetDB()
	tx, err := db.Begin()
	defer tx.Commit()
	if tx == nil {
		panic(err)
	}
	var resp []Response
	for user_id, password := range userinfos {
		username := stableNameGenerate(user_id)
		isExist := querySingleName(username, db)
		if isExist {
			resExist := ResponseExist{Userid: int(user_id), Exist: true}
			resp, _ := json.Marshal(resExist)
			return resp
		}
		stmt, err := tx.Prepare("INSERT INTO user (userID, userName, password, deviceName) VALUES (?, ?, ?, ?)")
		if err != nil {
			log.Println(err)
		}
		res, err := stmt.Exec(user_id, username, password, device_name)
		affected, err := res.RowsAffected()
		var insertRes bool = false
		if affected > 0 {
			insertRes = true
		}
		jsonType := Response{Userid: int(user_id), Result: insertRes}
		resp = append(resp, jsonType)
	}
	responseJSON, _ := json.Marshal(resp)
	return responseJSON
}
