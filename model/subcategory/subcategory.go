package subcategory

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var (
	table = "subcategories"
)

type Item struct {
	ID              uint32         `db:"id"`
	NameSubCategory string         `db:"name_category"`
	CreatedAt       mysql.NullTime `db:"created_at"`
	UpdatedAt       mysql.NullTime `db:"updated_at"`
	DeletedAt       mysql.NullTime `db:"deleted_at"`
	IdCategory      uint32         `db:"id_category"`
}

// Connection is an interface for making queries.
type Connection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

func Create(db Connection, namesubcategory, idcategory string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
    INSERT INTO %v
    (name_sub_category, id_category)
    VALUES
    (?,?)
  `, table), namesubcategory, idcategory)
	return result, err
}
