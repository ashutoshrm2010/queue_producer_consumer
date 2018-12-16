package main

import (
	"fmt"
	"time"
	"strconv"
	"testing"
)

type JsonData struct {
	Name string `json:"name"`
}

func main() {
	var done = make(chan bool)
	var Queue =[]JsonData{}
	var jsonData JsonData

	// producer is producing message in every one second
	go func() {
		i:=0 // i I am using for testing reference for name
		for range time.NewTicker(1 * time.Second).C {
			jsonData.Name = "ashu_"+strconv.Itoa(i+1)
			EnQueue(&jsonData, &Queue)
			i=i+1
		}
	}()



	//Consumer is checking queue in every two seconds
	go func() {
		for range time.NewTicker(2 * time.Second).C {
			DeQueue(&Queue)

		}
	}()

	// This function I tried for one more producer for testing

	/*go func() {
		jsonData.Name="123"
		EnQueue(&jsonData,&Queue)
		jsonData.Name="456"
		EnQueue(&jsonData,&Queue)
		jsonData.Name="789"
		EnQueue(&jsonData,&Queue)
	}()*/

	<-done
}

func EnQueue(item *JsonData,Queue*[]JsonData){
	//insert elements to queue
	*Queue=append(*Queue,*item)
	fmt.Println("from produce queue data ",*Queue)

}
func DeQueue(Queue*[]JsonData){
	if len(*Queue) > 1 {
		//remove elements fromm queue
		*Queue = (*Queue)[1:]
		fmt.Println("after consume queue data ",*Queue)
		//*Queue=nil
	} else {
		//Reset queue to empty
		*Queue =make([]JsonData,0)
		fmt.Println("after consume queue data ",*Queue)
		fmt.Println("from consume queue is empty")
	}


}

//Output is:------In FIFO Order
//Here I used two producer concurrently and one consumer
/*
from produce queue data  [{123}]
from produce queue data  [{123} {456}]
from produce queue data  [{123} {456} {789}]
from produce queue data  [{123} {456} {789} {ashu_1}]
from produce queue data  [{456} {789} {ashu_1} {ashu_2}]
after consume queue data  [{456} {789} {ashu_1}]
from produce queue data  [{456} {789} {ashu_1} {ashu_2} {ashu_3}]
from produce queue data  [{456} {789} {ashu_1} {ashu_2} {ashu_3} {ashu_4}]
after consume queue data  [{789} {ashu_1} {ashu_2} {ashu_3} {ashu_4}]
from produce queue data  [{789} {ashu_1} {ashu_2} {ashu_3} {ashu_4} {ashu_5}]
after consume queue data  [{ashu_1} {ashu_2} {ashu_3} {ashu_4} {ashu_5}]
from produce queue data  [{ashu_1} {ashu_2} {ashu_3} {ashu_4} {ashu_5} {ashu_6}]
from produce queue data  [{ashu_1} {ashu_2} {ashu_3} {ashu_4} {ashu_5} {ashu_6} {ashu_7}]
from produce queue data  [{ashu_1} {ashu_2} {ashu_3} {ashu_4} {ashu_5} {ashu_6} {ashu_7} {ashu_8}]
after consume queue data  [{ashu_2} {ashu_3} {ashu_4} {ashu_5} {ashu_6} {ashu_7} {ashu_8}]*/
