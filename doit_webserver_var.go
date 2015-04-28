package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func (ds *DoitServer) apiVarsHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		ds.ReturnInternalServerError(w, r)
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
		retVars, err := ds.GetVarsByDomain(d)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnJSON(retVars, w, r)
		if err != nil {
			return
		}
	default:
		ds.ReturnNotImplemented(w, r)
		return
	}
}

func (ds *DoitServer) apiVarHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		ds.ReturnInternalServerError(w, r)
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
		v, err := ds.GetVarByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnJSON(v, w, r)
		if err != nil {
			return
		}
	case "POST":
		_, err := ds.AddVar(d, reqName, "")
		if err != nil {
			//TODO: What error to throw here?
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnOK(w, r)
	case "PUT":
		v, err := ds.GetVarByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		err = ds.UpdateVar(d, v.ID, "")
		if err != nil {
			//TODO: WHAT TO RETURN HERE?
			ds.ReturnNotImplemented(w, r)
			return
		}
		ds.ReturnOK(w, r)
	case "DELETE":
		v, err := ds.GetVarByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		err = ds.RemoveVar(d, v)
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

func (ds *DoitServer) apiVarValueHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		ds.ReturnInternalServerError(w, r)
		return
	}
	vars := mux.Vars(r)
	domain := r.Form.Get("domain")
	reqValue := vars["value"]
	reqName := vars["name"]

	d, err := ds.DomainCheck(domain)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ds.logger(r, http.StatusBadRequest, 0)
		return
	}

	switch r.Method {
	case "GET":
		v, err := ds.GetVarByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnJSON(v, w, r)
		if err != nil {
			return
		}
	case "POST":
		_, err := ds.AddVar(d, reqName, reqValue)
		if err != nil {
			//TODO: What error to throw here?
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnOK(w, r)
	case "PUT":
		v, err := ds.GetVarByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		err = ds.UpdateVar(d, v.ID, reqValue)
		if err != nil {
			//TODO: WHAT TO RETURN HERE?
			ds.ReturnNotImplemented(w, r)
			return
		}
		ds.ReturnOK(w, r)
	case "DELETE":
		v, err := ds.GetVarByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		err = ds.RemoveVar(d, v)
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
