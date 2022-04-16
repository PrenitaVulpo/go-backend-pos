package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Event struct {
	Id          int    `json:"Id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type allEvents []Event

var events = allEvents{
	{
		Id:          1,
		Title:       "APIs",
		Description: "Descrição",
	},
}

// @title           Swagger API-Event
// @version         1.0
// @description     Documentação da API de eventos.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API-Event Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      https://go-events-pos.herokuapp.com
func main() {
	log.Println("Starting API")
	router := mux.NewRouter()
	router.HandleFunc("/", Home)
	router.HandleFunc("/health-check", Healtcheck).Methods("GET")
	router.HandleFunc("/events", GetAllEvents).Methods("GET")
	// port := os.Getenv("PORT")
	http.ListenAndServe(":8082", router)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Aplicação em execução!\n")
}

func Healtcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Aplicação em execução!\n")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Wornking\n")
}

// ShowAllEvents godoc
// @Summary      Show all events
// @Description  List all events
// @Tags         events
// @Accept       json
// @Produce      json
// @Success      200  {object}  Event
// @Router       /events [get]
func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}
