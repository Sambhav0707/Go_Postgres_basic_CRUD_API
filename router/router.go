package router


import(
	"github.com/gorilla/mux"
	"go_postgres/middleware"
)


func Router() *mux.Router{
	r := mux.NewRouter()

	r.HandleFunc("/api/stocks/{id}" , middleware.GetStockById ).Methods("GET" , "OPTIONS")
	r.HandleFunc("/api/stocks" , middleware.GetStocks ).Methods("GET" , "OPTIONS")
	r.HandleFunc("/api/newstock" , middleware.CreateStock ).Methods("POST" , "OPTIONS")
	r.HandleFunc("/api/stocks/{id}" , middleware.UpdateStock ).Methods("PUT" , "OPTIONS")
	r.HandleFunc("/api/deletestock/{id}" , middleware.DeleteStock ).Methods("DELETE" , "OPTIONS")

	return r

}