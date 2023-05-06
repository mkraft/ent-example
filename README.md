# ent-example

```shell
go run main.go
```
```shell
after Create	User(id=1, type=admin)	saved correctly
before Update	User(id=1, )		    type is nil
after Update	User(id=1, type=admin)	nil value is not saved by Update alone
after ClearType	User(id=1, )		    nil value is saved by ClearType
```