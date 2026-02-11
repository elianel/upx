# UPX

CLI tool for extracting `.unitypackage` archives.

It can also scan and detect potentially unsafe files such as: `.cs` and `.dll`

These file types can execute code on import in the Unity Editor and may contain malicious logic.  
For example, cosmetic assets (e.g., outfits downloaded from booth.pm) should not include executable scripts or assemblies unless explicitly intended.

---
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


