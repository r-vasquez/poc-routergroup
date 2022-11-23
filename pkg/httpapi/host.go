package httpapi

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetHostCPU(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Calling the host CPU!\n")
}

func GetHostBuses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Calling the host buses!\n")
}

func GetHostMemory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Calling the host memory!\n")
}
