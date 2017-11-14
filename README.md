# ffind
Find and delete hidden files/dir on Unix

This is a little hacky way of quickly listing and possibly deleting hidden files on Unix systems, i.e. systems where hidden files have their names starting with a dot.

This wont work on Windows!

## Building

```
$ git clone https://github.com/lyderic/ffind
$ cd ffind
$ go install
```

## Usage

- List all files, dirs, subdirs (including hidden files, dirs, subdir) in current directory:

```
$ ffind
```

- List all in <dir>

```
$ ffind dir
$ ffind ../relative/path/to/dir
$ ffind /absolute/path/to/dir
```

- List only hidden files, dirs, subdirs in current directory:

```
$ ffind -H
```

- List only hiddens in <dir>

```
$ ffind -H dir
$ ffind -H ../relative/path/to/dir
$ ffind -H /absolute/path/to/dir
```

- To remove all hidden files and dirs (!), please use the '-Hr' switch. Example:

```
$ ffind -Hr dir
```
