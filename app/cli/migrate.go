import (
    "github.com/jinzhu/gorm"
    "app/config"
    _ "github.com/go-sql-driver/mysql"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/mysql"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
    // db, err := sql.Open("postgres", "postgres://localhost:5432/database?sslmode=enable")
    db, err = gorm.Open("mysql", config.GetEnv().Database.FormatDSN())
    driver, err := mysql.WithInstance(db, &mysql.Config{})
    m, err := migrate.NewWithDatabaseInstance(
        "file:///app/migrations",
        "mysql", driver)
    m.Steps(2)
}