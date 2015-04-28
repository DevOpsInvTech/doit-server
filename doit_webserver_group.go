package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func (ds *DoitServer) apiGroupVarHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		w.WriteHeader(http.StatusInternalServerError)
		ds.logger(r, http.StatusInternalServerError, 0)
		return
	}
	vars := mux.Vars(r)
	domain := r.Form.Get("domain")
	reqName := vars["name"]
	varName := vars["varName"]

	d, err := ds.DomainCheck(domain)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ds.logger(r, http.StatusBadRequest, 0)
		return
	}

	switch r.Method {
	case "GET":
		g, err := ds.GetGroupByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnJSON(g, w, r)
		if err != nil {
			return
		}
	case "POST":
		_, err := ds.AddGroup(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnOK(w, r)
	case "PUT":
		//TODO: Add group items here
		w.WriteHeader(http.StatusNotImplemented)
		ds.logger(r, http.StatusNotImplemented, 0)
	case "DELETE":
		g, err := ds.GetGroupByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		err = ds.RemoveGroup(d, g)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ds.logger(r, http.StatusInternalServerError, 0)
			return
		}
		ds.ReturnOK(w, r)
	}
}

func (ds *DoitServer) apiGroupHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		w.WriteHeader(http.StatusInternalServerError)
		ds.logger(r, http.StatusInternalServerError, 0)
		return
	}
	vars := mux.Vars(r)
	domain := r.Form.Get("domain")
	reqName := vars["name"]

	d, err := ds.DomainCheck(domain)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ds.logger(r, http.StatusBadRequest, 0)
		return
	}

	switch r.Method {
	case "GET":
		g, err := ds.GetGroupByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnJSON(g, w, r)
		if err != nil {
			return
		}
	case "POST":
		_, err := ds.AddGroup(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnOK(w, r)
	case "PUT":
		//TODO: Add group items here
		w.WriteHeader(http.StatusNotImplemented)
		ds.logger(r, http.StatusNotImplemented, 0)
	case "DELETE":
		g, err := ds.GetGroupByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		err = ds.RemoveGroup(d, g)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ds.logger(r, http.StatusInternalServerError, 0)
			return
		}
		ds.ReturnOK(w, r)
	}
}

func (ds *DoitServer) apiGroupsHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		w.WriteHeader(http.StatusInternalServerError)
		ds.logger(r, http.StatusInternalServerError, 0)
		return
	}
	domain := r.Form.Get("domain")

	d, err := ds.DomainCheck(domain)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ds.logger(r, http.StatusBadRequest, 0)
		return
	}

	switch r.Method {
	case "GET":
		g, err := ds.GetGroupsByDomain(d)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnJSON(g, w, r)
		if err != nil {
			return
		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
		ds.logger(r, http.StatusNotImplemented, 0)
		return
	}
}
