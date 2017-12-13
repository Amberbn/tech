package category

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

//define table
var (
	table = "categories"
)

type Item struct {
	ID           uint32         `db:"id"`
	NameCategory string         `db:"name_category"`
	CreatedAt    mysql.NullTime `db:"created_at"`
	UpdatedAt    mysql.NullTime `db:"updated_at"`
	DeletedAt    mysql.NullTime `db:"deleted_at"`
}

// Connection is an interface for making queries.
type Connection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

// query all data category
func All(db Connection) ([]Item, bool, error) {
	var result []Item
	err := db.Select(&result, fmt.Sprintf(`
				SELECT id, name_category, created_at, updated_at, deleted_at
		 		FROM %v
		 		WHERE deleted_at IS NULL
		 		`, table))
	return result, err == sql.ErrNoRows, err
}

// query sava category
func Create(db Connection, namecategory string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
    INSERT INTO %v
    (name_category)
    VALUES
    (?)
  `, table), namecategory)
	return result, err
}

//query data category by id
func ById(db Connection, ID string) (Item, bool, error) {
	result := Item{}
	err := db.Get(&result, fmt.Sprintf(`
		SELECT id, name_category, created_at, updated_at, deleted_at
		FROM %v
		WHERE id = ?
		AND deleted_at IS NULL
	`, table), ID)
	return result, err == sql.ErrNoRows, err
}
