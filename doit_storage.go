package main

import (
	dt "github.com/DevOpsInvTech/doittypes"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type DoitStorage struct {
	Type     string
	Location string
	Conn     gorm.DB
}

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

func (s *DoitStorage) InitSchema(overwrite bool) {
	if overwrite {
		s.Conn.CreateTable(&dt.Host{})
		s.Conn.CreateTable(&dt.HostVar{})
		s.Conn.CreateTable(&dt.Var{})
		s.Conn.CreateTable(&dt.Domain{})
		s.Conn.CreateTable(&dt.Group{})
		s.Conn.CreateTable(&dt.GroupVar{})
		s.Conn.CreateTable(&dt.GroupMatrix{})
	} else {
		//TODO: Test schema
		//test schema
	}

}

func (s *DoitStorage) Close() error {
	err := s.Conn.Close()
	return err
}
