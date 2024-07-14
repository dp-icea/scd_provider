package httpServer

import (
	"encoding/json"
	"log"
	"net/http"
	"scd_provider/config"
	"scd_provider/scd"
	"scd_provider/scd/dss"
	"strings"
)

type HttpServer struct {
	Deconflictor scd.StrategicDeconfliction
	verifier     scd.TokenVerifier
	conf         config.Config
}

type RequestParser interface {
	ParseInjection(r *http.Request) error
	ParseNotification(r *http.Request) error

	ParseFetchOirResponse(intent dss.OperationalIntent) dss.GetOperationalIntentDetailsResponse
}

func (h HttpServer) Serve() {
	h.conf = *config.GetGlobalConfig()
	h.verifier = scd.JwtTokenVerifier{}

	http.HandleFunc("/injection", h.handleInjection)
	http.HandleFunc("/uss/v1/operational_intents/{entityid}", h.handleFetchOir)
	http.HandleFunc("/uss/v1/operational_intents/", h.handleNotification)
	http.HandleFunc("/uss/v1/versions/{systemid}", h.handleFetchVersion)

	log.Println("Starting server on port :" + h.conf.HostPort)
	log.Fatal(http.ListenAndServe(":"+h.conf.HostPort, nil))
}

func (h HttpServer) handleInjection(w http.ResponseWriter, r *http.Request) {
	parser := HttpRequestParser{}
	log.Println(*r)

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
	log.Println(*r)
	parser := HttpRequestParser{}

	if r.Method != http.MethodGet {
		log.Print("Invalid request method: " + r.Method)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	 authorization, err := h.verifyAuthorization(r, scd.StrategicCoordination)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if !authorization {
		log.Print("Unauthorized")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
 
	id := r.PathValue("entityid")
	oir, err := h.Deconflictor.FetchOir(id)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response := parser.ParseFetchOirResponse(oir)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h HttpServer) handleNotification(w http.ResponseWriter, r *http.Request) {
	log.Println(*r)

	if r.Method != http.MethodPost {
		log.Print("Invalid request method: " + r.Method)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	//TODO Update OIR Status
}

func (h HttpServer) handleFetchVersion(w http.ResponseWriter, r *http.Request) {
	log.Println(*r)

	if r.Method != http.MethodGet {
		log.Print("Invalid request method: " + r.Method)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
 
	authorization, err := h.verifyAuthorization(r, scd.GetVersion)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if !authorization {
		log.Print("Unauthorized")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}   

	systemid := r.PathValue("systemid")
	response := h.Deconflictor.FetchVersion(systemid)
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	return
}

func (h HttpServer) verifyAuthorization(r *http.Request, scope scd.AuthScope) (bool, error) {
	token := r.Header.Get("Authorization")
	splitToken := strings.Split(token, "Bearer ")
	if len(splitToken) != 2 {
		log.Println("Invalid Authorization Token")
		return false, nil
	}
	token = splitToken[1]

	return h.verifier.Verify(token, scope)
}
