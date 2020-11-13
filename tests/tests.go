package test

import (
    "fmt"
    "ioutil"
    "net/http"
)

func main() {
    fmt.Println("\nTesting putting a value : ")
    resp, err := http.Get("http://localhost:3000/put?key=testing?value=testingvalue")
    if err != nil {
	    // handle error
        fmt.Println("Done")
    } else {
        fmt.Println("\n Error  Didn't success adding value to the db \n")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

}
