package httpServer

import (
	"encoding/json"
	"log"
	"net/http"
	"scd_provider/config"
	"scd_provider/scd"
)

type HttpServer struct {
	Deconflictor scd.StrategicDeconfliction
}

type RequestParser interface {
	ParseInjection(r *http.Request) error
	ParseNotification(r *http.Request) error
}

func (h HttpServer) Serve() {
	conf := config.GetGlobalConfig()

	http.HandleFunc("/injection", h.handleInjection)
	http.HandleFunc("/uss/v1/operational_intents/{entityid}", h.handleFetchOir)
	http.HandleFunc("/uss/v1/operational_intents/", h.handleNotification)

	log.Fatal(http.ListenAndServe(":"+conf.HostPort, nil))
}

func (h HttpServer) handleInjection(w http.ResponseWriter, r *http.Request) {
	parser := HttpRequestParser{}
	if r.Method != http.MethodPut {
		log.Print("Invalid request method: " + r.Method)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	request, err := parser.ParseInjection(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := h.Deconflictor.Inject(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	return
}

func (h HttpServer) handleFetchOir(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Print("Invalid request method: " + r.Method)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.PathValue("entityid")
	response, err := h.Deconflictor.FetchOir(id)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h HttpServer) handleNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Print("Invalid request method: " + r.Method)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	//TODO Update OIR Status
}
