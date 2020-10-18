package tools

import (
    "fmt"
    "os"
    "io/ioutil"
)


dbName = "db.sqlite"

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



func read() (index int) {
	
    fmt.Println(index)
    for i := 0; i < 10; i++ {
        fmt.Println(index)
    }
    return

}

func add() (index int) {
    return
}

func delete() (index int) {
    return
}

