package main

import dt "github.com/DevOpsInvTech/doittypes"

//AddVar Add new var to datastore
func (ds *DoitServer) AddVar(d *dt.Domain, name string, value string) (v *dt.Var, err error) {
	v = &dt.Var{Name: name, Value: value, Domain: d}
	ds.Store.Conn.NewRecord(v)
	gormErr := ds.Store.Conn.Create(&v)
	return v, gormErr.Error
}

//UpdateVar Update Var
func (ds *DoitServer) UpdateVar(d *dt.Domain, id int, value string) error {
	v, err := ds.GetVar(d, id)
	if err != nil {
		return err
	}
	v.Value = value
	gormErr := ds.Store.Conn.Save(&v)
	if gormErr.Error != nil {
		return gormErr.Error
	}
	return nil
}

//RemoveVar Remove Var
func (ds *DoitServer) RemoveVar(d *dt.Domain, v *dt.Var) error {
	v, err := ds.GetVar(d, v.ID)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Delete(&v)
	if gormErr.Error != nil {
		return gormErr.Error
	}
	return nil
}

//GetVar Get Var from datastore
func (ds *DoitServer) GetVar(d *dt.Domain, id int) (*dt.Var, error) {
	v := &dt.Var{ID: id, Domain: d}
	gormErr := ds.Store.Conn.Where("id = ? and domain_id = ?", id, d.ID).First(&v)
	if gormErr.Error != nil {
		return v, gormErr.Error
	}
	return v, nil
}

//GetVarByName Get Var from datastore
func (ds *DoitServer) GetVarByName(d *dt.Domain, name string) (*dt.Var, error) {
	v := &dt.Var{Name: name, Domain: d}
	gormErr := ds.Store.Conn.Where("name = ? and domain_id = ?", name, d.ID).First(&v)
	if gormErr.Error != nil {
		return v, gormErr.Error
	}
	return v, nil
}

//GetVarsByDomain Get Vars from datastore
func (ds *DoitServer) GetVarsByDomain(d *dt.Domain) ([]*dt.Var, error) {
	vars := []*dt.Var{}
	gormErr := ds.Store.Conn.Where("domain_id = ? and ifnull(host_id, '') = '' and  ifnull(group_id, '') = ''", d.ID).Find(&vars)
	if gormErr.Error != nil {
		return vars, gormErr.Error
	}
	return vars, nil
}
