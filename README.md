# HN_Logs

The goal of this project is to provide two endpoints, the first one :
`/count/{datetime}` will return the amounts of requests done on the given datetime.
The second :
`/top/{date}/{size}` will return the the Nth top requests done on the given datime, 
N is specified with the size param.

All the data is index with the file `hn_logs.tsv` and are indexed in Tree.

# Example :

the request : `http://localhost:8080/top/2015/3` 
return :
```json
[
    {
        "Query": "http%3A%2F%2Fwww.getsidekick.com%2Fblog%2Fbody-language-advice",
        "Count": 6675
    },
    {
        "Query": "http%3A%2F%2Fwebboard.yenta4.com%2Ftopic%2F568045",
        "Count": 4652
    },
    {
        "Query": "http%3A%2F%2Fwebboard.yenta4.com%2Ftopic%2F379035%3Fsort%3D1",
        "Count": 3100
    }
]

```

the request :  `http://localhost:8080/count/2015-08-01`
return :
```json
{
    "Count": 432605
}
```

#Install

Go version 1.11 min.
`export GO111MODULE=on`
`go run main.go`
or
`go run server`