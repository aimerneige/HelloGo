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
export GOPATH="$HOME/Code/golang/"
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

## 第一个程序



