package models

type CreateRequest struct {
	DBName     string
	Username   string
	Password   string
	Port       int32
	Replicas   int32
	Capacity   string
	AccessMode string
}

type CreateResponse struct {
	Id string
}

type DeleteRequest struct {
	Id string
}

type UpdateRequest struct {
	Id       string
	Replicas int32
}
