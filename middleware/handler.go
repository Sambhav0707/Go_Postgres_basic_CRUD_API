package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go_postgres/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	// Load environment variables from .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Something went wrong")

	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES"))

	if err != nil {
		panic(err)
	}

	errr := db.Ping()

	if errr != nil {
		log.Fatal(errr)
	}

	fmt.Print("connection successfull")

	return db

}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	// we will assign the Stock from request body

	var stock models.Stock

	// now decode the json to pass it to the db

	err := json.NewDecoder(r.Body).Decode(&stock)

	// check the err

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// calling the insert stock func

	InsertID := insertStock(stock)

	// return the response
	res := response{
		ID:      InsertID,
		Message: "Stock created successfully",
	}

	json.NewEncoder(w).Encode(res)

}

func GetStockById(w http.ResponseWriter, r *http.Request) {
	// Get the id from the URL
	params := mux.Vars(r)

	// convert the id string into int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Could not convert the json into int %v", err)

	}

	// get the stock by the getStock func

	stock, err := getStock(int64(id))

	if err != nil {
		log.Fatalf("Could not get the stock: %v", err)
	}

	json.NewEncoder(w).Encode(stock)

}

func GetStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStock()

	if err != nil {
		log.Fatalf("Unable to get all stock. %v", err)
	}

	// send all the stocks as response
	json.NewEncoder(w).Encode(stocks)

}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	// get the stockid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// create an empty stock of type models.Stock
	var stock models.Stock

	// decode the json request to stock
	err = json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call update stock to update the stock
	updatedRows := updateStock(int64(id), stock)

	// format the message string
	msg := fmt.Sprintf("Stock updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)

}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Could not convert the json into int %v", err)
	}

	deletedRows := deleteStock(int64(id))

	msg := fmt.Sprintf("Stock deleted successfully. Total rows/record affected %v", deletedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)

}

func insertStock(stock models.Stock) int64 {
	db := CreateConnection()
	defer db.Close()

	query := `INSERT INTO stocks (name , price , company) VALUES ($1 , $2 , $3) RETURNING stockid`

	var id int64

	err := db.QueryRow(query, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)
	return id

}

func getStock(id int64) (models.Stock, error) {
	// create the db connection
	db := CreateConnection()

	// close the db connection
	defer db.Close()

	// create a stock of models.Stock type
	var stock models.Stock

	// create a query
	query := `SELECT * FROM stocks WHERE stockid=$1`

	// get the id from the url
	row := db.QueryRow(query, id)

	// unmarshal the row object to stock
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty stock on error
	return stock, err

}

func getAllStock() ([]models.Stock, error) {
	// create the postgres db connection
	db := CreateConnection()

	// close the db connection
	defer db.Close()

	var stocks []models.Stock

	// create the select sql query
	sqlStatement := `SELECT * FROM stocks`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var stock models.Stock

		// unmarshal the row object to stock
		err = rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the stock in the stocks slice
		stocks = append(stocks, stock)

	}

	// return empty stock on error
	return stocks, err
}

// update stock in the DB
func updateStock(id int64, stock models.Stock) int64 {

	// create the postgres db connection
	db := CreateConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE stocks SET name=$2, price=$3, company=$4 WHERE stockid=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// delete stock in the DB
func deleteStock(id int64) int64 {

	// create the postgres db connection
	db := CreateConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM stocks WHERE stockid=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
