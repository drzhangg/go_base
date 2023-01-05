package data

import "encoding/json"

type TfeInterface interface {
	UpdateStatus(req Request)
}

func NewInterface(kind string,data []byte) TfeInterface {
	switch kind {
	case "DatabaseAccount":
		return NewDatabaseAccount(data)
	case "DatabaseSchema":
		return NewDataBaseSchema(data)
	}

	return nil
}

type Request struct {
	Status string
	Id string
}

type Basic struct {
	DeployStatus string `json:"deployStatus"`
	RunId string `json:"runId"`
}

type DatabaseAccount struct {
	Basic
	Name string
}

func NewDatabaseAccount(data []byte) *DatabaseAccount {
	newData := &DatabaseAccount{}
	json.Unmarshal(data,&newData)
	return newData
}


func (s *DatabaseAccount) UpdateStatus(req Request) {
	s.DeployStatus = req.Status
	s.RunId = req.Id
}

type DataBaseSchema struct {
	Basic
	Address string
}

func NewDataBaseSchema(data []byte) *DataBaseSchema {
	newData := &DataBaseSchema{}
	json.Unmarshal(data,&newData)
	return newData
}

func (s *DataBaseSchema) UpdateStatus(req Request) {
	s.DeployStatus = req.Status
	s.RunId = req.Id
}
