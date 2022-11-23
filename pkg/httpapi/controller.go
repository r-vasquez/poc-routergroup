package httpapi

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetControllerCluster(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Calling the controller cluster!\n")
}

func GetControllerTopic(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Calling the controller topic!\n")
}

func GetControllerFeatures(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Calling the controller features!\n")
}
