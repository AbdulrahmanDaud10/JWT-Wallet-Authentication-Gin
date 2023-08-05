# JWT-Wallet-Authentication-Gin
It is becoming a common practice in the web3 world to let your users sign in using their wallet. handy for fast onboarding into your product.
I will walk through the different taken to make this process come true

### Prerequisites

The things you need before installing the software.

* Golang
* Gin
* JWT
* go=ethereum
* SIWE - Signing In With Ethereum


### Installation

A step by step guide that will tell you how to get the development environment up and running.

```
# Clone this repository
$ git clone https://github.com/AbdulrahmanDaud10/JWT-Etherium-Wallet-Authentication

# Go into the repository
$ cd JWT-Etherium-Wallet-Authentication

# Install dependencies
$ go mod tidy

# Run the app
$ go run main.go
```

## Getting Started

To achieve our goal we need 2 endpoints. First, `/nonce/:address` to generate a nonce for our wallet address, second, `/signin` to either sign-in or sign-up our user.

## Snippet Example

````
func Nonce(c *gin.Context) {
    // in my App I just have my db and redis
	app := c.MustGet("app").(*config.App)
	address := c.Param("address")

    // we check if it's a valid EVM like address
	if !evm.IsValidAddress(address) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid address format"})
		return
	}

    // we generate a nonce with SIWE
	nonce := siwe.GenerateNonce()

    // save the address annd nonce in redis
	err := app.Rdb.Set(ctx, address, nonce, 1*time.Minute).Err()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"nonce": nonce,
	})
}
````

## Enviroment Variable Example
````
REDIS_ADDR 		= "localhost:6379"
REDIS_PASSWORD 	= "xxxx"
DB_DSN  		=   "host=localhost user=admin password=supersecret dbname=db port=5432 sslmode=disable"
RPC_URL 		=   "https://eth-mainnet.g.alchemy.com/v2/apikey"
JWT_SECRET  	= "jflksdklfklsdjflsdjlfkjsdlkfjsdlkfjlsdkjfldksj999999"
````