package database

import (
	"database/sql"
	"fmt"
	chat_models "go-chat/models/chat"
	login_models "go-chat/models/login"
	"log"

	_ "github.com/lib/pq"
)

func OpenConn(dbName, dbUser, dbPassword, dbHost, dbPort string) (*sql.DB, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable", dbHost, dbPort, dbUser, dbPassword)
	postgres, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = postgres.Ping()
	if err != nil {
		panic(err)
	}
	defer postgres.Close()

	err = createDatabaseIfNotExists(dbName, dbUser, dbPassword, dbHost, dbPort)
	if err != nil {
		log.Fatalf("Database creating failed: %v\n", err)
	}

	connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = createTablesIfNotExists(db)
	if err != nil {
		log.Fatalf("Database tables creating failed: %v\n", err)
	}

	return db, err

}

func createDatabaseIfNotExists(dbName, dbUser, dbPassword, dbHost, dbPort string) error {

	// Para verificar, preciso conectar ao banco padrao do postgres
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable", dbHost, dbPort, dbUser, dbPassword)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	//defer db.Close()

	// Verifica se o banco de dados já existe
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s');", dbName)
	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		return err
	}

	// Se o banco de dados não existir, cria-o
	if !exists {
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
		if err != nil {
			return fmt.Errorf("error creating database: %v", err)
		}
		fmt.Printf("Database %s created successfully.\n", dbName)
	} else {
		fmt.Printf("Database %s already exists.\n", dbName)
	}

	return nil
}

func createTablesIfNotExists(db *sql.DB) error {
	// Define a query para criar a tabela, se ela não existir

	databases := []string{
		chat_models.CreateTableUsers(),
		login_models.CreateTableLogins(),
	}

	// Executa a query

	for _, createTableQuery := range databases {

		err := db.Ping()
		if err != nil {
			log.Fatalf("Error creating table: %v", err)
		}

		_, err = db.Exec(createTableQuery)
		if err != nil {
			return fmt.Errorf("error creating table: %v", err)
		}

		fmt.Println("Table created successfully.")
	}

	return nil
}
