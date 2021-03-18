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

func csv_to_struct() []Product{
	csv_file, err := os.Open("sample2.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csv_file.Close()

	r:=csv.NewReader(csv_file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var p Product
	var products []Product
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
		products = append(products, p)
	}
	return products

}

func struct_to_yaml(products []Product) {
	f, err := os.Create("sample.yaml")
	if err != nil {
		panic(err)
	}

	for i:= range products {
		pString := []byte("-" + strconv.Itoa(i) + ":\n" + "street: " + products[i]._1 +
			"\ncity: " + products[i]._2 + "\nzip: " + products[i]._3 + "\nstate: " +
			products[i]._4 + "\nbeds: " + products[i]._5 + "\nbaths: " +
			products[i]._6 + "\nsq__ft: " + products[i]._7 + "\ntype: " +
			products[i]._8 + "\nsale_date: " + products[i]._9 + "\nprice: " +
			products[i]._10 + "\nlatitude: " + products[i]._11 + "\nlongitude: " +
			products[i]._12 + "\n")
		n, err2 := f.Write(pString)
		if err2 != nil {
			panic(err2)
		}
		i++
		fmt.Printf(string(n), "\n")
		//err := ioutil.("sample.yaml", pString, 0644)

	}
}

func main() {
	start := time.Now()
	products := csv_to_struct()
	go struct_to_yaml(products)
	fmt.Printf("Process complete", time.Since(start))
}