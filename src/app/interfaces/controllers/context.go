package controllers

type Context interface {
	Bind(interface{}) error
	JSON(int, interface{})
	Status(int)
	Param(string) string
}
