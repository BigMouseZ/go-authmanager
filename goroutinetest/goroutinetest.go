package main

import (
	"fmt"
	`go-authmanager/goutils`
	"io/ioutil"
)

var complete chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	complete <- 1000 // 执行完毕了，发个消息
}

var ch chan int = make(chan int)

func foo(id int) { //id: 这个routine的标号
	ch <- id
}
func main() {
	/*	go loop() // 启动一个goroutine
		go loop()

		fmt.Println(<- complete) // 直到线程跑完, 取到消息. main在此阻塞住
		<- complete

		// 开启5个routine
		for i := 0; i < 5; i++ {
			go foo(i)
		}

		// 取出信道中的数据
		for i := 0; i < 5; i++ {
			fmt.Print(<- ch)
		}*/

	/*ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3*/

	/*fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
	fmt.Println(<-ch) // 3*/
	// 显式地关闭信道
	/*	close(ch)
		for v := range ch {
			fmt.Println(v)
		}*/

	url := "http://172.16.2.20:8100/expressway_track/shortpathtest/ShortPathTest.htm"
	abody:=make(map[string]string,0)
	abody["startGantryCode"] = "G008496949511004236"
	abody["endGantryCode"] = "G008162237386915590"
	resp,err:=goutils.PostUtil(url,abody,nil,nil )
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(body))

		/*ch := make(chan int, 3)
		ch <- 1
		ch <- 1
		ch <- 1
		ch <- 1 //这一行操作就会发生阻塞，因为前三行的放入数据的操作已经把channel填满了
*/
}

