# go-markdown-server

path base markdown view server

## Description

go-markdown-server is path base markdown view server.

* HTTP GET  - markdown viewer
* HTTP POST - markdown upload



### update markdown


````
$ curl -F "file=@README.md" -v http://[address:port]/my-markdown-path
````

### view markdown

open browser

````
http://[address:port]/my-markdown-path
````

## examples

POST and Web access

````
$ curl -F "file=@README.md" -v http://127.0.0.1:5000/example/readme
````







