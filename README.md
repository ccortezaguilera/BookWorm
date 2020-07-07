# BookWorm

A cli tool for tracking books read, owned, or should buy. 

```
go install BookWorm
```

## Commands

### Add
```sh
BookWorm add --title="1984" --read # will add the book called 1984
```

### List
```sh
BookWorm list  # show the first 20 records
```

or 

```sh
BookWorm list -n 6 # will only show the first 6 records
```
### Del
```sh
BookWorm del --identity=1  # deletes the record with id = 1
```

### Update
```sh
BookWorm update --identiy=1 --read --own # updates the record with id = 1 to own and has read
```
