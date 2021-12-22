package ch3

import (
	"errors"
	"time"
)

type ReusableObj struct {
}

type ObjPool struct {
	bufChan chan *ReusableObj
}

func NewObjPool(num int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, num)
	for i := 0; i < num; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

func (pool *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-pool.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (pool *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case pool.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
