package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	//"io"
	//"bufio"
)

//var wg sync.WaitGroup

type Product struct {
	_1 string //"id"
	_2 string //"Description"
	_3 string //"Name"
	_4 string //"number1"
	_5 string //"number2"
	_6 string //"number3"
	_7 string //"number4"
	_8 string
	_9 string
	_10 string
	_11 string
	_12 string
}

func csv_to_struct(products chan Product) {
	//defer wg.Done()
	csv_file, err := os.Open("sample2.csv")
	if err != nil {
		fmt.Println(err)
		close(products)
		return
	}
	defer csv_file.Close()

	r:=csv.NewReader(csv_file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		close(products)
		os.Exit(1)
	}

	var p Product

	for _, rec := range records {
		p._1 = rec[0]
		p._2 = rec[1]
		p._3 = rec[2]
		p._4 = rec[3]
		p._5 = rec[4]
		p._6 = rec[5]
		p._7 = rec[6]
		p._8 = rec[7]
		p._9 = rec[8]
		p._10 = rec[9]
		p._11 = rec[10]
		p._12 = rec[11]
		products <- p
	}
	close(products)

}

func struct_to_yaml(products chan Product, finit chan bool) {

	f, err := os.Create("sample.yaml")
	if err != nil {
		close(products)
	}
	i := 0
	for products:= range products {
		pString := []byte("-" + strconv.Itoa(i) + ":\n" + "street: " + products._1 +
			"\ncity: " + products._2 + "\nzip: " + products._3 + "\nstate: " +
			products._4 + "\nbeds: " + products._5 + "\nbaths: " +
			products._6 + "\nsq__ft: " + products._7 + "\ntype: " +
			products._8 + "\nsale_date: " + products._9 + "\nprice: " +
			products._10 + "\nlatitude: " + products._11 + "\nlongitude: " +
			products._12 + "\n")
		n, err2 := f.Write(pString)
		if err2 != nil {
			panic(err2)
		}
		i++
		fmt.Printf(string(n), "\n")
		//err := ioutil.("sample.yaml", pString, 0644)

	}
	finit<-true
}

func main() {
	start := time.Now()
	products := make(chan Product)
	finit := make(chan bool)
	//wg.Add(1)
	go csv_to_struct(products)
	go struct_to_yaml(products, finit)
	//go csv_to_struct(products)
	//go struct_to_yaml(products, finit)
	<-finit
	//wg.Wait()
	fmt.Printf("Process complete: ", time.Since(start))
}