# Simple Key Store

In memory key-value store through a HTTP API.


## Compile & Run

```
cd $workspace
go build sks && ./sks
```


## API

* Get /key: return 200 and body with value for that key, or 404 if key not defined;
* Post /key + body: save body content to key.


## TODO

* Store to disk: create file named after the key hash, put contents there.