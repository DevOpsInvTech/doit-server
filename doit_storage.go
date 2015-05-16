package main

import (
	dt "github.com/DevOpsInvTech/doittypes"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

//DoitStorage Storage manager
type DoitStorage struct {
	Type     string
	Location string
	Conn     gorm.DB
}

//NewStorage Create a new DoitStorage manager and connect to the database
func NewStorage(t string, loc string) (*DoitStorage, error) {
	db, err := gorm.Open(t, loc)
	if err != nil {
		return nil, err
	}

	s := &DoitStorage{Conn: db, Type: t, Location: loc}
	s.Conn.DB()
	db.DB().Ping()
	return s, nil
}

//InitSchema Initalize schema
func (s *DoitStorage) InitSchema(overwrite bool) {
	if overwrite {
		s.DeleteSchema()
		s.CreateSchema()
	} else {
		//TODO: Test schema
		//test schema
		tc := s.CheckSchema()
		if !tc {
			s.DeleteSchema()
			s.CreateSchema()
		}
	}

}

//CheckSchema Checks the existing schema
func (s *DoitStorage) CheckSchema() bool {
	var tc bool
	tc = s.Conn.HasTable(&dt.Host{})
	if !tc {
		return tc
	}
	tc = s.Conn.HasTable(&dt.Var{})
	if !tc {
		return tc
	}
	tc = s.Conn.HasTable(&dt.Domain{})
	if !tc {
		return tc
	}
	tc = s.Conn.HasTable(&dt.Group{})
	if !tc {
		return tc
	}
	tc = s.Conn.HasTable(&dt.GroupMatrix{})
	if !tc {
		return tc
	}
	return tc
}

//CreateSchema create the database schema
func (s *DoitStorage) CreateSchema() {
	s.Conn.CreateTable(&dt.Host{})
	s.Conn.CreateTable(&dt.Var{})
	s.Conn.CreateTable(&dt.Domain{})
	s.Conn.CreateTable(&dt.Group{})
	s.Conn.CreateTable(&dt.GroupMatrix{})
}

//DeleteSchema delete the existing schema
func (s *DoitStorage) DeleteSchema() {
	s.Conn.DropTable(&dt.Host{})
	s.Conn.DropTable(&dt.Var{})
	s.Conn.DropTable(&dt.Domain{})
	s.Conn.DropTable(&dt.Group{})
	s.Conn.DropTable(&dt.GroupMatrix{})
}

//Close close the database
func (s *DoitStorage) Close() error {
	err := s.Conn.Close()
	return err
}
