# createhandlerlib
Uses a text based file and creates a go package, which can be used as a http.Handler

## Install

To install this tool just use go get:

```go get github.com/as27/createhandlerlib```

## How to use

That an output package can be generated it is compulsory that the pkg flag is set otherwise the program exits with an error.

It is also compulsory to specify a source flag. At the moment the following source flags are availiable:

* __-url__ loads the source from the given url. The whole url (including http or https) as to be specified

Example:

```createhandlerlib -url http://example.com/cdn/javascriptlib.js -pkg jslib```