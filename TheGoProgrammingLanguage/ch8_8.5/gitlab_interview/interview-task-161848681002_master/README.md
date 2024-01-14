# Dupe

List all duplicate files.

## How to use

### Create test data

```sh
mkdir /tmp/dupe
go run cmd/generate/generate.go -rootDir /tmp/dupe
```

### Check for duplicates

```sh
go run cmd/dupe/dupe.go -rootDir /tmp/dupe
```