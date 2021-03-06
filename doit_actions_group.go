package main

import dt "github.com/DevOpsInvTech/doittypes"

//AddGroup Add group to datastore
func (ds *DoitServer) AddGroup(d *dt.Domain, name string) (g *dt.Group, err error) {
	domain, err := ds.GetDomain(d.ID)
	if err != nil {
		return g, err
	}
	g = &dt.Group{Name: name, Domain: domain}
	ds.Store.Conn.NewRecord(g)
	gormErr := ds.Store.Conn.Create(&g)
	return g, gormErr.Error
}

//AddGroupVars Add new Vars to Host
func (ds *DoitServer) AddGroupVars(d *dt.Domain, id int, vars ...*dt.Var) error {
	domain, err := ds.GetDomain(d.ID)
	if err != nil {
		return err
	}
	g, err := ds.GetGroup(domain, id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&g).Association("Vars").Append(&vars)
	return gormErr.Error
}

//RemoveGroupVars Remove Vars from Host
func (ds *DoitServer) RemoveGroupVars(d *dt.Domain, id int, vars ...*dt.Var) error {
	domain, err := ds.GetDomain(d.ID)
	if err != nil {
		return err
	}
	g, err := ds.GetGroup(domain, id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&g).Association("Vars").Delete(&vars)
	if gormErr != nil {
		return gormErr.Error
	}
	return nil
}

//AddGroupHosts Add new Host to Group
func (ds *DoitServer) AddGroupHosts(d *dt.Domain, id int, hosts ...*dt.Host) error {
	domain, err := ds.GetDomain(d.ID)
	if err != nil {
		return err
	}
	g, err := ds.GetGroup(domain, id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&g).Association("Hosts").Append(hosts)
	return gormErr.Error
}

//AddGroupHostVars Add new Vars to Host
func (ds *DoitServer) AddGroupHostVars(d *dt.Domain, g *dt.Group, h *dt.Host, vars ...*dt.Var) error {
	var err error
	h, err = ds.GetGroupHostByName(d, g, h.Name)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&h).Association("Vars").Append(vars)
	return gormErr.Error
}

//RemoveGroupHosts Remove Vars from Host
func (ds *DoitServer) RemoveGroupHosts(d *dt.Domain, id int, hosts ...dt.Host) error {
	domain, err := ds.GetDomain(d.ID)
	if err != nil {
		return err
	}
	g, err := ds.GetGroup(domain, id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&g).Association("Hosts").Delete(&hosts)
	if gormErr != nil {
		return gormErr.Error
	}
	return nil
}

//GetGroupHosts Get Group Hosts from datastore
func (ds *DoitServer) GetGroupHosts(d *dt.Domain, g *dt.Group) ([]*dt.Host, error) {
	v := []*dt.Host{}
	gormErr := ds.Store.Conn.Where("group_id = ? and domain_id = ?", g.ID, d.ID).Find(&v)
	if gormErr.Error != nil {
		return nil, gormErr.Error
	}
	return v, nil
}

//GetGroupHostByName Get GroupVar by name
func (ds *DoitServer) GetGroupHostByName(d *dt.Domain, g *dt.Group, name string) (*dt.Host, error) {
	h := &dt.Host{}
	gormErr := ds.Store.Conn.Where("name = ? and domain_id = ? and group_id = ?", name, d.ID, g.ID).First(&h)
	if gormErr.Error != nil {
		return nil, gormErr.Error
	}
	return h, nil
}

//GetGroupHostVars Get Host var by group
func (ds *DoitServer) GetGroupHostVars(d *dt.Domain, g *dt.Group, h *dt.Host) ([]*dt.Var, error) {
	v := []*dt.Var{}
	gormErr := ds.Store.Conn.Where("host_id = ? and domain_id = ? and group_id = ?", h.ID, d.ID, g.ID).Find(&v)
	if gormErr.Error != nil {
		return nil, gormErr.Error
	}
	return v, nil
}

//AddGroupDomains Add new Vars to Host
func (ds *DoitServer) AddGroupDomains(d *dt.Domain, id int, domains ...dt.Domain) error {
	domain, err := ds.GetDomain(d.ID)
	if err != nil {
		return err
	}
	g, err := ds.GetGroup(domain, id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&g).Association("Domains").Append(domains)
	return gormErr.Error
}

//RemoveGroupDomains Remove Vars from Host
func (ds *DoitServer) RemoveGroupDomains(d *dt.Domain, id int, domains ...dt.Domain) error {
	domain, err := ds.GetDomain(d.ID)
	if err != nil {
		return err
	}
	g, err := ds.GetGroup(domain, id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&g).Association("Domains").Delete(&domains)
	if gormErr != nil {
		return gormErr.Error
	}
	return nil
}

//RemoveGroup Remove group and its relationships to other objects
func (ds *DoitServer) RemoveGroup(d *dt.Domain, group *dt.Group) error {
	domain, err := ds.GetDomain(d.ID)
	if err != nil {
		return err
	}
	g, err := ds.GetGroup(domain, group.ID)
	if err != nil {
		return err
	}
	if len(g.Vars) > 0 {
		gormErr := ds.Store.Conn.Model(&g).Association("Vars").Delete(&g.Vars)
		if gormErr.Error != nil {
			return gormErr.Error
		}
	}
	if len(g.Hosts) > 0 {
		gormErr := ds.Store.Conn.Model(&g).Association("Hosts").Delete(&g.Hosts)
		if gormErr.Error != nil {
			return gormErr.Error
		}
	}
	hostErr := ds.Store.Conn.Delete(&g)
	if hostErr.Error != nil {
		return hostErr.Error
	}
	return nil
}

//GetGroup Get Var from datastore
func (ds *DoitServer) GetGroup(d *dt.Domain, id int) (*dt.Group, error) {
	g := &dt.Group{ID: id, Domain: d}
	gormErr := ds.Store.Conn.First(&g).Related(&g.Vars, "Vars").Related(&g.Hosts, "Hosts")
	if gormErr.Error != nil {
		return nil, gormErr.Error
	}
	return g, nil
}

//GetGroupByName Get host from datastore
func (ds *DoitServer) GetGroupByName(d *dt.Domain, name string) (*dt.Group, error) {
	g := &dt.Group{Name: name, Domain: d}
	gormErr := ds.Store.Conn.Where("name = ? and domain_id = ?", name, d.ID).First(&g).Related(&g.Vars, "Vars").Related(&g.Hosts, "Hosts")
	if gormErr.Error != nil {
		return nil, gormErr.Error
	}
	return g, nil
}

//GetGroupsByDomain Get Var from datastore
func (ds *DoitServer) GetGroupsByDomain(d *dt.Domain) ([]*dt.Group, error) {
	g := []*dt.Group{}
	gormErr := ds.Store.Conn.Find(&g).Where("domain_id = ?", d.ID)
	if gormErr.Error != nil {
		return nil, gormErr.Error
	}
	return g, nil
}

//GetGroupVar Get GroupVar from datastore
func (ds *DoitServer) GetGroupVar(d *dt.Domain, id int) (*dt.Var, error) {
	v := &dt.Var{}
	gormErr := ds.Store.Conn.Where("id = ? and domain_id = ?", id, d.ID).First(&v)
	if gormErr.Error != nil {
		return nil, gormErr.Error
	}
	return v, nil
}

//GetGroupVars Get GroupVars from datastore
func (ds *DoitServer) GetGroupVars(d *dt.Domain, g *dt.Group) ([]*dt.Var, error) {
	v := []*dt.Var{}
	gormErr := ds.Store.Conn.Where("group_id = ? and domain_id = ?", g.ID, d.ID).Find(&v)
	if gormErr.Error != nil {
		return nil, gormErr.Error
	}
	return v, nil
}

//GetGroupVarByName Get GroupVar by name
func (ds *DoitServer) GetGroupVarByName(d *dt.Domain, g *dt.Group, name string) (*dt.Var, error) {
	v := &dt.Var{}
	gormErr := ds.Store.Conn.Where("name = ? and domain_id = ? and group_id = ?", name, d.ID, g.ID).First(&v)
	if gormErr.Error != nil {
		return nil, gormErr.Error
	}
	return v, nil
}
