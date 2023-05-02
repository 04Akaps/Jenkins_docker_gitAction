package mysql

import (
	"database/sql"
	"fmt"
	"time"

	sqlc "github.com/04Akaps/Jenkins_docker_go.git/mysql/sqlc"
	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLClient(dbName string) *sqlc.Queries {
	sourceUri := fmt.Sprintf("admin:t2yq9m!!@tcp(golang-sns-v2.cynkkrrli9o4.ap-northeast-2.rds.amazonaws.com:3306)/%s", dbName)
	// 어차피 개인용이기 떄문에 Viper같은 패키지는 사용하지 않고 FIx형태로 DB Endpoint 작업
	// 현재 DB는 sns라는 DB를 사용 중
	dbInstance, err := sql.Open("mysql", sourceUri)
	if err != nil {
		panic(err.Error())
	}

	err = dbInstance.Ping()
	if err != nil {
		panic(err.Error())
	}

	dbInstance.SetConnMaxLifetime(time.Minute * 1)
	dbInstance.SetMaxIdleConns(3)
	dbInstance.SetMaxOpenConns(6)

	fmt.Println("Connected to the database")

	return sqlc.New(dbInstance)
}
