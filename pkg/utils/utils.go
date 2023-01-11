package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/tony-skywalks/my-web/pkg/models"

	"github.com/joho/godotenv"
)

var DEBUG = false

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DEBUG = strinToBool(os.Getenv("DEBUG"))
}

func strinToBool(s string) bool {
	if strings.ToLower(s) == "true" || s == "1" {
		return true
	}
	return false
}

func ParseBody(r *http.Request, unmarshaled_body interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), unmarshaled_body); err == nil {
			return
		}
		err = json.Unmarshal([]byte(body), unmarshaled_body)
	}
}

func PP(a ...interface{}) {
	if DEBUG {
		fmt.Println(a)
	}
}

func SendResponse(w http.ResponseWriter, status string, msg string, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")

	var res models.Response
	res.Staus = status
	res.Error = msg
	res.Data = data

	response, _ := json.Marshal(res)
	w.WriteHeader(statusCode)
	w.Write(response)
}
