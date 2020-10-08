package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
	fmt.Println("Hello, 世界");
    checkdb();
    for i:= 0; i < 10; i++ {
        fmt.Println(i)
    }
    return
}

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
}

func createfile() {
    err := ioutil.WriteFile(dbName, []byte("Hello"), 0755)
    if err != nil {
        fmt.Printf("Unable to write file: %v", err)
    }
}

func add(doc string) {
    return
}


func read(doc string) {
    return 
}
