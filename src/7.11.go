package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	fmt.Println(db)
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update",db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type database map[string]dollars

type dollars float32

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s : %.2f\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter,req *http.Request)  {
	item , price:=req.URL.Query().Get("item"),req.URL.Query().Get("price")
	_,ok:=db[item]
	if !ok{
		if price==""{
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}else {
			temp , err :=strconv.ParseFloat(price,32)
			if err!=nil{
				w.WriteHeader(http.StatusBadRequest) // 400
				fmt.Fprintf(w, "illegal input: %q\n", price)
				return
			}
			db[item]=dollars(temp)
			fmt.Fprintf(w,"add item:%s ,price:%q\n",item,price)
			return
		}
	}else{
		if price==""{
			delete(db,item)
			fmt.Fprintf(w, "deleted item:  %q\n", item)
			return
		}else {
			temp , err :=strconv.ParseFloat(price,32)
			if err!=nil{
				w.WriteHeader(http.StatusBadRequest) // 400
				fmt.Fprintf(w, "illegal input: %q\n", price)
				return
			}
			db[item]=dollars(temp)
			fmt.Fprintf(w,"change %s to %q\n",item,price)
			return
		}
	}
}