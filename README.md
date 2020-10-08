# b2db
Database backend fusion 
## How it works :

The aplication sets a server working on : 
```localhost:3000 ```

## Functions :

#### put( ) Method = [POST]
For adding a document or a collection or both at the same time
```
localhost:3000/put/
```
Params to send {

    index: "/main/app/..",
    key: "[YOUR_KEY]",
    data: "data"

    }

#### delete( ) Method = [POST]
For deleting a collection or a document ( Use carefuly )
```
localhost:3000/delete/
```
Params to send {

    index: "/main/app/..",
    key: "[YOUR_KEY]",

    }


#### read( ) Method = [POST]
Read a document or an enterie collection
```
localhost:3000/read/
```
{
Params to send {

    index: "/main/app/..",
       key: "[YOUR_KEY]",


   }
}

#### update( ) Method = [POST]
Update a document
```
localhost:3000/put/
```
Params to send {

    index: "/main/app/document",
    key: "privatekeyoftheserver",
    data: "data"

    }



