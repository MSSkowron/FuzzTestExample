package fuzzing

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ValuesRequest struct {
	Values []int `json:"values"`
}

func CalculateHighestValueHandler(w http.ResponseWriter, r *http.Request) {
	var valuesRequest ValuesRequest
	if err := json.NewDecoder(r.Body).Decode(&valuesRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	highestValue := 0
	for _, value := range valuesRequest.Values {
		if value > highestValue {
			highestValue = value
		}
	}

	if highestValue == 100 {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Something went wrong"))
		return
	}

	_, _ = fmt.Fprintf(w, "%d", highestValue)
}
