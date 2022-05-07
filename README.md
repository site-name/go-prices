Money manipulation with golang

All the ideas borrowed from:
https://github.com/mirumee/prices

thank to https://github.com/Rhymond/go-money, this library was done properly.

**Example**

```go
  package main
  
  import (
    money "github.com/site-name/go-prices"
    "log"
  )

  money1, err := money.NewMoney(34.56, "USD")
  if err != nil {
    log.Fatalln(err)
  }
  money2, err := money.NewMoney(23, "usd")
  if err != nil {
    log.Fatalln(err)
  }
  
  sum, err := money1.Add(money2)
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Println(sum.String())
```
