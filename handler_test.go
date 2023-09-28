package fuzzing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func FuzzCalculateHighestValueHandler(f *testing.F) {
	server := httptest.NewServer(http.HandlerFunc(CalculateHighestValueHandler))
	defer server.Close()

	testCases := []ValuesRequest{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10}},
		{[]int{-1, -2, -3, -4, -5, 6, 7, 8, 9, 10}},
		{[]int{1}},
	}

	for _, testCase := range testCases {
		data, _ := json.Marshal(testCase)
		f.Add(data)
	}

	f.Fuzz(func(t *testing.T, data []byte) {
		if !json.Valid(data) {
			t.Skip()
		}

		var valuesRequest ValuesRequest
		if err := json.Unmarshal(data, &valuesRequest); err != nil {
			t.Skip()
		}

		resp, err := http.DefaultClient.Post(server.URL, "application/json", bytes.NewBuffer(data))
		if err != nil {
			t.Errorf("Error reaching API: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d instead", http.StatusOK, resp.StatusCode)
		}

		var response int
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Errorf("Error decoding response: %v", err)
		}
	})
}
