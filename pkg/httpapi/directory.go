package httpapi

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetDirectory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Calling the directory!\n")
}
