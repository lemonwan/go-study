package test

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	now := time.Now()
	t.Log(now)
	t.Log(runtime.NumCPU())
	t.Log(time.Since(now))

	st, err := time.ParseDuration("-24h")
	add := time.Now().Add(st)
	t.Log(add, err)
	date, month, day := add.Date()
	t.Log(date, month.String(), day)
}

func TestChan(t *testing.T) {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go send(ch)
	go receive(ch, &wg)
	wg.Wait()
}

func send(c chan<- int) {
	c <- 10
}

func receive(c <-chan int, w *sync.WaitGroup) {
	defer w.Done()
	t := <-c
	fmt.Println(t)
}

func TestLock(t *testing.T) {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	t.Log("Locking (G0)")
	mutex.Lock()
	t.Log("Locked (G0)")
	wg.Add(3)
	for i := 1; i < 4; i++ {
		go func(i int) {
			t.Logf("Locking (G%d)\n", i)
			mutex.Lock()
			t.Logf("Locked (G%d)\n", i)

			time.Sleep(time.Second * 2)
			mutex.Unlock()
			t.Logf("unlocked (G%d)\n", i)
			wg.Done()
		}(i)
	}
	time.Sleep(time.Second * 5)
	t.Log("ready unlock (G0)")
	mutex.Unlock()
	t.Log("unlocked (G0)")
	wg.Wait()
}

// sync.Mutex用作结构体的一部分，这样这个struct就可以防止多线程修改数据
type Book struct {
	Name string
	L    *sync.Mutex
}

func (b *Book) SetName(wg *sync.WaitGroup, name string) {
	defer func() {
		fmt.Printf("Unlock set name: %v\n", name)
		b.L.Unlock()
		wg.Done()
	}()

	b.L.Lock()
	fmt.Printf("Lock set name: %v\n", name)
	time.Sleep(time.Second * 2)
	b.Name = name
}

func TestMutexStruct(t *testing.T) {
	b := new(Book)
	b.L = new(sync.Mutex)
	wg := &sync.WaitGroup{}
	books := []string{"《三国演义》", "《道德经》", "《西游记》"}
	for _, book := range books {
		wg.Add(1)
		go b.SetName(wg, book)
	}
	wg.Wait()
}

// 读写锁使用
func TestRWMutex(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(20)
	var rwMutex sync.RWMutex
	Data := 0
	for i := 0; i < 10; i++ {
		// 读
		go func() {
			rwMutex.RLock()
			defer rwMutex.RUnlock()
			t.Logf("Read Data: %v\n", Data)
			wg.Done()
			time.Sleep(time.Second * 2)
		}()
		// 写
		go func(n int) {
			rwMutex.Lock()
			defer rwMutex.Unlock()
			Data += n
			t.Logf("Write Data: %v %d\n", Data, n)
			wg.Done()
			time.Sleep(time.Second * 2)
		}(i)
	}
	time.Sleep(time.Second * 5)
	wg.Wait()
}

func TestSyncMap(t *testing.T) {
	m := make(map[string]string)
	m["a"] = "b"
	t.Log(m)

	var m1 sync.Map
	m1.Store("name", "lemon")
	m1.Store("gender", "male")

	v, ok := m1.LoadOrStore("name", "hello")
	t.Log(ok, v)

	v, ok = m1.LoadOrStore("name1", "hello")
	t.Log(ok, v)

	v, ok = m1.Load("name")
	if ok {
		t.Log("key 存在，值是：", v)
	} else {
		t.Log("key 值不存在")
	}

	f := func(k, v interface{}) bool {
		t.Log(k, v)
		return true
	}
	m1.Range(f)

	m1.Delete("name1")
	t.Log(m1.Load("name1"))

}

func TestReflect(t *testing.T) {
	var a int = 5
	v := reflect.ValueOf(a)
	w := reflect.TypeOf(a)
	t.Log(v, w, v.Type(), v.Kind(), w.Kind())

	var b [5]int = [5]int{1, 2, 3, 4}
	t.Log(b)
	t.Log(reflect.TypeOf(b), reflect.ValueOf(b))
	t.Log(reflect.TypeOf(b).Kind(), reflect.TypeOf(b).Elem())
	t.Log(reflect.ValueOf(b).Kind(), reflect.ValueOf(b).Type())
}

func TestReflectModify(t *testing.T) {
	a := 10
	t.Log(reflect.TypeOf(a), reflect.ValueOf(a))
	t.Log(reflect.TypeOf(a).Kind(), reflect.TypeOf(&a).Elem())
	t.Log(reflect.ValueOf(a).Kind(), reflect.ValueOf(&a).Elem())

	elem := reflect.ValueOf(&a).Elem()
	t.Log(elem.CanAddr())
	elem.SetInt(10000)
	t.Log(a)

	var b [5]int = [5]int{5, 6, 7, 8}
	t.Log(reflect.TypeOf(b), reflect.TypeOf(b).Kind(), reflect.TypeOf(b).Elem())
}
