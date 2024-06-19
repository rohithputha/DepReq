package DepReq

import (
	"fmt"
)

func (d *DepR) putInstance(req *request) error {
	d.instanceMap[req.ii.instanceName] = req.ii
	fmt.Println("Instance added to map")
	d.responseChan <- &response{ii: req.ii, err: nil}
	return nil
}

func (d *DepR) getInstance(req *request) {
	if _, ok := d.instanceMap[req.ii.instanceName]; !ok {
		d.responseChan <- &response{ii: req.ii, err: fmt.Errorf("instance not found")}
		return
	}
	d.responseChan <- &response{ii: d.instanceMap[req.ii.instanceName], err: nil}

}

func (d *DepR) requestRouter() {
	for {
		req:= <-d.requestChan
		if req.requestType == "put" {
			d.putInstance(req)
		} else if req.requestType == "get" {
			d.getInstance(req)
		}
	}
}