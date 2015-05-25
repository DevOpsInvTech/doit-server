package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	dt "github.com/DevOpsInvTech/doittypes"
	log "github.com/Sirupsen/logrus"
)

func (ds *DoitServer) ReturnBadRequest(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusBadRequest)
	ds.logger(r, http.StatusBadRequest, 0)
	return nil
}

func (ds *DoitServer) ReturnInternalServerError(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusInternalServerError)
	ds.logger(r, http.StatusInternalServerError, 0)
	return nil
}

func (ds *DoitServer) ReturnNotImplemented(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNotImplemented)
	ds.logger(r, http.StatusNotImplemented, 0)
	return nil
}

func (ds *DoitServer) ReturnOK(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	ds.logger(r, http.StatusOK, 0)
	return nil
}

func (ds *DoitServer) ReturnNotFound(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNotFound)
	ds.logger(r, http.StatusNotFound, 0)
	return nil
}

func (ds *DoitServer) ReturnJSON(dStruct interface{}, w http.ResponseWriter, r *http.Request) error {
	data, err := json.Marshal(dStruct)
	if err != nil {
		log.Errorln("Unable to marshal json", data)
		w.WriteHeader(http.StatusInternalServerError)
		ds.logger(r, http.StatusInternalServerError, 0)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	ds.logger(r, http.StatusOK, len(data))
	return nil
}

func (ds *DoitServer) ReturnEmptyJSON(w http.ResponseWriter, r *http.Request) error {
	data := "{}"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(data))
	ds.logger(r, http.StatusOK, len(data))
	return nil
}

func (ds *DoitServer) DomainCheck(dName string) (d *dt.Domain, err error) {
	if len(dName) > 0 {
		var err error
		d, err = ds.GetDomainByName(dName)
		if err != nil {
			return nil, err
		}
		return d, nil
	}
	return nil, errors.New("Domain string not valid")
}

func (ds *DoitServer) homeHandler(w http.ResponseWriter, r *http.Request) {

}

func (ds *DoitServer) logger(r *http.Request, status int, retSize int) {
	t := time.Now()
	zone, _ := t.Zone()
	log.Infof("%s %s %s [%s] \"%s %s %s\" %d %d", r.RemoteAddr, "-", "-", fmt.Sprintf("%d/%s/%d:%d:%d:%d %s", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second(), zone), r.Method, r.URL.RequestURI(), r.Proto, status, retSize)
}
