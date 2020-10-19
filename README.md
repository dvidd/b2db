# keyvalue
Distributed key value store made simple using go and leveldb, made really simple

## Get started
```
$ git clone https://github.com/dvidd/keyvalue/
```
```
$ cd keyvalue
```
```
$ docker run 
```


## How it works :
The aplication sets a server working on : 
```localhost:3000 ```

## Functions :

#### /put/ Method = [POST]
For adding a document or a collection or both at the same time

```json
data {

    index: "/main/app/..",
    key: "[YOUR_KEY]",
    data: "data"

    }
```
#### /delete Method = [POST]
For deleting a collection or a document ( Use carefuly )
```javascript
data {

    index: "/main/app/..",
    key: "[YOUR_KEY]",

    }
```

#### /read Method = [POST]
Read a document or an entire collection
```javascript
data {
    
    index: "/main/app/..",
    key: "[YOUR_KEY]",


   }
}
```
#### /update Method = [POST]
Update a document
```javascript
data{

    index: "/main/app/document",
    key: "privatekeyoftheserver",
    data: "data"

    }


```


MIT License

Copyright (c) 2020 

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

