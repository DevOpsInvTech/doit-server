package main

import (
	"net/http"

	dt "github.com/DevOpsInvTech/doittypes"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func (ds *DoitServer) apiGroupVarHandler(w http.ResponseWriter, r *http.Request) {
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
		g, err := ds.GetGroupByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		gv, err := ds.GetGroupVarByName(d, g, varName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnJSON(gv, w, r)
	case "POST":
		g, err := ds.GetGroupByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		err = ds.AddGroupVars(d, g.ID, &dt.GroupVar{Name: varName, Value: value, Domain: d, Group: g})
		if err != nil {
			//TODO: What error to throw here?
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnOK(w, r)
	case "PUT":
		//TODO: Add group items here
		ds.ReturnNotImplemented(w, r)
	case "DELETE":
		g, err := ds.GetGroupByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		err = ds.RemoveGroupVars(d, g.ID)
		if err != nil {
			ds.ReturnInternalServerError(w, r)
			return
		}
		ds.ReturnOK(w, r)
	default:
		ds.ReturnNotImplemented(w, r)
	}
}

func (ds *DoitServer) apiGroupVarsHandler(w http.ResponseWriter, r *http.Request) {
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
		g, err := ds.GetGroupByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		hv, err := ds.GetGroupVars(d, g)
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

func (ds *DoitServer) apiGroupHandler(w http.ResponseWriter, r *http.Request) {
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
		ds.ReturnNotFound(w, r)
	case "DELETE":
		g, err := ds.GetGroupByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		err = ds.RemoveGroup(d, g)
		if err != nil {
			ds.ReturnInternalServerError(w, r)
			return
		}
		ds.ReturnOK(w, r)
	}
}

func (ds *DoitServer) apiGroupsHandler(w http.ResponseWriter, r *http.Request) {
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
		ds.ReturnNotImplemented(w, r)
		return
	}
}
