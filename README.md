# UPX

upx is a simple cli utility and library for parsing and extracting unity packages.  
it also contains functionality to examine unity packages for suspicious files (scripts and dlls)  
they are suspicious because they can contain malware. when i download an outfit on booth.pm it shouldn't come with any scripts

## Import

`go get github.com/elianel/upx`

## Usage

```sh
$ ./upx

Usage: upx [options] [command] ...

Options:
  -h, --help Help

Commands:
  s         examine sus level of unity package
  x         extract unity package

```

## Example

`./upx x --src outfit.unitypackage --dst ./outfit`


