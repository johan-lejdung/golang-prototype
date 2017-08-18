# golang-prototype
A GoLang prototype that contains an example of how one might use Protobuf together with either gRPC or PubSub to communicate with other services and clients.

**Resources:**
- [How to Write Go Code](https://golang.org/doc/code.html)
- [A Tour of Go (Interactive coding tour)](https://tour.golang.org/welcome/1)

# Covers

- JWT Token creation & authentication
- Protobuf
- gRPC
- PubSub

# Installation

Install golang: https://golang.org/doc/install

*Sidenote: GoLang keeps all of it's code (and your code) in a huge rootfolder, to more easily reuse common packages. This is was is refered as GOPATH.*

If not already done, add the following lines to ~/.bach_profile or ~/.zshrc (if using zsh)

    `export GOPATH=/Users/<username>/go`

    `export PATH=$GOPATH/bin:$PATH`

Where `<username>` is the username of your profile.

**Go to your GOPATH/src and then clone this repo:**

    `git clone git@github.com:Storytel/golang-prototype.git`

**Then install govendor:**

*Sidenote: govendor is used to keep packages locally in projects rather than in the GOPATH, it also works much like npm/yarn*

    `go get -u github.com/kardianos/govendor`

**Type `govendor sync` in both**:
 - `golang-prototype/consumer`
 - `golang-prototype/producer`
 
**Then go back to root `golang-prototype/`**


A file named *general.env* must exist in subdirectories `consumer` and `producer`. Copy *general.env.template* and rename the copy to *general.env*.

**Run `docker-compose up`**

# Testing

There are 2 endpoints to the producer. You can easily access these using the program [postman](https://www.getpostman.com/postman). 

**Endpoints:**
 - `localhost:8081/produce/grpc/msg`
 - `localhost:8081/produce/pubsub/msg`
 - `localhost:8081/produce/jwt/fetch`
 - `localhost:8081/produce/jwt/auth`
 
**PubSub/gRPC**
 
`localhost:8081/produce/grpc/msg` will transport the message sent via gRPC and `localhost:8081/produce/pubsub/msg` will do it using the PubSub Emulator.

In postman make sure the header is:
`Content-Type` : `application/json`

And the body is:
`{
	"msg" : "This is a custom message and sender",
	"sender": 4
}`

**If everything work as expected you will recieve the same body in the response back!**
In the case of gRPC the response is the actual response from the `consumer`, via `producer`. That is not the case with PubSub, it's more of a confirmation in that usecase.

You can also observe the logs from the docker containers to see what is happening.

**For JWT**

Call `localhost:8081/produce/jwt/fetch` and save the Token from the response.

In Postman, add a header `Authorization` with value `Bearer <token>` where `<token>` is the token from the previous response.

The message in the response will say `All good, you are free to enter!` if the token is valid.
