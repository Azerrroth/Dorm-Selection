package models

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"register/pkg/setting"
)

var db *gorm.DB

var userDB, dormDB *gorm.DB

type Model struct {
	// gorm.Model
	ID        uint      `gorm:"primaryKey;PRIMARY_KEY;AUTO_INCREMENT;NOT NULL;" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

func init() {
	var (
		err                                                     error
		dbName, user, password, userHost, dormHost, tablePrefix string
		waitTime, retryTimes                                    int
		fillNums                                                = 100
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	// dbType = sec.Key("TYPE").String() // mysql
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	// host = sec.Key("HOST").String()
	userHost = sec.Key("USER_HOST").String()
	dormHost = sec.Key("DORM_HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	waitTime, _ = sec.Key("WAIT_TIME").Int()
	retryTimes, _ = sec.Key("RETRY_TIMES").Int()

	rand.Seed(time.Now().Unix())
	// db, err = ConnectDB(user, password, host, dbName, tablePrefix)
	userDB, err = ConnectDB(user, password, userHost, dbName, tablePrefix)
	dormDB, err = ConnectDB(user, password, dormHost, dbName, tablePrefix)

	if err != nil {
		fmt.Println(err)
		for i := 0; i < retryTimes; i = i + 1 {
			time.Sleep(time.Duration(waitTime) * time.Millisecond)
			userDB, err = ConnectDB(user, password, userHost, dbName, tablePrefix)
			dormDB, err = ConnectDB(user, password, dormHost, dbName, tablePrefix)
			fmt.Printf("Error: connect error, retry times: %d/%d. \n", i, retryTimes)
			if err == nil {
				break
			}
		}
	}

	userDB.AutoMigrate(&User{})
	dormDB.AutoMigrate(&Dorm{})
	FillDormIfEmpty(fillNums)
	FillUserIfEmpty(fillNums)

	if err != nil {
		log.Println(err)
	}

}

func ConnectDB(user string, password string, host string, dbName string, tablePrefix string) (db *gorm.DB, err error) {
	return gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: tablePrefix,
		},
	})
}

// func CloseDB() {
// 	defer db.Close()
// }
