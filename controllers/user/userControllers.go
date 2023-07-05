package userController

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	models "oct.loc/models"
	userService "oct.loc/services/user"
)

func Store(w http.ResponseWriter, r *http.Request) {

	user := userService.User{}

	user.SetFirstname(r.FormValue("firstname"))
	user.SetSecondname(r.FormValue("secondname"))
	user.SetEmail(r.FormValue("email"))
	user.SetPhone(r.FormValue("phone"))
	user.SetPassword(r.FormValue("password"))

	id, _ := models.UserStore(user)

	resp := make(map[string]int64)
	resp["user"] = id
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)

}

func Show(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	u64, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		panic(err)
	}

	u32 := uint32(u64)

	user := models.UserGet(u32)

	jsonResp, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)

}

func Index(w http.ResponseWriter, r *http.Request) {

	users := models.GetUsers()

	jsonResp, err := json.Marshal(users)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://oct.front")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
