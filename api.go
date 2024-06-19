package main

import ("reflect" 
		"sync")

var depReqApi *DepReqApi
var once sync.Once

type DepReqApi struct {
	depR *DepR
}



func GetDepReqApi() *DepReqApi{
	once.Do(func(){
		depReqApi = &DepReqApi{depR: &DepR{
			instanceMap: make(map[string] *instanceInfo),
			requestChan: make(chan *request),
			responseChan: make(chan *response),
		}}
		go depReqApi.depR.requestRouter()
	})
	return depReqApi
}

func (d *DepReqApi) Put(instName string, inst interface{}) error{
	request := &request{ii: &instanceInfo{instanceName: instName, instance: reflect.ValueOf(inst), instanceType: reflect.TypeOf(inst)}, requestType: "put"}
	d.depR.requestChan <- request
	response := <-d.depR.responseChan
	if response.err != nil{
		return response.err
	}
	return nil
}

func (d *DepReqApi) Get(instName string) (interface{}, error){
	request := &request{ii: &instanceInfo{instanceName: instName}, requestType: "get"}
	d.depR.requestChan <- request
	response := <-d.depR.responseChan
	if response.err != nil{
		return nil, response.err
	}
	return response.ii.instance.Interface(), nil
}
