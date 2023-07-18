package main

import (
	"database/sql"
	"fmt"

	"github.com/cleitonbalonekr/go-intensivo/internal/infra/database"
	"github.com/cleitonbalonekr/go-intensivo/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close() // espera rodar e depois executa o close
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)
	input := usecase.OrderInput{
		ID:    "2",
		Price: 10.0,
		Tax:   1.0,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(*output)
}
