package main

import (
"net/http"
"time"
)

func main() {
 	mux := http.NewServeMux()
  p("HashSequencePortfolio", version(), "started at", config.Address)
  // index
	mux.HandleFunc("/", index)
  mux.HandleFunc("/GetUserData", GetUserData)
  mux.HandleFunc("/GetUserDataObj", GetUserDataObj)
  mux.HandleFunc("/GetUserExperienceData", GetUserExperienceData)

  server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()

}
