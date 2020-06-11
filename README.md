# Nicepay-go 

Unofficial Nicepay library for golang

### Installation

> go get -u github.com/yisiper/nicepay-go



### Quick Start

```go
client := nicepay.NewClient("iMid-Key", "Merchant-Key", "CallbackUrl")
client.Env = nicepay.Development // nicepay.Production

gateway := nicepay.NewCoreGateway(client)
resp, err := gateway.Registration(&nicepay.RegistrationRequest{
    // fill the struct detail
})
```



### License

See [LICENSE](LICENSE)