# Simple Key Store (Go)

A simple key-value (SKS) store exposing a HTTP API. It was also implemented in Scala [here](https://github.com/lucastorri/scala-sks).


## Setting the Project

I'm quite inexperienced with Go, so a better way of doing it is probably out there. For now, that's what we have.

Create a workspace directory, and clone this project using:

```
export GOPATH=/path/to/workspace

mkdir -p $GOPATH

cd $GOPATH

git clone https://github.com/lucastorri/golang-sks.git src/github.com/lucastorri/sks

cd src/github.com/lucastorri/sks

gpm
```


## Compile & Run

```
cd $GOPATH

go install github.com/lucastorri/sks

./bin/sks
```

### Options

The following flags can be used:
  
  * `-port=<number>`: set the HTTP port to be used
  * `-store=<details>`: how files will be store, with two options:
    * `men`: store files in memory;
    * `dir:<path>`: store in the given path and access them through memory mapped files.


## API

* `GET /{key}`: return 200 and body with value for that key, or 404 if key not defined;
* `POST /key`: save body content to key.


## Testing

Load some files into the ska, generate a list with URLs and use it as an input to [siege](http://www.joedog.org/siege-home/).

```
for i in {1..1000}; do 
    url=http://localhost:12121/$i
    curl -X POST --data "@/path/to/some/file.txt" $url
    echo $url >> urls.txt
done

siege -f urls.txt -i -t 10S -d 0 -c 10 -b
```

Results (more or less around this):

```
Transactions:		       16377 hits
Availability:		      100.00 %
Elapsed time:		        9.44 secs
Data transferred:	       29.82 MB
Response time:		        0.00 secs
Transaction rate:	     1734.85 trans/sec
Throughput:		        3.16 MB/sec
Concurrency:		        3.79
Successful transactions:       16377
Failed transactions:	           0
Longest transaction:	        0.01
Shortest transaction:	        0.00
```
