# Combination Generator

Generates random combinations of the lines in the file with a given length, and then writes the results to an output file.

## Usage

`input.txt` example:
```
abcd
123
pass
home
pc
foo
bar
bla
```

```sh
go run main.go input.txt output.txt 2
```

The generated combinations are separated by spaces and each combination is written on a new line in the output file.
