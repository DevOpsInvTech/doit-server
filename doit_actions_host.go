package main

import (
	"fmt"

	dt "github.com/DevOpsInvTech/doittypes"
)

//AddHost Add new host to the datastore
func (ds *DoitServer) AddHost(d *dt.Domain, name string) (h *dt.Host, err error) {
	domain, err := ds.GetDomain(d.ID)
	if err != nil {
		return h, err
	}
	h = &dt.Host{Name: name, Domain: domain}
	ds.Store.Conn.NewRecord(h)
	gormErr := ds.Store.Conn.Create(&h)
	return h, gormErr.Error
}

//AddHostVars Add new Vars to Host
func (ds *DoitServer) AddHostVars(d *dt.Domain, id int, vars ...*dt.Var) error {
	h, err := ds.GetHost(d, id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&h).Association("Vars").Append(vars)
	return gormErr.Error
}

//RemoveHostVars Remove Vars from Host
func (ds *DoitServer) RemoveHostVars(d *dt.Domain, id int, vars ...*dt.Var) error {
	h, err := ds.GetHost(d, id)
	if err != nil {
		return err
	}
	for i, v := range vars {
		fmt.Println(i)
		rmVar, err := ds.GetHostVar(d, v.ID)
		fmt.Println(rmVar)
		if err != nil {
			return err
		}
		varErr := ds.Store.Conn.Delete(&rmVar)
		if varErr.Error != nil {
			return varErr.Error
		}
	}
	gormErr := ds.Store.Conn.Model(&h).Association("Vars").Delete(&vars)
	if gormErr != nil {
		return gormErr.Error
	}
	return nil
}

//RemoveHost Remove host from datastore
func (ds *DoitServer) RemoveHost(d *dt.Domain, host *dt.Host) error {
	h, err := ds.GetHost(d, host.ID)
	if err != nil {
		return err
	}
	if len(h.Vars) > 0 {
		gormErr := ds.Store.Conn.Model(&h).Association("Vars").Delete(&h.Vars)
		if gormErr.Error != nil {
			return gormErr.Error
		}
	}
	hostErr := ds.Store.Conn.Delete(&h)
	if hostErr.Error != nil {
		return hostErr.Error
	}
	return nil
}

//GetHost Get host from datastore
func (ds *DoitServer) GetHost(d *dt.Domain, id int) (*dt.Host, error) {
	h := &dt.Host{ID: id, Domain: d}
	gormErr := ds.Store.Conn.First(&h).Related(&h.Vars, "Vars")
	if gormErr.Error != nil {
		return nil, gormErr.Error
	}
	return h, nil
}

//GetHostByName Get host from datastore
func (ds *DoitServer) GetHostByName(d *dt.Domain, name string) (*dt.Host, error) {
	h := &dt.Host{Name: name, Domain: d}
	gormErr := ds.Store.Conn.Where("name = ? and domain_id = ?", name, d.ID).Find(&h).Related(&h.Vars, "Vars")
	if gormErr.Error != nil {
		return nil, gormErr.Error
	}
	return h, nil
}

//GetHostsByDomain Get host from datastore
func (ds *DoitServer) GetHostsByDomain(d *dt.Domain) ([]*dt.Host, error) {
	h := []*dt.Host{}
	gormErr := ds.Store.Conn.Where("domain_id = ? and ifnull(id, '') = '' and  ifnull(group_id, '') = ''", d.ID).Find(&h)
	if gormErr.Error != nil {
		return nil, gormErr.Error
	}
	return h, nil
}

//GetHostVar Get HostVar from datastore
func (ds *DoitServer) GetHostVar(d *dt.Domain, id int) (*dt.Var, error) {
	v := &dt.Var{}
	gormErr := ds.Store.Conn.Where("id = ? and domain_id = ?", id, d.ID).First(&v)
	if gormErr.Error != nil {
		return nil, gormErr.Error
	}
	return v, nil
}

//GetHostVars Get HostVar from datastore
func (ds *DoitServer) GetHostVars(d *dt.Domain, h *dt.Host) ([]*dt.Var, error) {
	v := []*dt.Var{}
	gormErr := ds.Store.Conn.Where("host_id = ? and domain_id = ?", h.ID, d.ID).Find(&v)
	if gormErr.Error != nil {
		return nil, gormErr.Error
	}
	return v, nil
}

//GetHostVarByName Get host var by name
func (ds *DoitServer) GetHostVarByName(d *dt.Domain, h *dt.Host, name string) (*dt.Var, error) {
	v := &dt.Var{}
	gormErr := ds.Store.Conn.Where("name = ? and domain_id = ? and host_id = ?", name, d.ID, h.ID).First(&v)
	if gormErr.Error != nil {
		return nil, gormErr.Error
	}
	return v, nil
}
