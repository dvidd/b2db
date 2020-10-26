

package main


import (
    "fmt"
    "net/http"
    "log"
    "github.com/syndtr/goleveldb/leveldb" 
)


/* 

Key Value store api really simple

This file sets a server in port 3000 that is waiting 3 types of petitions

/get
/put
/delete
 

*/


/* get value from db  */
func get(w http.ResponseWriter, r *http.Request) {
        
    value := r.FormValue("value")
    key := r.FormValue("key")
    
    fmt.Print("%s,%s",key, value)
    db, err := leveldb.OpenFile("db", nil)
    
    if err != nil {
        fmt.Println(err)
    }
    data, err := db.Get([]byte(key),nil)
    fmt.Println(data)
    defer db.Close()
 
}
/* put value into store using a key */
func put(w http.ResponseWriter, r *http.Request) {

    value := r.FormValue("value")
    key := r.FormValue("key")
    
    fmt.Print("%s,%s",key, value)
    db, err := leveldb.OpenFile("db", nil)
    
    if err != nil {
        fmt.Println(err)
    }
    data, err := db.Get([]byte(key),nil)   
    fmt.Println(data)
    defer db.Close()
}
/* deleting value with key */
func delete(w http.ResponseWriter, r *http.Request) {
            
    
    value := r.FormValue("value")
    key := r.FormValue("key")
    
    fmt.Print("%s,%s",key, value)
    db, err := leveldb.OpenFile("db", nil)
    
    if err != nil {
        fmt.Println(err)
    }
    data, err := db.Get([]byte(key),nil)
    fmt.Println(data)
    defer db.Close()

}
func main() {
    
    fmt.Printf("Starting server at port 3000\n")
    
    http.HandleFunc("/put", put)
    http.HandleFunc("/get", get)
    http.HandleFunc("/delete",delete)

    if err := http.ListenAndServe(":3000", nil); err != nil {
            log.Fatal(err)
                                            
    }
                    
                    
}
