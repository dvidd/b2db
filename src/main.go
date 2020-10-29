

package main


import (
    "fmt"
    "net/http"
    "log"
    "golang.org/x/tools/godoc/util"
    "github.com/syndtr/goleveldb/leveldb" 
)


/* 

Key Value store api really simple

This file sets a server in port 3000 that is waiting 3 types of petitions

/get
/put
/delete
/subset : for iterate on element in the same collection subset

*/


/* get value from db  */
func get(w http.ResponseWriter, r *http.Request) {
        
    key := r.FormValue("key")
    
    fmt.Print(key)
    db, err := leveldb.OpenFile("db", nil)
    
    if err != nil {
        fmt.Println(err)
    }
    data ,err2 := db.Get([]byte(key),nil)
    fmt.Println(data)
    if err2 != nil {
        fmt.Println(err2)
    }
    w.Write(data)
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
    err = db.Put([]byte(key),[]byte(value), nil)   
    
    fmt.Println(err)
    defer db.Close()
}
/* deleting value with key */

func delete(w http.ResponseWriter, r *http.Request) {
            
    key := r.FormValue("key")
    
    fmt.Print(key)
    db, err := leveldb.OpenFile("db", nil)
    
    if err != nil {
        fmt.Println(err)
    }
    err = db.Delete([]byte(key),nil)
    fmt.Println(err)

    defer db.Close()

}
/* Iterate over collection subst of elements  */
func subset(w http.ResponseWriter, r *http.Request) {
            
    key := r.FormValue("key")
    
    fmt.Print(key)
    db, err := leveldb.OpenFile("db", nil)
    
    if err != nil {
        fmt.Println(err)
    }
    ter := db.NewIterator(util.BytesPrefix([]byte(key)), nil)
    for iter.Next() {
            fmt.Println(iter)
    }
    iter.Release()
    err = iter.Error()
    defer db.Close()

}

func main() {
    
    fmt.Printf("Starting server at port 3000\n")
    
    http.HandleFunc("/put", put)
    http.HandleFunc("/get", get)
    http.HandleFunc("/delete",delete)
    http.HandleFunt("/subset", subset)
    if err := http.ListenAndServe(":3000", nil); err != nil {
            log.Fatal(err)
                                            
    }
                    
                    
}
