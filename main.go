package main

import (
	"database/sql"
	"fmt"
	"github.com/JacquesMarques/go-intensivo-jul/internal/infra/database"
	"github.com/JacquesMarques/go-intensivo-jul/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	orderRepository := &database.OrderRepository{
		Db: db,
	}
	uc := usecase.NewCalculateFinalPrince(orderRepository)

	input := usecase.OrderInput{
		ID:    "123",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
