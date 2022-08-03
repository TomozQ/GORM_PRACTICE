package my

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Migrate program.
func Migrate () {
	// gorm.OpenでデータベースをOpen
	db, er := gorm.Open("sqlite3", "data.sqlite3")	//("ドライバ名", "データベース名")
	fmt.Println("hogehoge")
	if er != nil {
		fmt.Println(er)
		return
	}

	defer db.Close()
	// 引数に指定したモデルがデータベースに反映される
	db.AutoMigrate(&User{}, &Group{}, &Post{}, &Comment{})	// 引数にはマイグレーションするモデルの構造体の値をポインタで渡す
}