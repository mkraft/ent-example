# ent-example

Simple Ent example: one model with two fields.

```shell
go run main.go
```
```shell
after create
	User(id=1, name=Jane Doe, type=admin)
before update
	User(id=1, , )
after update (fields are not persisted as null)
	User(id=1, name=Jane Doe, type=admin)
after update with ClearType and ClearName (fields are persisted as null)
	User(id=1, , )
```