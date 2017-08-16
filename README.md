# golang-prototype
A GoLang prototype that contains an example of how one might use Protobuf together with either gRPC or PubSub to communicate with other services and clients 

# Installation

Install golang: https://golang.org/doc/install

*Sidenote: GoLang keeps all of it's code (and your code) in a huge rootfolder, to more easily reuse common packages. This is was is refered as GOPATH.*

If not already done, add the following lines to ~/.bach_profile or ~/.zshrc (if using zsh)

    `export GOPATH=/Users/<username>/go`

    `export PATH=$GOPATH/bin:$PATH`

Where <username> is the username of your profile.

**Go to your GOPATH/src and then clone this repo:**

    `git clone git@github.com:Storytel/golang-prototype.git`

**Then install govendor:**

    `go get -u github.com/kardianos/govendor`

**Type `govendor sync` in both**:
 - `golang-prototype/consumer`
 - `golang-prototype/producer`
 
**Then go back to root `golang-prototype/`**


A file named *general.env* must exist. Copy *general.env.template* and rename the copy to *general.env*.

**Run `docker-compose up`**

# Testing

There are 2 endpoints to the producer. You can easily access these using the program [postman](https://www.getpostman.com/postman). 

**Endpoints:**
 - `localhost:8081/produce/grpc/msg`
 - `localhost:8081/produce/pubsub/msg`
 
The first one will transport the message sent via gRPC and the second one will do it using the PubSub Emulator.

In postman make sure the header is:
`Content-Type` : `application/json`

And the body is:
`{
	"msg" : "This is a custom message and sender",
	"sender": 4
}`
