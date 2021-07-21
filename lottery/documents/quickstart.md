## Models
xorm-cmd

go get xorm.io/xorm
github.com\go-sql-driver\mysql
go build -tags mysql
xorm reverse mysql "root:wzmmysql@tcp(127.0.0.1:3306)/lottery?charset=utf8" templates/goxorm
C:\Users\wangzhumo\go\src\github.com\go-xorm\cmd\xorm\xorm.exe reverse mysql "root:wzmmysql@tcp(127.0.0.1:3306)/lottery?charset=utf8" templates/goxorm


Step1:
```shell
go get github.com/go-xorm/cmd/xorm
go get github.com/go-xorm/xorm
```
需要特别注意的是，在gopaht下在，记得修改`GO111MODULE=on` ->> `GO111MODULE=auto`



Step2:
```shell
C:\Users\wangzhumo\go\src\github.com\go-xorm\cmd\xorm

go build
```
会生成`xorm.exe`

Step3:
根目录生成templates/goxorm目录
添加文件
- config
  ```text
  lang=go
  genJson=1
  prefix=
  ```

- struct.go.tpl
  ```js
    package {{.Models}}

    {{$ilen := len .Imports}}
    {{if gt $ilen 0}}
    import (
        {{range .Imports}}"{{.}}"{{end}}
    )
    {{end}}

    {{range .Tables}}
    type {{Mapper .Name}} struct {
    {{$table := .}}
    {{range .ColumnsSeq}}{{$col := $table.GetColumn .}} {{Mapper $col.Name}}	{{Type $col}} {{Tag $table $col}}
    {{end}}
    }

    {{end}}

  ```

Step4:
```shell
C:\Users\wangzhumo\go\src\github.com\go-xorm\cmd\xorm\xorm.exe reverse mysql "root:wzmmysql@tcp(127.0.0.1:3306)/lottery?charset=utf8" templates/goxorm
```