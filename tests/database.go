package main

import(
    "fmt"
    "github.com/syndtr/goleveldb/leveldb"
)

func main() {
    db, err := leveldb.OpenFile(db, nil)
    if err != nil {
        fmt.Sprintf("Opening database failed: %s", err))
                  
    }
    defer db.Close()


    data, err := db.Get([]byte("key"), nil)
    err = db.Put([]byte("key"), []byte("value"), nil)




    fmt.Println(data)
    err = db.Delete([]byte("key"), nil)
    defer db.Close()

    }
}
