package main

import (
	dt "github.com/DevOpsInvTech/doittypes"
)

//GetAllByDomain Get all items by domain
func (ds *DoitServer) GetAllByDomain(d *dt.Domain) (*dt.Domain, error) {
	var err error
	d.Vars, err = ds.GetVarsByDomain(d)
	if err != nil {
		return d, err
	}
	d.Hosts, err = ds.GetHostsByDomain(d)
	if err != nil {
		return d, err
	}
	for i, h := range d.Hosts {
		hVars, err := ds.GetHostVars(d, h)
		if err != nil {
			return d, err
		}
		d.Hosts[i].Vars = hVars
	}
	d.Groups, err = ds.GetGroupsByDomain(d)
	if err != nil {
		return d, err
	}
	for i, g := range d.Groups {
		gVars, err := ds.GetGroupVars(d, g)
		if err != nil {
			return d, err
		}
		d.Groups[i].Vars = gVars
		gHosts, err := ds.GetGroupHosts(d, g)
		if err != nil {
			return d, err
		}
		d.Groups[i].Hosts = gHosts
		for ih, h := range d.Groups[i].Hosts {
			ghVars, err := ds.GetGroupHostVars(d, g, h)
			if err != nil {
				return d, err
			}
			d.Groups[i].Hosts[ih].Vars = ghVars
		}
	}
	return d, nil
}
