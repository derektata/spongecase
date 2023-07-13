# spongecase

<img src="docs/spongebob.png" width="200">

```
❯ ./spongecase -t 'easily generate spongebob-cased text'
```
EaSiLy GeNeRaTe SpOnGeBoB-cAsEd TeXt

## Usage

```
Usage of ./spongecase:
  -c, --clipboard     copy the output to the clipboard
  -f, --file string   the path to a file containing the text to convert
  -o, --overwrite     overwrite the input file
  -t, --text string   the text to convert
```

## Examples

```
# convert text and print to console
./spongecase -t "Hello, world!"

# convert text and copy to clipboard
./spongecase -t "Hello, world!" -c

# convert file content and copy to clipboard
./spongecase -f "/path/to/your/file.txt" -c

# convert file content and overwrite the file
./spongecase -f "/path/to/your/file.txt" -o
```

## License

MIT License