## go-resource

Resource compiler for Golang.

## DESCRIPTION

From time to time You need to add something like binary or text file
into target binary.

There are several ways to do that, but the package is very simple.

If You want to add any files into Your binary, follow the instruction:

1. Install `go-resource` CLI.

```
$ go install github.com/unera/go-resource@latest

$ go-resource version
0.0.1

```

2. Create infrastructure in your project:

```

$ go-resource create my-resources

2024/11/20 18:02:12 Directory `my-resources` created.
2024/11/20 18:02:12 Config `my-resources/resources.yaml` created.
2024/11/20 18:02:12 Please edit the file (my-resources/resources.yaml) and place your resources into the directory.
2024/11/20 18:02:12 After that run: go-resource build my-resources <file.go>.
2024/11/20 18:02:12 Note: You can use go:generate tags in Your files.
2024/11/20 18:02:12 example: //go:generate go-resource build my-resources resources.go

```

3. Edit resource config:

- Add exclude patterns into section: `resources.yaml/.ignore`
- Choose if resource getter returns error or panics:
  `resources.yaml/.use_errors`
- Choose package name for resource accessors. 

4. Put Your files into `my-resources/` directory (or its subdirectories)

5. Build go-file from resources:

```

$ go-resource build my-resource res.go
$ go-resource build my-resource > res.go   # the same

```

6. Add go:generate tag into Your file (if You want)

```
//go:generate -v go-resource build my-resource res.go
```

7. Build Your project.

To access resources from Your code You can use one of the following
patterns:

```go

internalYamlBlob := R("my-internal.yaml")   // returns []byte
internalYamlStr := Rs("my-internal.yaml")   // returns strinng


```

8. Enjoy
