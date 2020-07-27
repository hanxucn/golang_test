package main

import (
	"database/sql"
	"fmt"
	"encoding/json"
	"log"
	"strings"
)


func changeDevice(userId int, verified bool, newDevice_info []string, db *sql.DB) []byte {
	/*
		newDevice_info: ["device_id", "platform"]
	 */
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}
	defer tx.Commit()
	if verified == true {
		var resp Response
		device_name := strings.Join(newDevice_info, "-")
		res, _ := tx.Exec(" UPDATE user set deviceName=? WHERE userID=?", device_name, userId)
		RowsAffected, err := res.RowsAffected()
		if err != nil {
			log.Println("UPDATE user device failed")
		}
		if RowsAffected > 0 {
			ret := addDevice(device_name, newDevice_info[1], userId, tx)
			if ret == true {
				resp = Response{Userid: int(user_id), Result: true}
				responseJSON, _ := json.Marshal(resp)
				return responseJSON
			}
		}
		resp = Response{Userid: int(user_id), Result: false}
		responseJSON, _ := json.Marshal(resp)
		return responseJSON
	}

}

func addDevice(newDevice_name string, platform string, userId int, tx *sql.Tx) bool {

	stmt, err := tx.Prepare("INSERT INTO deviceInfo (deviceName, platform, userID) VALUES (?, ?, ?)")
	if err != nil {
		log.Println(err)
	}
	res, err := stmt.Exec(newDevice_name, platform, userId)
	if err != nil {
		log.Println(err)
		return false
	}
	affected, err := res.RowsAffected()
	if affected > 0 {
		return true
	} else {
		return false
	}
}
