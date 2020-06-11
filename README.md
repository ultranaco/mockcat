# mockcat
Just another mock server but this one is chido. Mockcat is a command line application to raise a http server from any file type also allow make CRUD operations with JSON file that contains an array of elements or one element at root

## Install:

```
$ go install github.com/ultranaco/mockcat
```

## Usage:

Generates a simple http server to make CRUD requests http://localhost:8080/collection/item/ or http://localhost:8080/collection/item/{id}

`{id}` is default root property in a JSON file that contains an array, default property can be changed with the optional parameter property-matcher `-m url`

```
$ mockcat collection-item.json
```

## Options

```
-m   :property matcher used to retrieve an item from array of elements through root properties, default 'id'
-p   :port to listen incoming requests, default '8080'
```



