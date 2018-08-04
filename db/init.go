package db

import (
    "github.com/jinzhu/gorm"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

var (
    Mysql *gorm.DB
)

func Database(connString string) gin.HandlerFunc {
    db, err := gorm.Open("mysql", connString)
    db.LogMode(true)
    // Error
    if err != nil {
        panic(err)
    }
    Mysql = db
    
    return func(c *gin.Context) {
        c.Set("DB", db)
        c.Next()
    }
}