package main

import (
	"net/http"

	dt "github.com/DevOpsInvTech/doittypes"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func (ds *DoitServer) apiHostVarHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		ds.ReturnInternalServerError(w, r)
		return
	}
	vars := mux.Vars(r)
	domain := r.Form.Get("domain")
	reqName := vars["name"]
	varName := vars["varName"]
	value := vars["value"]

	d, err := ds.DomainCheck(domain)
	if err != nil {
		ds.ReturnBadRequest(w, r)
		return
	}

	switch r.Method {
	case "GET":
		h, err := ds.GetHostByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		hv, err := ds.GetHostVarByName(d, h, varName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnJSON(hv, w, r)
	case "POST":
		h, err := ds.GetHostByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		err = ds.AddHostVars(d, h.ID, &dt.HostVar{Name: varName, Value: value, Domain: d, Host: h})
		if err != nil {
			//TODO: What error to throw here?
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnOK(w, r)
	case "PUT":
		//TODO: Add host items here
		ds.ReturnNotImplemented(w, r)
	case "DELETE":
		h, err := ds.GetHostByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		err = ds.RemoveHostVars(d, h.ID, &dt.HostVar{Name: varName, Value: value, Domain: d})
		if err != nil {
			//TODO: What error to throw here?
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnOK(w, r)
	default:
		ds.ReturnNotImplemented(w, r)
		return
	}
}

func (ds *DoitServer) apiHostHandler(w http.ResponseWriter, r *http.Request) {
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
		ds.ReturnBadRequest(w, r)
		return
	}

	switch r.Method {
	case "GET":
		h, err := ds.GetHostByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnJSON(h, w, r)
		if err != nil {
			return
		}
	case "POST":
		_, err := ds.AddHost(d, reqName)
		if err != nil {
			//TODO: What error to throw here?
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnOK(w, r)
	case "PUT":
		//TODO: Add host items here
		ds.ReturnNotImplemented(w, r)
	case "DELETE":
		h, err := ds.GetHostByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
		}
		err = ds.RemoveHost(d, h)
		if err != nil {
			ds.ReturnInternalServerError(w, r)
		}
		ds.ReturnOK(w, r)
	default:
		ds.ReturnNotImplemented(w, r)
		return
	}
}

func (ds *DoitServer) apiHostsHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		ds.ReturnInternalServerError(w, r)
		return
	}
	domain := r.Form.Get("domain")

	d, err := ds.DomainCheck(domain)
	if err != nil {
		ds.ReturnBadRequest(w, r)
		return
	}

	switch r.Method {
	case "GET":
		h, err := ds.GetHostsByDomain(d)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnJSON(h, w, r)
		if err != nil {
			return
		}
	default:
		ds.ReturnNotImplemented(w, r)
		return
	}
}

func (ds *DoitServer) apiHostVarsHandler(w http.ResponseWriter, r *http.Request) {
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
		ds.ReturnBadRequest(w, r)
		return
	}

	switch r.Method {
	case "GET":
		h, err := ds.GetHostByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		hv, err := ds.GetHostVars(d, h)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnJSON(hv, w, r)
		if err != nil {
			return
		}
	default:
		ds.ReturnNotImplemented(w, r)
		return
	}
}
