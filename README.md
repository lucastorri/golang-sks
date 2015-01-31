# Simple Key Store

Memory-mapped key-value store through a HTTP API.


## Compile & Run

```
cd $workspace
gpm
go install github.com/lucastorri/sks && ./bin/sks
```


## API

* Get /key: return 200 and body with value for that key, or 404 if key not defined;
* Post /key + body: save body content to key.
