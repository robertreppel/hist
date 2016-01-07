/*
Package filestore implement file system persistence for the hist eventstore.

Events are stored in files. Each aggregate type is a directory. Each aggregate instance is a file, with events appended
when they are saved. For example, given a data directory "/data", a "User" aggregate and a user with id "12345", when an
"EmailChanged" event is saved it is appended to:

  /data/events/User/12345.events
*/
package filestore
