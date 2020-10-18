package main


import (
    "fmt"
    // "log"
    "net/http"
    "os"
    "io/ioutil"
    "encoding/json"
)

// Check if index collection 
const dbName = "db.json"


func checkdb() {
    if _, err := os.Stat(dbName); err == nil {
        fmt.Println("\n Database exists");
   }else if os.IsNotExist(err) {
        fmt.Println("\n Database does not exists");
        fmt.Println("\n Creating file...");
        createfile();
    } else {
        fmt.Println("\n Something went wrong while creatring the database");
}

}

func createfile() {
    err := ioutil.WriteFile(dbName, []byte("Hello"), 0755)
    if err != nil {
        fmt.Printf("Unable to write file: %v", err)
    }
}

func write() {
    plan, _ := ioutil.ReadFile(dbName)
    var data interface{}
    err := json.Unmarshal(plan, &data)
    fmt.Println(json.Marshal(err))

    
}

func read() {
    data, err := ioutil.ReadFile(dbName)
    data = string(data)
    var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(data), &result)

	//Reading each value by its key
	fmt.Println("Id :", result["id"],
		"\nName :", result["name"],
		"\nDepartment :", result["department"],
        "\nDesignation :", result["designation"])
        
    if err != nil {
        fmt.Println("File reading error", err)
        return
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

    read()


    
    // fmt.Printf("Starting server at port 8080\n")
    //         if err := http.ListenAndServe(":8080", nil); err != nil {
    //         log.Fatal(err)
                                            
    // }
                    
                    
}
