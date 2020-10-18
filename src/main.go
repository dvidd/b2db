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
    "net/http"
    "log"
    "github.com/syndtr/goleveldb/leveldb" 
)

// func _auth(key string) {
    
// }

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
            
            fmt.Println(index,key,data)
            
            // tool_put(index,key,data)



            default:
                fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
                                                                                                                 
    }
            
}



func main() {
    db, err := leveldb.OpenFile("db.db", nil)
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()
    fmt.Printf("Starting server at port 8080\n")
            if err := http.ListenAndServe(":8080", nil); err != nil {
            log.Fatal(err)
                                            
    }
                    
                    
}
