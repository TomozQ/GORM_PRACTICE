package main

import (
	"gorm-practice/my"	// go modules初期化の際に指定したモジュール名がポイント。→ 今回の場合go mod init gorm-practice モジュール名/パッケージ名
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// main program
func main() {
	my.Migrate()
}