/* 
Key Value store api really simple

This file sets a server in port 3000 that is waiting 4 types of petitions

read()
add()
delete()
update()

There is two types of structures which are ;
    - collection
    - document

You can read, add, delete a collection and read,update and delete a document 

To get any petition it have to have the secret key

 */



package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"net/http"
)


func add(index string, value string) {
	
	db, err := leveldb.OpenFile("/tmp/foo.db", nil)
	if err != nil {
		log.Fatal("Yikes!")
	}
	defer db.Close()


	err = db.Put([]byte("index"), []byte("value"), nil)
	if err != nil {
		log.Fatal("Yikes!")
	}
}

func server(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/add" {
		http.Error(w, "You should try on POST", http.StatusNotFound)
		return
							
	}
	switch r.Method {
		case "GET":     
		  http.ServeFile(w, r, "Try with Post")
		case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
			if  err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
																					 
		}   
		
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		index := r.FormValue("index")
		key := r.FormValue("key")
		data := r.FormValue("data")
		
		fmt.Println("Put request")
		
		fmt.Println(index,key,data)
		
		// tool_put(index,key,data)



		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
																											 
}
		
}



func main() {
	db, err := leveldb.OpenFile("/tmp/foo.db", nil)
	add("users/p/david/", "dvidd")
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("key: %s | value: %s\n", key, value)
	}
	if err != nil {
		log.Fatal("Yikes!")
	}
}