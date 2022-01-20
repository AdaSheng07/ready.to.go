## Go Module 进阶

### 🔸 Module 的使用

- 像普通的包一样直接使用函数
- 编写代码时，会遇到找不到`package`、`function`的错误，一些较老的项目无法使用 Go Module，这时需要运行：
  ```
  go get <package> // as updated in golang 1.17
  go mod tidy
  ```
- 像正常一样使用代码
- 如果在 VS Code 中打开项目，注意`go.mod`文件务必位于 VS Code 锁打开文件夹的位置（golang 1.17），GoLand 中不存在这个问题
- 运行`go mod tidy`来保证项目中使用到的 go modules 被包含在项目依赖中

### 🔸 Go Module 替换（replace）

Go Module 替换（replace）是用另一个实现替换默认要使用的实现，类似于"狸猫换太子"。一般使用的场景有：
- 默认`go mod tidy`出来的都是最新版本，如果最新版本不合适，或不能使用等，需要将其替换为相对旧一些的版本。比如：
  ```
  replace github.com/spf13/cobra => github.com/spf13/cobra v1.2.0
  ```
- upstream代码（上游代码库）不太符合我们的需求，我们可以在此基础上做自身的定制，并开源在`git`服务上，需要将`module`重定向。在 Golang 中永远都是从源码出发的，可以直接从实现出发来使用。
- 本地有些项目的代码，没有`git`支持，需要作为独立的`module`使用。比如：
  ```
  replace learn.go.tools => ../learn.go.tools
  ```
> 【注意】本地项目代码作为`module`替换：
> 1. 替换的是`module`，寻找的是替换后的`go.mod`所在的路径位置**相对**于原来的`go.mod`所在文件夹的位置。
> 2. 可以通过`Terminal`验证`replace`时的重定向路径是否存在和正确，保证`module`能够成功`replace`。
> 3. `replace`之后的版本一定被使用。

> 【常见问题】使用`go mod`时报错`i/o timeout`
>
>      解决办法是在执行前在`Terminal`中设置环境变量：
>   ```
>   $ export GO111MODULE=on
>   $ export GOPROXY=https://goproxy.io 
>   或
>   $ export GO111MODULE=on
>   $ export GOPROXY=https://goproxy.cn
> 
>   对于Golang 1.13及以上版本，还可以：
>   $ go env -w GOPROXY=https://goproxy.cn
>   ```

### 🔸 Vendor

Vendor 是把项目`module`定义的所有依赖制作一个副本保存在项目的`vendor`目录。可以理解为：将项目定义的依赖做一个**快照**并保存下来，避免项目依赖的变更影响项目的一致性。Vendor 是几乎所有大型项目中都会使用的，稳定性强，保证项目的一致性。

Go Module 深度支持`vendor`，通过在`go.mod`所在目录下运行命令行`go mod vendor`将项目的依赖保存到`vendor`中：
  ```
  $ go mod vendor
  $ go mod vendor -v
  # github.com/inconshreveable/mousetrap v1.0.0
  ## explicit
  github.com/inconshreveable/mousetrap
  # github.com/spf13/cobra v1.3.0
  ## explicit; go 1.15
  github.com/spf13/cobra
  # github.com/spf13/pflag v1.0.5
  ## explicit; go 1.12
  github.com/spf13/pflag
  ```
Go Module 的定义与 `vendor`本身就是**共存**的，通过`go build/run`的命令行选项来控制`-mod=<vendor|mod|readonly>`：
- `vendor`——使用`vendor`中的依赖编译项目
- `mod`——使用`go module`定义的依赖编译项目，并且会自动更新`go.mod`定义
- `readonly`——使用`go module`定义的依赖编译项目，并且不做任何依赖的升级

> 【注意】`vendor`目录存在/不存在：
> 1. 项目中如果有`vendor`目录，在编译时默认使用`vendor`提供的依赖，如果在`Terminal`中运行了`go mod tidy`后，必须要再运行`go mod vendor`才能使依赖生效
> 2. 如果没有`vendor`目录，在编译时默认使用`readonly`
> 3. 在开发时最好先将目录中的`vendor`目录删除，在最后提交代码前再对`go.mod`文件执行`go mod vendor`
