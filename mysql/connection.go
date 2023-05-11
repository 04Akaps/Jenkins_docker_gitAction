package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	sqlc "github.com/04Akaps/Jenkins_docker_go.git/mysql/sqlc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattes/migrate/source/file"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLClient(dbName string) *sqlc.Queries {
	sourceUri := fmt.Sprintf("admin:t2yq9m!!@tcp(golang-sns-v2.cynkkrrli9o4.ap-northeast-2.rds.amazonaws.com:3306)/%s", dbName+"?parseTime=true")
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

	log.Printf("Connected to the database : %s", dbName)

	return sqlc.New(dbInstance)
}

func MigratMysql(db *sql.DB) {
	// 초기에만 사용할 Migrate
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal("launchpad instance Error : ", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://mysql/migrate",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal("db instance not found : ", err)
	}

	version, dirty, _ := m.Version()

	if version == 0 {
		version += 1
	}

	if dirty {
		_ = m.Drop()
	}

	if err := m.Migrate(version); err != nil {
		if err != migrate.ErrNoChange {
			log.Fatal("-----: ", err)
		}
	}

	_ = m.Steps(2)
}
