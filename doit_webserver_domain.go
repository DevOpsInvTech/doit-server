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
		ds.ReturnInternalServerError(w, r)
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
		err = ds.ReturnJSON(d, w, r)
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
		ds.ReturnNotFound(w, r)
	case "DELETE":
		d, err := ds.GetDomainByName(reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		err = ds.RemoveDomain(d)
		if err != nil {
			ds.ReturnInternalServerError(w, r)
			return
		}
		ds.ReturnOK(w, r)
	default:
		ds.ReturnNotImplemented(w, r)
		return
	}
}

func (ds *DoitServer) apiDomainsHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		ds.ReturnInternalServerError(w, r)
		return
	}

	switch r.Method {
	case "GET":
		d, err := ds.GetDomains()
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		err = ds.ReturnJSON(d, w, r)
		if err != nil {
			return
		}
	default:
		ds.ReturnNotFound(w, r)
		return
	}
}
