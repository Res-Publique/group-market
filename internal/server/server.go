package server

import (
	"net/http"

	"github.com/Res-Publique/group-market/internal/data"
	"github.com/sirupsen/logrus"
)

func setCORS(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Max-Age", "300")
}

func Initialize(db data.Database, w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodOptions {
		setCORS(w, req)
		w.WriteHeader(http.StatusOK)
		return
	}

	if method := req.Method; method != http.MethodPost {
		logrus.Info("Got ", method, " - abort request")
		http.Error(w, "/init is POST handle", http.StatusBadRequest)
		return
	}

	if content := req.Header.Get("Content-type"); content != "application/json" {
		logrus.Info("Got content-type ", content, " - abort request")
		http.Error(w, "Content-type is not application/json", http.StatusBadRequest)
		return
	}

	defer req.Body.Close()

	setCORS(w, req)
	w.WriteHeader(http.StatusOK)
}
