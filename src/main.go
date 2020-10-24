

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
func get(key string) {
        
    db, err := leveldb.OpenFile("db", nil)
    
    if err != nil {
        fmt.Println(err)
    }
    data, err := db.Get([]byte(key),nil)
    fmt.Println(data)
    defer db.Close()
 
}
/* put value into store using a key */
func put(key string,value string) {

    db, err := leveldb.OpenFile("db", nil)
    
    if err != nil {
        fmt.Println(err)
    }
    data, err := db.Get([]byte(key),nil)   
    fmt.Println(data)
    defer db.Close()
}
/* deleting value with key */
func delete(key string) {
    db, err := leveldb.OpenFile("db", nil)
    
    if err != nil {
        fmt.Println(err)
    }
    data, err := db.Get([]byte(key),nil)
    fmt.Println(data)
    defer db.Close()

}

func add(w http.ResponseWriter, r *http.Request) {
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
            

            fmt.Fprintf(w, "Index = %s\n", index)
            fmt.Fprintf(w, " Key  = %s\n", key)
            fmt.Fprintf(w, "Data = %s\n", data)
            // tool_put(index,key,data)



            default:
                fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
                                                                                                                 
    }
    
}

func main() {

    
    fmt.Printf("Starting server at port 3000\n")
    
    
    http.HandleFunc("/add", add)

    if err := http.ListenAndServe(":3000", nil); err != nil {
            log.Fatal(err)
                                            
    }
                    
                    
}
