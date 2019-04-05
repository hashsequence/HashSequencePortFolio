package main

import (
  "./data"
  "fmt"
  "encoding/json"
  "log"
  "net/http"
)

func index(w http.ResponseWriter, request *http.Request) {
  fmt.Println("this is the index")
}


func GetUserData(w http.ResponseWriter, request *http.Request) {
  userJson, err := json.MarshalIndent(data.GetUser(1), "", "  ")
  if err != nil {
		log.Fatal("Error Getting User JSON", err)
	}
  fmt.Println("GetUserData: ",userJson)
  w.Header().Set("Access-Control-Allow-Origin", "*")
   w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
  w.Write(userJson)

}

func GetUserDataObj(w http.ResponseWriter, request *http.Request) {
  userJson, err := json.MarshalIndent(data.GetUserObj(1), "", "  ")
  if err != nil {
		log.Fatal("Error Getting User JSON", err)
	}
  fmt.Println("GetUserData: ",userJson)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
  w.Write(userJson)

}

func GetUserExperienceData(w http.ResponseWriter, request *http.Request) {
  userExpJsonArr, err := json.MarshalIndent(data.GetUserExperience(1), "", "  " )
  if err != nil {
		log.Fatal("Error Getting User Experience JSON", err)
	}
  w.Header().Set("Access-Control-Allow-Origin", "*")

   w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
  w.Write(userExpJsonArr)
}
