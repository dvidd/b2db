 keyvalue
Distributed key value store API made simple using go and leveldb, made really simple

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

#### /put Method = [POST]
Read a document or an entire collection

```javascript
params {
    key: "[key]",
    value: "[value]",
}
``````

#### /get Method = [POST]
For adding a document or a collection or both at the same time

```javascript
params {
    value: "[value]",
}
```
#### /delete Method = [POST]
For deleting a collection or a document ( Use carefuly )

```javascript
params {
    value: "[value]"
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

