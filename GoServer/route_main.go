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
  userJson, err := json.Marshal(data.GetUser(1))
  if err != nil {
		log.Fatal("Error Getting User JSON", err)
	}
  fmt.Println("GetUserData: ",userJson)
  w.Write(userJson)

}

func GetUserExperienceData(w http.ResponseWriter, request *http.Request) {
  userExpJsonArr, err := json.Marshal(data.GetUserExperience(1))
  if err != nil {
		log.Fatal("Error Getting User Experience JSON", err)
	}
  w.Write(userExpJsonArr)
}
