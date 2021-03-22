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
	titles []string
	columns []string
}

func csv_to_struct(products chan Product) {
	//defer wg.Done()
	csv_file, err := os.Open("FL_insurance_sample.csv")
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
	in := 0
	for _, rec := range records {
		if in == 0 {
			for i := range rec {
				p.titles = append(p.titles, rec[i])
			}
		} else {
			for i := range rec {
				p.columns = append(p.columns, rec[i])
			}
		}
		in++
		products <- p
		p.columns = nil
	}
	close(products)

}

func struct_to_yaml(products chan Product, finit chan bool) {

	f, err := os.Create("sample.yaml")
	if err != nil {
		close(products)
	}
	i:= 0
	for products:= range products {
		var pString string
		if i > 0 {
			pString = "-" + strconv.Itoa(i) + ":\n"
		}
		x := 0
		for _, in := range products.columns {
			pString = pString + products.titles[x]+": " + in+"\n"
			x++
		}
		pByte := []byte(pString)
		n, err2 := f.Write(pByte)
		if err2 != nil {
			panic(err2)
		}
		fmt.Printf(string(n), "\n")
		//err := ioutil.("sample.yaml", pString, 0644)
		i++
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