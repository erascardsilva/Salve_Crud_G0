package dados

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

const dbdados = "./dados/arquivo.sql"

func Connect() {
	dados()
	dados().Close()

}

func dados() *sql.DB {
	if err := os.MkdirAll(filepath.Dir(dbdados), os.ModePerm); err != nil {
		panic(err)
	}
	if _, err := os.Stat(dbdados); os.IsNotExist(err) {
		createDatabase()
	}
	db, err := sql.Open("sqlite3", dbdados)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Conectado ao Banco")
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS data (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		Titulo TEXT,
		Conteudo TEXT
	)
	`)
	if err != nil {
		panic(err)
	}

	return db
}

func createDatabase() {
	file, err := os.Create(dbdados)
	if err != nil {
		panic(err)
	}
	file.Close()
	fmt.Println("Bando de dados criado", dbdados)

}
