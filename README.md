# Hist: A Simple Eventstore in Go.

## Usage

See example.

```
cd example
go build
./example
```

## Design

Events are stored in files. Each aggregate type is a directory. Each aggregate instance is a file, with events appended
when they are saved. For example, given a data directory _"/data"_, a _"User"_ aggregate and a user with id _"12345"_, when an
"EmailChanged" event is saved it is appended to _"/data/events/User/12345.events"_

## Tests

Uses http://goconvey.co/. Run it to see BDD-style details about hist's business rules and behaviour.

## Production Use

Hist is considered alpha. Not recommended for production use.
