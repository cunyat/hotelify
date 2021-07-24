package service

import (
	"errors"
	"os"
	"time"

	"github.com/cunyat/hotelify/internal/rooms/adapters/storage"
	"github.com/cunyat/hotelify/internal/rooms/domain/room"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewRoomRepository() room.Repository {
	// config, err := loadConfig()
	// if err != nil {
	//	panic(err.Error())
	// }

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/rooms", config.user, config.password, config.host, config.port)
	// db, err := sql.Open("mysql", dsn)
	db := sqlx.MustOpen("sqlite3", "database.sqlite")
	_, err := db.Exec(
		`create table if not exists rooms (
    uuid varchar(38) not null primary key,
    num varchar(61) not null,
    floor integer not null,
    services varchar(255) not null

) CHARACTER SET utf8mb4
  COLLATE utf8mb4_bin;
		
		create table if not exists beds (
			id integer not null primary key autoincrement,
			room_uuid varchar(38) not null,
			bed_type varchar(32) not null,
			count int not null
		) CHARACTER SET utf8mb4
			COLLATE utf8mb4_bin;
`)
	if err == nil {
		panic(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return storage.NewMysqlRoomRepository(db)
}

func loadConfig() (repoConfig, error) {
	user, ok := os.LookupEnv("MYSQL_USER")
	if !ok {
		return repoConfig{}, errors.New("MYSQL_USER env var not foud")
	}

	passord, ok := os.LookupEnv("MYSQL_PASSWORD")
	if !ok {
		return repoConfig{}, errors.New("MYSQL_PASSWORD env var not foud")
	}

	host, ok := os.LookupEnv("MYSQL_HOST")
	if !ok {
		return repoConfig{}, errors.New("MYSQL_HOST env var not foud")
	}
	port, ok := os.LookupEnv("MYSQL_PORT")
	if !ok {
		return repoConfig{}, errors.New("MYSQL_PORT env var not foud")
	}
	return repoConfig{
		user:     user,
		password: passord,
		host:     host,
		port:     port,
	}, nil
}

type repoConfig struct {
	user     string
	password string
	host     string
	port     string
}
