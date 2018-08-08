# wosr
Code samples for Women of Silicon Roundabout talk, October 2017.

To run this code:

Download and install Go - see https://golang.org/dl/ for instructions.

Download and install git, if you don't already have it - see https://git-scm.com/downloads for instructions.

Run `go get github.com/sjwells/wosr`

Go to `$GOWORKSPACE/src/github.com/sjwells/wosr/helloworld`

Run `go run helloworld.go`

You should see output like this:

    dlrow olleH

Then feel free to explore the other code samples!

## helloworld
Reverses a string and prints it to standard out

## functions
Calls a function that returns two values

## rectangles
Calculates the area of a rectangle, two ways

## simpleconcurrency
Kicks off a separate go routine. Doesn't wait for it.

## concurrency
Kicks off five separate go routines. Waits for the first one that says it's finished (uses channels)

## webservice
A small webservice that returns author information as json. Make a request using:

`curl -s 'http://localhost:8000/authors/456'`
