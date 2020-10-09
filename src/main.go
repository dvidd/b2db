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


func tool_put(index string,key string,j map[string]interface{}) {
    fmt.Println(index)
    fmt.Println(key)
    
    for k, v := range j {
        fmt.Printf("%v %v\n", k, v)
                         
    }

}




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
  //   jsonFile, err := os.Open(dbName)
}

func createfile() {
    err := ioutil.WriteFile(dbName, []byte("Hello"), 0755)
    if err != nil {
        fmt.Printf("Unable to write file: %v", err)
    }
}



func put(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/put" {
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

    var b = []byte(`{"a":"b", "c":1, "d": ["e", "f"]}`)
    var j map[string]interface{}
        err := json.Unmarshal(b, &j)
            if (err != nil) {
        return
                                 
    }


    tool_put("index","key",j)

    


    //checkdb()

    //http.HandleFunc("/put", put)
    //http.HandleFunc("/delete" , delete)
    //http.HandleFunc("/read" , read)
    //http.HandleFunc("/update", update)
    
    
    fmt.Printf("Starting server at port 8080\n")
            if err := http.ListenAndServe(":8080", nil); err != nil {
            log.Fatal(err)
                                            
    }
                    
                    
}
