package main

import (
	"net"
	"testing"
	"time"
)

func TestConn(t *testing.T){
	go func() {
		WEB()
	}()
	time.Sleep(time.Second*5)
	_,err:=net.Dial("tcp",":8888")
	if err!=nil {
		t.Errorf(`port 8888 do not have server`)
	}
}