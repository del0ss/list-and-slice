package main

//var Pepe = make(chan int)

func main() {
	var store model.Storage
	store = list.NewList()
	store.Add(88)
	zxc := store.Get(0)
	fmt.Println(xzc)
	store = &slice.Slice{}
	store.Add(22)

// 	defer sleeper()
// 		var mutex sync.Mutex
// 		for i := 1; i < 4; i++ {
// 			go sleeper(i, Pepe, &mutex)
// 		}
// 		<-Pepe
// 		fmt.Println("Функция main отработала")
	
}

//func sleeper(number int, ch chan int, mutex *sync.Mutex) {
//	mutex.Lock()
//	time.Sleep(3 * time.Second)
//	u := number
//	fmt.Println("Функция sleeper отработала")
//	mutex.Unlock()
//	Pepe <- u
//}
