package utils

import "sync"
import "errors"
import "fmt"

type MapWithRWMutex struct {
	mapInst map[interface{}]interface{}
	sync.RWMutex
}

func NewMapWithRWMutex() *MapWithRWMutex {
	return &MapWithRWMutex{
		mapInst: map[interface{}]interface{}{},
	}
}

func (p *MapWithRWMutex) Add(key interface{}, value interface{}) (err error) {
	p.Lock()
	defer p.Unlock()
	_, ok := p.mapInst[key]
	if ok {
		err = errors.New(fmt.Sprintf("Allready has :", key))
	} else {
		p.mapInst[key] = value
	}
	return
}

func (p *MapWithRWMutex) Get(key interface{}) (value interface{}, ok bool) {
	p.RLock()
	defer p.RUnlock()
	value, ok = p.mapInst[key]
	return
}

func (p *MapWithRWMutex) PopOne() (value interface{}, ok bool) {
	p.Lock()
	defer p.Unlock()
	if len(p.mapInst) > 0 {
		for k, v := range p.mapInst {
			value = v
			delete(p.mapInst, k)
			break
		}
		ok = true
	} else {
		ok = false
	}
	return
}

func (p *MapWithRWMutex) Remove(key interface{}) (err error) {
	// fmt.Println("call Remove begin")
	p.Lock()
	// fmt.Println("call Remove middle1")
	defer p.Unlock()
	// fmt.Println("call Remove middle2")
	_, ok := p.mapInst[key] //此时不能调用p.Has函数，否则就锁重入了，会卡住
	if ok {
		// fmt.Println("call Remove middle3")
		delete(p.mapInst, key)
		// fmt.Println("call Remove middle3.1")
	} else {
		// fmt.Println("call Remove middle4")
		err = errors.New(fmt.Sprint("Do not have:", key))
	}
	//fmt.Println("call Remove end")
	return
}

func (p *MapWithRWMutex) Has(key interface{}) bool {
	p.RLock()
	defer p.RUnlock()
	_, ok := p.mapInst[key]
	return ok
}

func (p *MapWithRWMutex) Len() int {
	return int(len(p.mapInst))
}

func (p *MapWithRWMutex) Clear() {
	p.Lock()
	defer p.Unlock()
	p.mapInst = map[interface{}]interface{}{}
}

func (p *MapWithRWMutex) IsEmpty() bool {
	if p.Len() == 0 {
		return true
	}
	return false
}

func (p *MapWithRWMutex) Items() (keys []interface{}, values []interface{}) {
	p.RLock()
	defer p.RUnlock()
	for key, value := range p.mapInst {
		keys = append(keys, key)
		values = append(values, value)
	}
	return
}

func (p *MapWithRWMutex) Keys() (keys []interface{}) {
	keys, _ = p.Items()
	return
}

func (p *MapWithRWMutex) Values() (values []interface{}) {
	_, values = p.Items()
	return
}
