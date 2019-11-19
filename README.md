# forgan

File ORGANize utils written with Go.

## Requirements

- [Go](https://golang.org/)

## Installation

### Homebrew

```bash
$ brew install muniere/triv/forgan
```

### Manual

```bash
# clone
$ git clone git@github.com:muniere/forgan.git

# install
$ cd forgan
$ make && make install
```

## Usage

### Count entries in directory

```bash
# default
$ count dir1 dir2 dir3

# filter with pattern
$ count --pattern="*.txt" dir1 dir2 dir3

# include hidden files
$ count --all dir1 dir2
```

### Numberize filenames

```bash
# default
$ numberize *.jpg

# with options
$ numberize --start=1234 --length=4 --prefix="image_" *.jpg

# dry run
$ numberize -n *.jpg
```

### Randomize filenames

```bash
# default
$ randomize *.jpg

# with options
$ randomize --length=5 --prefix="image_" *.jpg

# dry run
$ randomize -n *.jpg
```
