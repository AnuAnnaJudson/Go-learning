package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project-5/models"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	//set content type

	userIdString := r.URL.Query().Get("user_id") // this will fetch the

	//convert this to uint (strconv.ParseUint)
	userId, err := strconv.ParseUint(userIdString, 10, 64)

	// if error happens report to the client in json // write status code
	if err != nil {

		log.Println("Error: ", err)

		errorInConversion := map[string]string{"msg": "not a valid number"}

		jsonData, err := json.Marshal(errorInConversion)

		if err != nil {
			log.Println("Error while converting error to json", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
			return
		}

		w.WriteHeader(http.StatusBadRequest)

		w.Write(jsonData)

		return

	}

	//appErr := map[string]string{"Message": http.StatusText(http.StatusBadRequest)}

	//call the FetchUser function from the models package

	uData, err := models.FetchUser(userId)

	if err != nil {

		fetchError := map[string]string{"msg": "user not found"}

		errData, err := json.Marshal(fetchError)

		if err != nil {

			log.Println("Error while parsing fetchuser error conversion: ", err)

			w.WriteHeader(http.StatusInternalServerError)

			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))

			return

		}

		w.WriteHeader(http.StatusInternalServerError)

		w.Write(errData)

		return

	}

	// write to the client with user data or error

	userData, err := json.Marshal(uData)

	if err != nil {

		log.Println("Error while converting user data to json", err)

		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))

		return

	}

	w.Write(userData)

}

// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	//set content type
// 	w.Header().Set("Content-Type", "application/json")

// 	userIdString := r.URL.Query().Get("user_id")
// 	//convert this to uint (strconv.ParseUint)
// 	userId, err := strconv.ParseUint(userIdString, 10, 64)
// 	// if error happens report to the client in json // write status code
// 	if err != nil {
// 		conversionErr := map[string]string{"Message": "Enter numeric values"}
// 		conversionErrJson, err := json.Marshal(conversionErr)
// 		if err != nil {
// 			log.Println(err)

// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte(err.Error()))
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(conversionErrJson))
// 		return
// 	}
// 	//appErr := map[string]string{"Message": http.StatusText(http.StatusBadRequest)}

// 	//call the FetchUser function from the models package
// 	u, err := models.FetchUser(userId)
// 	if err != nil {
// 		conversionErr := map[string]string{"Message": err.Error()} //to get error from the function and show,
// 		//to convert err of type error to string do err.Error()
// 		conversionErrJson, err := json.Marshal(conversionErr)
// 		if err != nil {
// 			log.Println(err)

// 			w.WriteHeader(http.StatusBadRequest)
// 			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
// 			// w.Write([]byte(err.Error()))

// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(conversionErrJson))
// 		return
// 	}
// 	jsonData, err := json.Marshal(u)
// 	if err != nil {
// 		conversionErr := map[string]string{"Message": "Cannot convert to json"}
// 		conversionErrJson, err := json.Marshal(conversionErr)
// 		if err != nil {
// 			log.Println(err)

// 			w.WriteHeader(http.StatusBadRequest)
// 			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))

// 			// w.Write([]byte(err.Error()))
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(conversionErrJson))
// 		return
// 	}
// 	// write to the client with user data or error

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(jsonData))

// }
