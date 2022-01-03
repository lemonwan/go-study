package ch3

type Request interface{}

type Response interface{}

type Filter interface {
	Process(data Request) (Request, error)
}
