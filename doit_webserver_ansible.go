package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func (ds *DoitServer) ansibleGroupHandler(w http.ResponseWriter, r *http.Request) {
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
		a, err := ds.GetAllByDomain(d)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ad := a.MarshalAnsible()
		ag := ad["groups"]
		gr := make(map[string]interface{})
		for _, v := range ag.([]map[string]interface{}) {
			for k, va := range v {
				gr[k] = va
			}
		}
		err = ds.ReturnJSON(gr, w, r)
		if err != nil {
			return
		}
	default:
		ds.ReturnNotImplemented(w, r)
		return
	}
}

func (ds *DoitServer) ansibleHostVarsHandler(w http.ResponseWriter, r *http.Request) {
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
		meta := make(map[string]interface{})
		hostVars := make(map[string]interface{})
		hosts := make(map[string]interface{})
		for i := range h {
			h[i].Vars, _ = ds.GetHostVars(d, h[i])
			ha := h[i].MarshalAnsible()
			if vars, ok := ha[h[i].Name]; ok {
				hosts[h[i].Name] = vars
			}
		}
		hostVars["hostvars"] = hosts
		meta["_meta"] = hostVars
		err = ds.ReturnJSON(meta, w, r)
		if err != nil {
			log.Println(err)
			ds.ReturnEmptyJSON(w, r)
			return
		}
		ds.ReturnEmptyJSON(w, r)
	default:
		ds.ReturnNotImplemented(w, r)
		return
	}
}

func (ds *DoitServer) ansibleHostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		ds.ReturnInternalServerError(w, r)
		return
	}
	vars := mux.Vars(r)
	reqName := vars["name"]
	domain := r.Form.Get("domain")

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
		h.Vars, err = ds.GetHostVars(d, h)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		log.Println(h)
		ah := h.MarshalAnsible()
		log.Println(ah)
		if vars, ok := ah[h.Name]; ok {
			varsm := make(map[string]interface{})
			vs := vars.(map[string]interface{})
			for _, v := range vs["vars"].([]map[string]interface{}) {
				for k, vv := range v {
					varsm[k] = vv
				}
			}
			err = ds.ReturnJSON(varsm, w, r)
			if err != nil {
				ds.ReturnEmptyJSON(w, r)
				return
			}
			return
		}
		ds.ReturnEmptyJSON(w, r)
	default:
		ds.ReturnNotImplemented(w, r)
		return
	}
}
