package repository

import (
	"log"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DBManager struct {
	DB *sqlx.DB
}

var (
	db   *sqlx.DB
	once sync.Once
)

func GetDB() *sqlx.DB {
	once.Do(func() {
		var err error
		db, err = dbInit()
		if err != nil {
			panic(err)
		}

		if err = db.Ping(); err != nil {
			log.Printf("repository.GetDB: %v", err)
			panic(err)
		}
	})

	err := db.Ping()
	if err != nil {
		// 接続がタイムアウトしていたら、もう一度接続を行う
		db, err = dbInit()
		if err != nil {
			panic(err)
		}
	}

	return db
}

func dbInit() (*sqlx.DB, error) {
	// SQLiteファイルのパスを環境変数から取得
	dbPath := os.Getenv("DatabasePath")
	if dbPath == "" {
		dbPath = "./db.sqlite3" // デフォルトパス
	}

	// SQLiteに接続
	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		panic("Error connecting to SQLite database")
	}

	// 外部キー制約を有効にする
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		db.Close()
		return nil, err
	}

	// 接続のタイムアウト時間を設定（SQLiteではあまり意味がないが一応設定）
	db.SetConnMaxLifetime(0) // 無制限

	// 接続が成功したかどうかを確認
	err = db.Ping()
	if err != nil {
		panic("Error pinging SQLite database")
	}
	return db, err
}
