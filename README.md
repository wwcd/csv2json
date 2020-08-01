### CSV2JSON

#### Cmd Usage 

```txt
❯ cat input.csv
h1,h2,h3,h4
a,b,c,d
aa,bb,cc,dd
aaa,bbb,ccc,ddd
aaaa,bbbb,cccc,dddd
aaaaa,bbbbb,ccccc,ddddd

❯ ./csv2json input.csv --from-col=1 --to-col=2 --from-row=1 --to-row=2 | jq .
[
  {
    "h2": "b",
    "h3": "c"
  },
  {
    "h2": "bb",
    "h3": "cc"
  }
]
```

---
#### Lib Usage

```go
	input := `h1,h2,h3,h4
a,b,c,d
aa,bb,cc,dd
aaa,bbb,ccc,ddd
aaaa,bbbb,cccc,dddd
aaaaa,bbbbb,ccccc,ddddd`

    output := &bytes.Buffer{}

    err := csv2json.Conv(bytes.NewBufferString(input), output, With(1,2), WithRow(2,3))
    if err != nil {
        panic(err)
    }

    // OUTPUT: [{"h2":"bb","h3":"cc"},{"h2":"bbb","h3":"ccc"}] 
    fmt.Println(output.String())
```
