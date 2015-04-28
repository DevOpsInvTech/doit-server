package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func (ds *DoitServer) apiDomainHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		w.WriteHeader(http.StatusInternalServerError)
		ds.logger(r, http.StatusInternalServerError, 0)
		return
	}
	vars := mux.Vars(r)
	reqName := vars["name"]

	switch r.Method {
	case "GET":
		d, err := ds.GetDomainByName(reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnJSON(d, w, r)
		if err != nil {
			return
		}
	case "POST":
		_, err := ds.AddDomain(reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnOK(w, r)
	case "PUT":
		w.WriteHeader(http.StatusNotImplemented)
		ds.logger(r, http.StatusNotImplemented, 0)
	case "DELETE":
		d, err := ds.GetDomainByName(reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		err = ds.RemoveDomain(d)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ds.logger(r, http.StatusInternalServerError, 0)
			return
		}
		ds.ReturnOK(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		ds.logger(r, http.StatusNotImplemented, 0)
		return
	}
}

func (ds *DoitServer) apiDomainsHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		w.WriteHeader(http.StatusInternalServerError)
		ds.logger(r, http.StatusInternalServerError, 0)
		return
	}

	switch r.Method {
	case "GET":
		d, err := ds.GetDomains()
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnJSON(d, w, r)
		if err != nil {
			return
		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
		ds.logger(r, http.StatusNotImplemented, 0)
		return
	}
}
