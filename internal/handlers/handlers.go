package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"gitlab.com/zapirus/task/internal/models"
	"gitlab.com/zapirus/task/internal/service"
)

type Handler struct {
	service service.Service
	wg      sync.WaitGroup
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}
func (h *Handler) MonitoringChangeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	now := time.Now()

	resp := make(map[string]int)
	resp["Status"] = http.StatusOK

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Error(err)
		return
	}

	headers := r.Header

	auth := make(map[string]string)

	for header, values := range headers {
		for _, value := range values {
			auth[header] = value
		}
	}
	auths := auth["X-Tantum-Authorization"]

	data := models.ReqData{
		Headers: headers,
		Body:    body,
	}
	headersJSON, err := json.Marshal(data.Headers)
	if err != nil {
		logrus.Errorf("failed to convert headers to JSON: %s", err)
		return
	}

	bodyJSON, err := json.Marshal(data.Body)
	if err != nil {
		logrus.Errorf("failed to convert headers to JSON: %s", err)
		return
	}

	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		if err = h.service.Take(now, auths, headersJSON, bodyJSON); err != nil {
			logrus.Error(err)
			return
		}
	}()
	h.wg.Wait()

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		logrus.Errorf("Error %s", err)
		return
	}
	w.Write(jsonResp)
	return

}
