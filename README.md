# go-uuid-qrcode

generate uuid and save as PNG of QRCode(H).


## Build

```
$ go get github.com/mattn/gom
$ go get github.com/mitchellh/gox

$ gom install
$ gom exec gox -output "dist/qruuid_{{.OS}}_{{.Arch}}/qruuid"
```

## command qruuid

```
$ ./qruuid
e1018559-cc0a-4a29-a6ac-2369f626e92c

$ ls
e1018559-cc0a-4a29-a6ac-2369f626e92c.png
```

### Windows

click `qruuid.exe`.


## Release(note for me.)

master branch is released automaticaly to `giraffi/go-uuid-qrcode/releases/tag/master` by travis.

### stable release

```
go get github.com/tcnksm/ghr
find dist/* -type d -exec zip -r {}.zip {} \;
gom exec ghr --username giraffi --token $GITHUB_TOKEN v0.2.0 dist/
```

## LICENSE

MIT
