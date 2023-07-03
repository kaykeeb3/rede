package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os/exec"
)

type PredictionRequest struct {
    X []float32 `json:"x"`
}

type PredictionResponse struct {
    Y []float32 `json:"y"`
}

func predictHandler(w http.ResponseWriter, r *http.Request) {
    var request PredictionRequest
    err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    cmd := exec.Command("python", "predict.py")
    cmd.Stdin = bytes.NewBufferString(jsonString)
    output, err := cmd.Output()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var response PredictionResponse
    err = json.Unmarshal(output, &response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/predict", predictHandler)

    log.Println("Starting server on http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
