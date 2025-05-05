
package db

import (
    "fmt"
    "log"
    "os"
    "persona-core/model"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Init() *gorm.DB {
    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        host, user, password, dbname, port)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to database: ", err)
    }

    db.AutoMigrate(&model.User{})
    return db
}
