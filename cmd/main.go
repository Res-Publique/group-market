package main

import (
	"net/http"

	"github.com/Res-Publique/group-market/internal/data"
	"github.com/Res-Publique/group-market/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := data.CreateDatabase()
	if err != nil {
		logrus.Fatal(err)
	}

	http.HandleFunc("/init", func(w http.ResponseWriter, req *http.Request) {
		server.Initialize(db, w, req)
	})

	err = http.ListenAndServe(":3000", nil)
	logrus.Fatal(err)
}
