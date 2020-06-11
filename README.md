# mockcat
Just another mock server but this one is chido. Mockcat is a command line application to raise a http server from any file type also allow make CRUD operations with JSON file that contains an array of elements or one element at root

## Install:

```
$ go install github.com/ultranaco/mockcat
```

## Usage:

With de given file `collection-item.json`

```
[{
  "id": 1,
  "url": "foo"
},
{
  "id": 2,
  "url": "bar"
}]
```

Generates a simple http server to make CRUD requests http://localhost:8080/collection/item/ or http://localhost:8080/collection/item/{id}

`{id}` is default root property in a JSON file that contains an array, default property can be changed with the optional parameter property-matcher `-m url`

```
$ mockcat collection-item.json
```

## Options:

```
-m   :property-matcher used to retrieve an item from an array of elements through some root property, default id
-p   :port to listen incoming requests, default 8080
```

## Example:

The following command listen incoming request at port `8082` and property-matcher is changed to seek items by `url`, `http://localhost:8082/collection/item/{url}` or retrive all collection `http://localhost:8082/collection/item/`

```
$ mockcat collection-item.json -p 8082 -m url
```

