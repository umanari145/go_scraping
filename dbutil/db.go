package dbutil

import (
	"fmt"
	"os"

	"github.com/go-easylog/el"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

//Connect DBへの接続
func Connect() (*gorm.DB, error) {
	DBMS := os.Getenv("DB_TYPE")
	USER := os.Getenv("POSTGRES_DBUSER")
	PASS := os.Getenv("POSTGRES_PASSWORD")
	DBNAME := os.Getenv("POSTGRES_DBNAME")

	CONNECT := "host=" + os.Getenv("POSTGRES_DBHOST") +
		" port=5432" +
		" user=" + USER +
		" dbname=" + DBNAME +
		" password=" + PASS +
		" sslmode=disable"

	fmt.Printf("DB_CONNECT:%s", CONNECT)

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		el.Error("DB cannnot connect")
		el.Errorf(" error message %s", err.Error())
		return nil, err
	}
	db.LogMode(true)
	el.Info("DB connect")

	return db, nil
}
