package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)

var DB *sql.DB

func ConnectDB() error {	
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/api")
	if err != nil {
       fmt.Println("Erro ao conectar com o banco de dados MySQL")
    }
    DB = db

    // Testar a conexão com o banco de dados
    err = DB.Ping()
    if err != nil {
        fmt.Println("Erro ao testar a conexão com o banco de dados MySQL")
        return err
    }

    fmt.Println("Conexão com o banco de dados MySQL estabelecida com sucesso!")
    return nil
}

// CloseDB fecha a conexão com o banco de dados
func CloseDB() {
    DB.Close()
}
