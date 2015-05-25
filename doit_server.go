package main

import (
	"net"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

//DoitServer A webserver to frontend a DOIT database
type DoitServer struct {
	Store *DoitStorage
}

//OpenDatastore open datastore for writing
func (ds *DoitServer) OpenDatastore(t string, loc string) (err error) {
	s, err := NewStorage(t, loc)
	ds.Store = s
	return err
}

//CloseDatastore close datastore
func (ds *DoitServer) CloseDatastore() error {
	err := ds.Store.Close()
	return err
}

//Listen Starts the webserver for listening
func (ds *DoitServer) Listen(port *string, config *DoitConfig) (err error) {
	err = ds.OpenDatastore(config.Storage.Type, config.Storage.Location)
	if err != nil {
		return err
	}
	ds.Store.InitSchema(false)
	r := mux.NewRouter()

	//ansible
	r.HandleFunc("/api/v1/ansible/groups", ds.ansibleGroupHandler).Methods("GET")
	r.HandleFunc("/api/v1/ansible/host/{name}", ds.ansibleGroupHandler).Methods("GET")

	//APIv1 Subrouter
	apiV1R := r.PathPrefix("/api/v1").Subrouter()
	//All items in a domain
	apiV1R.HandleFunc("/all", ds.apiGetAllByDomain).Methods("GET")

	//domains
	apiV1R.HandleFunc("/domain/{name}", ds.apiDomainHandler).Methods("POST", "DELETE", "PUT", "GET")
	apiV1R.HandleFunc("/domains", ds.apiDomainsHandler).Methods("GET")
	//vars
	apiV1R.HandleFunc("/var/{name}/value/{value}", ds.apiVarValueHandler).Methods("POST", "DELETE", "PUT", "GET")
	apiV1R.HandleFunc("/var/{name}", ds.apiVarHandler).Methods("POST", "DELETE", "PUT", "GET")
	apiV1R.HandleFunc("/vars", ds.apiVarsHandler).Methods("GET")
	//hosts
	apiV1R.HandleFunc("/host/{name}/var/{varName}/value/{value}", ds.apiHostVarHandler).Methods("POST", "DELETE", "PUT", "GET")
	apiV1R.HandleFunc("/host/{name}/vars", ds.apiHostVarsHandler).Methods("GET")
	apiV1R.HandleFunc("/host/{name}", ds.apiHostHandler).Methods("POST", "DELETE", "PUT", "GET")
	apiV1R.HandleFunc("/hosts", ds.apiHostsHandler).Methods("GET")
	//groups
	apiV1R.HandleFunc("/group/{name}/var/{varName}/value/{value}", ds.apiGroupVarHandler).Methods("POST", "DELETE", "PUT", "GET")
	apiV1R.HandleFunc("/group/{name}/vars", ds.apiGroupVarsHandler).Methods("GET")
	apiV1R.HandleFunc("/group/{name}/host/{hostName}", ds.apiGroupHostHandler).Methods("POST", "DELETE", "PUT", "GET")
	apiV1R.HandleFunc("/group/{name}/hosts", ds.apiGroupHostsHandler).Methods("GET")
	apiV1R.HandleFunc("/group/{name}/host/{hostName}/var/{varName}/value/{value}", ds.apiGroupHostVarHandler).Methods("POST", "DELETE", "PUT", "GET")
	apiV1R.HandleFunc("/group/{name}/host/{hostName}/vars", ds.apiGroupHostVarsHandler).Methods("GET")
	apiV1R.HandleFunc("/group/{name}", ds.apiGroupHandler).Methods("POST", "DELETE", "PUT", "GET")
	apiV1R.HandleFunc("/groups", ds.apiGroupsHandler).Methods("GET")
	//templates
	apiV1R.HandleFunc("/template/{name}", nil)
	apiV1R.HandleFunc("/templates", nil)
	//object
	apiV1R.HandleFunc("/object/{name}", nil)
	apiV1R.HandleFunc("/objects", nil)

	//home
	r.HandleFunc("/", ds.homeHandler)

	//handle root requests
	http.Handle("/", r)

	log.Infoln("Staring webserver")
	if err := http.ListenAndServe(net.JoinHostPort("", *port), nil); err != nil {
		log.Errorln(err)
		return err
	}
	return err
}
