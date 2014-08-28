package models

import (
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Mobile struct {
	Mobile   string
	Province string
	City     string
	Opera    string
	Fullname string
}

func FindLocation(mobile string) (area Mobile) {
	db, err := sql.Open("sqlite3", "data/location.db")
	if err != nil {
		log.Println(err)
	}
	rows, err := db.Query("select * from phone_location where _id = " + mobile)
	checkErr(err)

	for rows.Next() {
		var _id int
		var area string
		err = rows.Scan(&_id, &area)
		checkErr(err)
		if len(area) == 0 {
			return Mobile{Mobile:mobile}
		} else if len(area) == 20 {
			return Mobile{Mobile:mobile, Province:area[0:6], City:area[6:12], Opera:area[13:19], Fullname:area}
		}else {
			province:= area[0:9];
			if(province == "黑龙江"){
				return Mobile{Mobile:mobile, Province:area[0:9], Opera:area[10:16], Fullname:area}
			}else{
				return Mobile{Mobile:mobile, Province:area[0:6], Opera:area[7:13], Fullname:area}
			}
		}
	}
	defer db.Close()
	return Mobile{Mobile:mobile}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
