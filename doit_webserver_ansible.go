package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
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
		err = ds.ReturnJSON(ad, w, r)
		if err != nil {
			return
		}
	default:
		ds.ReturnNotImplemented(w, r)
		return
	}
}
