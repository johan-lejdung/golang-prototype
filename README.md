# golang-proto
A GoLang prototype that contains an example of how one might use Protobuf together with either gRPC or PubSub to communicate with other services and clients 

# Installation

Install golang: https://golang.org/doc/install

If not already done, add the following lines to ~/.bach_profile or ~/.zshrc (if using zsh)

    export GOPATH=/Users/<username>/go

    export PATH=$GOPATH/bin:$PATH

Where <username> is the username of your profile.

Run: 

Then install govendor:

go get -u github.com/kardianos/govendor

Run this command in the correct folder:

    govendor init
    govendor add +external

A file named *general.env* must exist, with data similar to that in *general.env.template*

# Future

* JWT for auth, https://jwt.io/
* Write simple tests √
* More packages √
* Analyze comment with perspective API, https://www.perspectiveapi.com/ √
* Handle the analyze, bring out best data
