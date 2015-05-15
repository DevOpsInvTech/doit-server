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
	ds.Store.InitSchema(true)
	r := mux.NewRouter()

	//domains
	r.HandleFunc("/api/v1/domain/{name}", ds.apiDomainHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/v1/domains", ds.apiDomainsHandler).Methods("GET")
	//vars
	r.HandleFunc("/api/v1/var/{name}/value/{value}", ds.apiVarValueHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/v1/var/{name}", ds.apiVarHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/v1/vars", ds.apiVarsHandler).Methods("GET")
	//groups
	r.HandleFunc("/api/v1/group/{name}/var/{varName}/value/{value}", ds.apiGroupVarHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/v1/group/{name}/vars", ds.apiGroupVarsHandler).Methods("GET")
	r.HandleFunc("/api/v1/group/{name}/host/{hostName}", ds.apiGroupHostHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/v1/group/{name}/hosts", ds.apiGroupHostsHandler).Methods("GET")
	r.HandleFunc("/api/v1/group/{name}", ds.apiGroupHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/v1/groups", ds.apiGroupsHandler).Methods("GET")
	//hosts
	r.HandleFunc("/api/v1/host/{name}/var/{varName}/value/{value}", ds.apiHostVarHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/v1/host/{name}/vars", ds.apiHostVarsHandler).Methods("GET")
	r.HandleFunc("/api/v1/host/{name}", ds.apiHostHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/v1/hosts", ds.apiHostsHandler).Methods("GET")
	//All items in a domain
	r.HandleFunc("/api/v1/all", nil).Methods("GET")

	//home
	r.HandleFunc("/", ds.homeHandler)
	//templates
	r.HandleFunc("/api/v1/template/{name}", nil)
	r.HandleFunc("/api/v1/templates", nil)
	//object
	r.HandleFunc("/api/v1/object/{name}", nil)
	r.HandleFunc("/api/v1/objects", nil)
	//ansible
	r.HandleFunc("/api/ansible/domain/{name}", ds.ansibleHandler).Methods("GET")
	//handle root requests
	http.Handle("/", r)

	log.Infoln("Staring webserver")
	if err := http.ListenAndServe(net.JoinHostPort("", *port), nil); err != nil {
		log.Errorln(err)
		return err
	}
	return err
}
