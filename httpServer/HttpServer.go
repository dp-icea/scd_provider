package httpServer

import (
	"encoding/json"
	"icea_uss/scd"
	"log"
	"net/http"
)

type HttpServer struct {
}

type RequestParser interface {
	ParseInjection(r *http.Request) error
	ParseNotification(r *http.Request) error
}

func (h HttpServer) Serve() {
	parser := HttpRequestParser{}
	http.HandleFunc("/injection", func(w http.ResponseWriter, r *http.Request) {
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

		response, err := scd.Deconflictor{}.Injection(request)
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

	})

	http.HandleFunc("/uss/v1/operational_intents/{entityid}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Print("Invalid request method: " + r.Method)
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		//TODO Return OIR
	})

	http.HandleFunc("/uss/v1/operational_intents/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Print("Invalid request method: " + r.Method)
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		//TODO Update OIR Status
	})

	log.Fatal(http.ListenAndServe(":9091", nil))
}
