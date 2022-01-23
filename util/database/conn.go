package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	// Driver database
	Driver string
	// DataSource connection script
	DataSource string
)

// SQLi interface for wrapping sqlx.DB and sqlx.Tx
// type SQLi interface {
// 	PrepareNamedContext(context.Context, string) (*sqlx.NamedStmt, error)
// 	MustExecContext(context.Context, string, ...interface{}) sql.Result
// 	SelectContext(context.Context, interface{}, string, ...interface{}) error
// 	GetContext(context.Context, interface{}, string, ...interface{}) error
// 	NamedExecContext(context.Context, string, interface{}) (sql.Result, error)
// 	BeginTxx(context.Context, *sql.TxOptions) (*sqlx.Tx, error)
// 	sqlx.ExtContext
// 	sqlx.PreparerContext
// }

// SqlxDBMysql is struct for Sqlx Connection
type SqlxDBMysql struct{}

func open() (*sql.DB, error) {
	return sql.Open(Driver, DataSource)
}

func (d *SqlxDBMysql) buildConnection() (*sqlx.DB, error) {
	db, err := open()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("DB Server is connected!")

	return sqlx.NewDb(db, Driver), nil
}

// GetSqlxConnection create connection for sqlx
func GetSqlxConnection() (*sqlx.DB, error) {
	urlConnection := fmt.Sprint(os.Getenv("DB_USER")) + ":" + fmt.Sprint(os.Getenv("DB_PASSWORD")) + "@/" + fmt.Sprint(os.Getenv("DB_NAME"))

	log.Println("Connecting to DB Server " + fmt.Sprint(os.Getenv("DB_HOST")) + ":" + fmt.Sprint(os.Getenv("DB_PORT")) + "...")

	Driver = os.Getenv("DB_DRIVER")
	DataSource = urlConnection

	db := SqlxDBMysql{}
	return db.buildConnection()
}
