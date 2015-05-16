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
		err = ds.AddGroupVars(d, g.ID, &dt.Var{Name: varName, Value: value, Domain: d, Group: g})
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

func (ds *DoitServer) apiGroupHostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		ds.ReturnInternalServerError(w, r)
		return
	}
	vars := mux.Vars(r)
	domain := r.Form.Get("domain")
	reqName := vars["name"]
	hostName := vars["hostName"]

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
		gv, err := ds.GetGroupHostByName(d, g, hostName)
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
		err = ds.AddGroupHosts(d, g.ID, &dt.Host{Name: hostName, Domain: d, Group: g})
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

//TODO: Fix this!
//apiGroupHostVarsHandler handles requests for group var
func (ds *DoitServer) apiGroupHostVarHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		ds.ReturnInternalServerError(w, r)
		return
	}
	vars := mux.Vars(r)
	domain := r.Form.Get("domain")
	reqName := vars["name"]
	hostName := vars["hostName"]
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
		gh, err := ds.GetGroupHostByName(d, g, hostName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnJSON(gh, w, r)
	case "POST":
		g, err := ds.GetGroupByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		gh, err := ds.GetGroupHostByName(d, g, hostName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		err = ds.AddGroupHostVars(d, g, gh, &dt.Var{Name: varName, Value: value, Domain: d, Host: gh, Group: g})
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

//apiGroupHostsHandler returns all hosts for a given group
func (ds *DoitServer) apiGroupHostVarsHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		ds.ReturnInternalServerError(w, r)
		return
	}
	vars := mux.Vars(r)
	domain := r.Form.Get("domain")
	reqName := vars["name"]
	hostName := vars["hostName"]

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
		h, err := ds.GetGroupHostByName(d, g, hostName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		gh, err := ds.GetGroupHostVars(d, g, h)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnJSON(gh, w, r)
	default:
		ds.ReturnNotImplemented(w, r)
		return
	}
}

//apiGroupHostsHandler returns all hosts for a given group
func (ds *DoitServer) apiGroupHostsHandler(w http.ResponseWriter, r *http.Request) {
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
		hv, err := ds.GetGroupHosts(d, g)
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

//apiGroupHandler Handles Group requests
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

//apiGroupsHandler returns all groups for a given domain
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
