# Go 语言入门

## Go 语言的优势

- 编译速度很快
- 协程 (goroutine) 和信道 (channel)
- 语法糖
- 拥有垃圾收集器 (GC)
- 一致性
- 适合网络编程

## 配置环境

### 下载安装 go

在 [golang的官网](https://golang.org/) 找到 [下载界面](https://golang.org/dl/) 下载对应系统版本的安装包，并设置环境变量。（下载需要代理服务器）

以 `Mac OS` 系统和 `zsh` shell 为例，你需要在 `~/.zshrc` 下写入下面这样的环境变量：（不要直接复制，按照自己的安装路径和文件路径修改）

```bash
export GOROOT="/usr/local/go"
export GOPATH="$HOME/Code/golang"
export PATH="$GOROOT/bin:$GOPATH/bin:$PATH"
```

其中，第一行表示 go 的**安装路径**，第二行表示 go 的**工作区目录**，第三行表示**可执行文件**的路径。

修改好环境变量后，可以使用下面的指令来查看当前安装的 go 语言版本，以检查是否正确安装。

```
go version
```

注意，go 语言将所有的 go 文件放在同一个 _工作区_ 中，你需要新建一个文件夹来保存所有的源代码。下面的例子展示了一个工作区的状态。

```bash
bin/
    streak
    todo
pkg/
    code.google.com/p/goauth2/
        oauth.a
    github.com/nf/todo/
        task.a
src/
    code.google.com/p/goauth2/
        .hg/
        oauth/
            oauth.go
            oauth_test.go
    github.com/nf/
        streak/
        .git/
            oauth.go
            streak.go
        todo/
        .git/
            task/
                task.go
            todo.go
```

你必须在工作区目录下新建三个文件夹：

1. bin
2. pkg
3. src

这三个文件夹分别用来存储编译好的可执行文件、包和源代码文件。

### 编辑器配置

#### vscode

安装好 `vscode` 后需要下载一个插件，直接搜索 `go` 下载即可。

安装好插件后关闭 `vscode`，在工作区中新建一个文件，比如 `hello.go` ，使用 `vscode` 打开它， `vscode` 会询问你是否要通过 `go get` 安装一些依赖，我们可以选择 `Install All` 来安装它们，它们可以提供一些格式化，代码提示等等这些功能。稍微等待安装完成后，我们就可以正式使用 `vscode` 写 `go` 语言程序了！

由于 `GWF` 的封锁，你可能无法下载这些依赖，你需要准备一个代理服务器。但是即使开了代理依然不能下载是为什么呢？这是因为 `go get` 默认使用 `git` 作为版本管理工具，首先你需要为 `git` 设置代理，但是即使你配置了 `git` 的代理依然无法安装依赖，这是因为有一些依赖使用的版本控制工具是 `svn` 而不是 `git`。 比如你要下载 `golang.org/x/tools/cmd/guru` 这个包，为了确定使用的是 `svn` 还是 `git`，  `go get` 会访问 `https://golang.org/x/tools/cmd/guru?go-get=1` 这个网站来获取版本库的类型，这个请求也是需要代理的，所以还需要额外设置终端的代理。

如果你需要设置终端代理，将下面的内容保存在 `~/.zshrc` 中即可 （端口改为你自己的端口）

```bash
export https_proxy=socks5://127.0.0.1:8123
export http_proxy=socks5://127.0.0.1:8123
```

如果你需要设置 git 的代理，直接执行下面的命令：

```bash
git config --global http.proxy 'socks5://127.0.0.1:8123'
git config --global https.proxy 'socks5://127.0.0.1:8123'
```

当然你也可以修改 `~/.gitconfig`

```conf
[http]
	proxy = socks5://127.0.0.1:8123
[https]
	proxy = socks5://127.0.0.1:8123
```

<!-- FUCK GFW -->

#### goland

这个 IDE 开箱即用，没什么好说的。

> 本项目完全使用 vscode 开发。

## 第一个程序

首先为程序创建一个目录，你需将它放在你的 _工作区目录_ 下，比如这个项目我放在了 `$GOPATH/src/aimerneige.com/HelloGo/` 下，之后就可以开始写代码了。

首先惯例当然是要在终端打印一个 `Hello World!`.

我们创建一个 `hello.go` 的文件来写入下面的代码：

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

你并不需要写 `import "fmt"` 这一行，你也不需要特别在意代码的格式化，当你保存文件之后，go 会帮你自动导包和格式化。

当代码是 main 包并且有一个 main 函数的时候，它可以被编译为可执行文件。之后我们可以在源代码存储的路径下通过下面的指令直接运行它：

```bash
go run hello.go
```

如果你想要编译它，执行下面的指令：

```bash
go build hello.go
```

在当前目录下你会得到一个与当前系统相关的可执行文件，你可以直接执行它，也会得到终端输出。

构建完成后，如果你想要吧这个可执行文件安装到系统，在源文件的目录下执行下面的指令：

```
go install
```

执行完这个指令后，你会发现在任何路径下执行 `hello` 都可以得到输出。

我们可以使用下面的命令查看可执行文件被存储在哪了：

```bash
which hello
```

我们会发现这个文件被保存在了 _工作区_ 下的 `bin` 文件夹中，而这个文件夹已经在之前配置环境中被加入到了环境变量中，所以我们可以直接执行它。

如果不需要了，可以直接删除它。

## 第一个库

首先创建一个新的目录用来保存我们的库文件。本项目中我新建了一个名为 `stringutil` 的文件夹，在下面创建新的文件 `reverse.go`，写入如下内容：

```go
package stringutil

// Reverse 将其实参字符串以符文为单位左右反转。
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
```

我们会发现函数前有一行注释，如果不写的话，vscode 会弹出一个警告，要求你写注释，而且必须以函数名开头。之后加上解释。

写好之后，我们可以通过下面的指令来测试该包的编译：

```bash
go build aimerneige.com/HelloGo/stringutil
```

如果当前目录就是源码所在目录，可以直接通过下面的指令编译：

```bash
go build
```

由于没有主函数，编译并不会产生任何结果，如果你没有看到输出，证明它没有问题并可以通过编译。如果想要输出文件，可以通过在源代码所在目录下执行下面的指令在工作区的 pkg 目录中生成包的对象：

```bash
go install
```

以本项目为例，在 `$GOPATH/pkg/darwin_amd64/aimerneige.com/HelloGo/` 目录下可以看到编译后的库文件 `stringutil.a`。

之后，我们可以在其他项目中直接导入这个库。

新建一个 `pkghello` 的文件夹，复制之前写的第一个程序中的代码，并修改，得到下面的代码：

```bash
package main

import (
	"fmt"

	"aimerneige.com/HelloGo/stringutil"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(stringutil.Reverse("Hello World!"))
}
```

如果我们要导入多个包的时候，我们需要使用括号。之后就可以使用刚才写的 Reverse 函数了。

直接运行后构建后执行这个文件，我们可以得到一个正序的 `Hello World!` 和一个反序的 `!dlroW olleH`。

当我们在构建成功后通过 `go install` 安装新的 `hello` 程序时，go 工具会安装它所依赖的任何东西，之后再执行 `hello`，我们依然可以得到相同的结果。

当我们通过下面的指令安装的时候，`stringutil` 包也会被自动安装。

```
go install aimerneige.com/HelloGo/pkghello
```

如果你和我使用同样的环境，当成功的完成安装之后，工作空间应该是这样的：

```bash
bin/
    pkghello # 可执行文件
pkg/
    darwin_amd64/ # 这里会反映出你的操作系统和架构
        aimerneige.com/
            HelloGo/
                stringutil.a # 包对象
src/
    aimerneige.com/
        HelloGo/
            stringutil/
                reverse.go # 包源码
            pkghello/
                hello.go # 命令源码
```

Go 的可执行命令是静态链接的；在运行Go程序时，包对象无需存在。

## Go 的包管理

Go 中的包实际上就是计算机中的目录，或是叫文件夹，通过它们进行目录结构和文件的组织。Go 只是将文件目录称为了包而已。Go 语言中，包名和文件所在的目录名是一致的。

### 包的命名

#### Go 语言的包命名

规范：

1. 简洁
2. 小写
3. 和 Go 文件所在目录同名

这样的命名规范有助于我们引用，书写以及快速定位查找。

```go
package main

import "net/http"

func main() {
	http.ListenAndServe("127.0.0.1:8080", handler)
}
```

#### 以域名命名包

如果你有自己的域名，你可以使用自己的域名命名包：

```go
package main

import "aimerneige.com/utils"
```

#### 使用 GitHub 账号命名包

如果你没有自己的域名，你也可以使用 GitHub 账号的方式命名：

```go
package main

import "github.com/aimerneige/utils"
```

由于笔者有自己的域名 `aimerneige.com` ，在本项目中我使用 `以域名命名包` 的方式。

### main 包

当把一个 go 文件的包声明为 `main` 时，就等于告诉 go 编译程序，我这个是一个可执行的程序，那么 go 编译程序会尝试把它编译为一个二进制的可执行文件。

一个 `main` 的包，一定要包含一个 `main()` 函数。它是程序的入口，没有这个函数，程序无法执行。

> 在 go 语言里，同时要满足 `main` 包和包含 `main()` 函数，才会被编译为一个可执行文件。

### 导入包

```go
import "fmt"

import (
    "fmt"
    "net/http"
)
```

编译器首先会优先在 `$GOROOT` 下查找，其次是 `$GOPATH`，一旦找到，会立即停止搜索，如果最终都没有找到，就会报编译异常。

### 远程包导入

```go
import "github.com/biezhi/moe"
```

1. 在 `$GOPATH` 下搜索
2. `go get` 下载
3. 保存在 `$GOPATH/src` 下以便于之后使用

### 命名导入

如果不同的包之间重名了，我们可以通过下面的方式对包进行重命名：

```go
package main

import (
    "fmt"
    myfmt "mylib/fmt"
)

func main() {
    fmt.Println()
    myfmt.Println()
}
```

有时候，如果有一个包没有用到，但是却要导入它，就可以通过重命名为下滑线都形式来导入：

```go
package main
import (
    _ "mylib/fmt"
)
```

### 包的 init 函数

#### init 函数

每个包可以有任意多的 `init` 函数。

`init` 函数会在 `main` 函数执行之前执行，通常我们会用 `init` 函数来初始化变量，设置包，读取一些配置等。

```go
package mysql

import (
    "database/sql"
)

func init() {
    sql.Register("mysql", &MySQLDriver{})
}
```

如果我们只想执行一个包的 init 方法，并不想使用这个包，我们可以在导入这个包的时候，使用 `_` 重命名包，避免编译错误。

#### 静默导入

```go
import "database/mysql"
import _ "github.com/go-sql-driver/mysql"

db, err != sql.Open("mysql", "user:password@/dbname")
```

## Go 命令

在终端输入 go 并回车，你会得到下面的输出

```bash
➜  ~ go
Go is a tool for managing Go source code.

Usage:

	go <command> [arguments]

The commands are:

	bug         start a bug report
	build       compile packages and dependencies
	clean       remove object files and cached files
	doc         show documentation for package or symbol
	env         print Go environment information
	fix         update packages to use new APIs
	fmt         gofmt (reformat) package sources
	generate    generate Go files by processing source
	get         add dependencies to current module and install them
	install     compile and install packages and dependencies
	list        list packages or modules
	mod         module maintenance
	run         compile and run Go program
	test        test packages
	tool        run specified go tool
	version     print Go version
	vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

	buildconstraint build constraints
	buildmode       build modes
	c               calling between Go and C
	cache           build and test caching
	environment     environment variables
	filetype        file types
	go.mod          the go.mod file
	gopath          GOPATH environment variable
	gopath-get      legacy GOPATH go get
	goproxy         module proxy protocol
	importpath      import path syntax
	modules         modules, module versions, and more
	module-get      module-aware go get
	module-auth     module authentication using go.sum
	module-private  module configuration for non-public modules
	packages        package lists and patterns
	testflag        testing flags
	testfunc        testing functions

Use "go help <topic>" for more information about that topic.
```

### go build

```bash
usage: go build [-o output] [-i] [build flags] [packages]
```

```bash
go build
go build .
go build hello.go
```

编译指定包

```bash
go build aimerneige.com/HelloGo/stringutil
```

编译所有包

```bash
go build aimerneige.com/HelloGo/stringutil/...
```

查看环境信息

```bash
➜  ~ go env
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/Users/aimerneige/Library/Caches/go-build"
GOENV="/Users/aimerneige/Library/Application Support/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOINSECURE=""
GOMODCACHE="/Users/aimerneige/Code/golang/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="darwin"
GOPATH="/Users/aimerneige/Code/golang"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
GCCGO="gccgo"
AR="ar"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/63/vqnp4j053z91rlqrms0lpfph0000gn/T/go-build667255972=/tmp/go-build -gno-record-gcc-switches -fno-common"
```

`GOOS` 是目标操作系统，它的值为：

- darwin
- freebsd
- linux
- windows
- android
- dragonfly
- netbsd
- openbsd
- plan9
- solaris

`GOARCH` 是目标处理器的架构，目前支持的有：

- arm
- arm64
- 386
- amd64
- ppc64
- ppc64le
- mips64
- mips64le
- s390x

跨平台编译 / 跨平台编译

```bash
GOOS=linux GOARCH=amd64 go build aimerneige.com/HelloGo/hello
```

> 更多内容 https://golang.org/doc/install/source#enviroment
> 
> 也可以在终端执行如下指令查看帮助信息
> 
> ```bash
> go help build
> ```

### go clean

用于清除可执行文件等编译结果

```bash
usage: go clean [clean flags] [build flags] [packages]
```

```bash
go help clean
```

### go run

我们想要运行一个源文件时，首先要 `build`，接下来需要执行可执行文件，而 `go run` 是吧这俩步合成了一步，直接得到结果。

```bash
usage: go run [build flags] [-exec xprog] package [arguments...]

Run compiles and runs the named main Go package.
Typically the package is specified as a list of .go source files from a single directory,
but it may also be an import path, file system path, or pattern
matching a single known package, as in 'go run .' or 'go run my/cmd'.

By default, 'go run' runs the compiled binary directly: 'a.out arguments...'.
If the -exec flag is given, 'go run' invokes the binary using xprog:
	'xprog a.out arguments...'.
If the -exec flag is not given, GOOS or GOARCH is different from the system
default, and a program named go_$GOOS_$GOARCH_exec can be found
on the current search path, 'go run' invokes the binary using that program,
for example 'go_js_wasm_exec a.out arguments...'. This allows execution of
cross-compiled programs when a simulator or other execution method is
available.

The exit status of Run is not the exit status of the compiled binary.

For more about build flags, see 'go help build'.
For more about specifying packages, see 'go help packages'.

See also: go build.
```

接收参数

假设我们有这样的一个文件 `hello.go`

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Input Args:", os.Args[1]);
}
```

我们使用下面的指令来执行它，并给它传递一个参数：

```bash
go run hello.go HelloWorld
```

之后我们会在终端得到 `Input Args: HelloWorld`。

### go install

安装编译结果。

```bash
➜  ~ go help install
usage: go install [-i] [build flags] [packages]

Install compiles and installs the packages named by the import paths.

Executables are installed in the directory named by the GOBIN environment
variable, which defaults to $GOPATH/bin or $HOME/go/bin if the GOPATH
environment variable is not set. Executables in $GOROOT
are installed in $GOROOT/bin or $GOTOOLDIR instead of $GOBIN.

When module-aware mode is disabled, other packages are installed in the
directory $GOPATH/pkg/$GOOS_$GOARCH. When module-aware mode is enabled,
other packages are built and cached but not installed.

The -i flag installs the dependencies of the named packages as well.

For more about the build flags, see 'go help build'.
For more about specifying packages, see 'go help packages'.

See also: go build, go get, go clean.
```

### go get

下载 Go 依赖库

```bash
go get github.com/biezhi/moe
```

更新依赖

```bash
go get -u github.com/biezhi/moe
```

查看进度

```bash
go get -v github.com/biezhi/moe
```

> 更多内容请在终端执行 `go help get`

### go fmt

可以格式化源代码中的排版。在 `vscode` 下，保存代码后会自动对代码执行格式化。

如果你想要查看它的作用，可以使用其他编辑器（比如没有经过配置的`vim`）调整排版，然后使用 `go fmt` 来格式化代码。

你可以像这样来格式化代码：

```bash
go fmt hello.go
```

`go fmt` 可以为我们统一代码风格，所有的代码都是统一的，这可以便于团队协作，所以我们在将代码提交到 `GitHub` 前一定要进行格式化。幸运的是，vscode 自动帮我们做了这件事。

### go vet

用于检查代码中的错误。

1. `Println` 这类的函数调用时，类型匹配了错误的参数。
2. 定义常用的方法时，方法签名错误。
3. 错误的结构体标签。
4. 没有指定字段名的结构字面量。

```bash
usage: go vet [-n] [-x] [-vettool prog] [build flags] [vet flags] [packages]
```

### go test

这个命令用于运行 Go 的单元测试，它也是接受一个包名作为参数，如果没有指定，使用当前目录。

go test 运行的单元测试必须符合 go 的测试要求。

1. 写有单元测试的文件名，必须以 `_test.go` 结尾。
2. 测试文件要求包含若干个测试函数。
3. 这些测试函数要以 Test 为前缀，还要接受一个 *testing.T 类型的参数。

```go
package main

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) == 3 {
		t.Log("1 + 2 = 3")
	}
	if Add(1, 1) == 3 {
		t.Error("1 + 1 = 3")
	}
}
```

> 更多内容请在终端执行 `go help test`

## 语法结构

### 文件名、关键字、标识符

#### Go 标记

Go 程序可以由多个标记组成，可以是关键字，标识符，常量，字符串，符号。

```go
package main

import "fmt"

func main() {
	fmt.Println("hello, world")
}
```

#### 行分隔符

在 Go 中，一行代表一个语句，行尾不需要分号 `;`，这是因为编译器会帮我们处理。

以下为俩个语句：

```go
	fmt.Println("Hello World!")
	fmt.Println("https://aimerneige.com/")
```

#### 注释

Go 语言中的注释和 C/Java 等语言相同，包含单行和多行注释。

```go
// 这是个单行注释

/*
这是多行注释
我有很多行
*/
```

#### 标识符

规范与 C/C++ 等语言相同，以字母、数字和下划线组成，其中数字不能开头。

- 不能以数字开头
- 不能使用 `$` 等特殊符号
- 不能是 Go 的关键字

#### 关键字

|          |             |        |           |        |
| -------- | ----------- | ------ | --------- | ------ |
| break    | default     | func   | interface | select |
| case     | defer       | go     | map       | struct |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continue | for         | import | return    | var    |

#### 预定义标识符

|        |         |         |         |        |         |           |            |         |
| ------ | ------- | ------- | ------- | ------ | ------- | --------- | ---------- | ------- |
| append | bool    | byte    | cap     | close  | complex | complex64 | complex128 | uint16  |
| copy   | false   | float32 | float64 | imag   | int     | int8      | int16      | uint32  |
| int32  | int64   | itoa    | len     | make   | new     | nil       | panic      | uint64  |
| print  | println | real    | recover | string | true    | uint      | uint8      | uintptr |

程序一般由关键字、常量、变量、运算符、类型和函数组成。

程序中可能会使用到这些分隔符： 括号`()` 中括号`[]`和大括号`{}`

程序中可能出现这些标点符号： `.` `,` `;` `:` `...`

如果需要将多个语句写在同一行，必须使用分号来区分。但是非常不推荐这样写。

### 包的导入和可见性

每一段代码只会被编译一次

一个 Go 程序是通过 `import` 关键字将一组包链接在一起。

如果需要导入多个包，它们可以被分别导入：

```go
import "fmt"
import "os"
```

或

```go
import "fmt"; import "os"
```

当然也有更加优雅的方法（推荐）：

```go
import (
	"fmt"
	"os"
)
```

该方法被称为因式分解关键字，该方法同样可以适用于 `const` `var` 和 `type` 的声明或定义。

#### 可见性规则

方法首字母小写 private 对外不可见
方法首字母大写 public 对外可见

例如我们在 `syntaz` 目录下的 `main.go` 中尝试调用 `pkg1` 下定义在 `package01.go` 中的函数时，`Foo()` 由于首字母大写，对外是 `public` 的，可以成功调用，而 `bar()` 由于首字母小写，对外是 `private` 的，无法调用。

### Go 程序的一般结构

Go 程序的整体结构如下：

1. 包的 import
2. 常量、变量和类型的定义和声明
3. 如果存在 init 函数，则定义 init 函数
4. 如果当前包是 main 包，定义 main 函数
5. 然后定义其余的函数
   1. 首先定义类型的方法
   2. 按照 main 函数中调用的先后顺序来定义相关函数
   3. 如果有很多函数，还可以使用字母顺序进行排序

```go
package main

import (
	"fmt"
)

const c = "C"

var v int = 5

type T struct {}

func init() { // 包级别的初始化
}

func main() {
	var a int
	Func1()
	// ...
	fmt.Println(a)
}

func (t, T) Method1() {
	// ...
}

func Func1() { // 导出函数 Func1
	// ...
}
```

Go 程序的执行（程序启动）顺序如下：

1. 按顺序导入所有被 main 包引用的其他包，然后在每个包中执行如下流程：
2. 如果该包又导入了其他的包，则重第一步开始递归执行，但是每个包只会被导入一次。
3. 然后以返回的顺序在每个包中初始化常量与变量，如果该包含有 init 函数的话，则调用该函数。
4. 在完成这一切后，main 也执行同样的过程，最后调用main函数开始执行程序。

### 类型转换

在 Go 中不存在隐式类型转换，因此所有的转换都必须显示说明，就像调用一个函数一样（类型在这里的作用可以看作是一种函数）：

```go
valueOfTypeB = typeB(valueOfTypeA)
```

类型B的值 = 类型B（类型A的值）

示例：

```go
a := 5.0
b := int(a)
```

但这只能在定义正确的情况下抓换成功，例如从一个取值范围较小的类型转换到一个取值类型较大的类型（例如将 int16 转换为 int32）。

当从一个取值范围较大的转换为一个取值范围较小的类型时（例如将 int32 转换为 int16），会触发精度丢失（截断）的情况。当编译器捕捉到非法的类型转换时会引发编译时错误，否者将引发运行时错误。

具有相同底层类型的变量之间可以相互转换：

```go
var a IZ = 5
c := int(a)
d := IZ(c)
```

### 命名规范

干净、可读的代码和简洁性是 Go 最求的目标。通过 gofmt 来强制实现统一的代码风格。

Go 语言中对象命名也应该是简洁且有意义的。

像 `Java` 和 `Python` 中那样使用混合着大小写和下划线的冗长的名称会严重降低代码的可读性。名称不需要指出自己所属的包，因为在调用的时候会使用包名作为限定符。

返回某个对象的函数或方法的名称一般都是使用名词，没有 `Get...` 之类的字符，如果是用于修改某个对象，则使用 `SetName`。

有必须要的话可以使用大小写混合的方式，如 `MixedCaps` 或 `mixedCaps`，而不是使用下划线来分割多个名称。
