package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"transaction_ex1/internal/handler"
	"transaction_ex1/internal/repo"

	_ "github.com/lib/pq"
)

func main() {
	// Соединение с базой
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	productRepo := repo.NewProductRepo(db)
	userRepo := repo.NewUserRepo(db)
	orderRepo := repo.NewOrderRepo(db)
	orderUsecase := handler.NewOrderUsecase(productRepo, userRepo, orderRepo)

	http.HandleFunc("/", formHandler)
	http.HandleFunc("/make-order", makeOrderHandler(orderUsecase, db))

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./internal/template/form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func makeOrderHandler(uc *handler.OrderUsecase, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// получаем данные из формы
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userId, err := strconv.Atoi(r.Form.Get("userId"))
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		productId, err := strconv.Atoi(r.Form.Get("productId"))
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		processingTime, err := strconv.Atoi(r.Form.Get("processingTime"))
		if err != nil {
			http.Error(w, "Invalid processing time", http.StatusBadRequest)
			return
		}

		// стартуем транзакцию
		// можно сделать изящней, но тут схематично
		tx, err := db.Begin()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// вызываем usecase создания заказа
		err = uc.MakeOrder(userId, productId, processingTime)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// закрываем транзакцию
		err = tx.Commit()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Заказ успешно создан"))
	}
}
