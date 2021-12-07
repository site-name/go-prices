Money manipulation with golang

All the ideas borrowed from:
https://github.com/mirumee/prices

thank to https://github.com/Rhymond/go-money, this library was done properly.

**Example**

```go  
  import (
    money "github.com/site-name/go-prices"
  )

  money1, _ := money.NewMoney(34.56, "USD")
  money2, _ := money.NewMoney(23, "usd")
  
  sum, _ := money1.Add(money2)
  
  fmt.Println(sum.String())
```
