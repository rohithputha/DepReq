package DepReq

import "reflect"


type DepR struct {
	instanceMap map[string] *instanceInfo
	requestChan chan *request
	responseChan chan *response
}
type request struct{
	ii *instanceInfo
	requestType string
}
type response struct{
	ii *instanceInfo
	err error
}

type instanceInfo struct {
	instanceName string
	instanceType reflect.Type
	instance reflect.Value
}