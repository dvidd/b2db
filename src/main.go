package main


import (
    "fmt"
    "log"
    "net/http"
    "os"
    "io/ioutil"
    "encoding/json"

)

// Check if index collection 

const dbName = "db.json"




func checkdb() {
    if _, err := os.Stat(dbName); err == nil {
        // path/to/whatever exists
        fmt.Println("\n Database exists");

   }else if os.IsNotExist(err) {
        // path/to/whatever does *not* exist
        fmt.Println("\n Database does not exists");
        fmt.Println("\n Creating file...");
        createfile();

    } else {
        // Schrodinger: file may or may not exist. See err for details.
        fmt.Println("\n Something went wrong while creatring the database");
        // Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence

}
    jsonFile, err := os.Open(dbName)
}

func createfile() {
    err := ioutil.WriteFile(dbName, []byte("Hello"), 0755)
    if err != nil {
        fmt.Printf("Unable to write file: %v", err)
    }
}



func add(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/add" {
            http.Error(w, "404 not found.", http.StatusNotFound)
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
            name := r.FormValue("name")
            address := r.FormValue("address")
            fmt.Fprintf(w, "Name = %s\n", name)
            fmt.Fprintf(w, "Address = %s\n", address)
    
            default:
            fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
                                                                                                                 
    }
            
}





func main() {

    checkdb()

    http.HandleFunc("/add", add)
    //http.HandleFunc("/delete" , delete)
    //http.HandleFunc("/read" , read)
    //http.HandleFunc("/update", update)
    
    
    fmt.Printf("Starting server at port 8080\n")
            if err := http.ListenAndServe(":8080", nil); err != nil {
            log.Fatal(err)
                                            
    }
                    
                    
}
