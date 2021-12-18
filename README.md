# go-study
record go study data

## golang base knowledge
### struct
- 理解
    - 数组是一个同一数据类型的数据集合；结构体可以是同一类型、也可以不是同一类型数据的集合；
- 结构体定义
    - 关键字：type和struct
        - struct语句：定义一个新的数据类型
        - type语句：设定结构体的名称
    - 定义结构体
    ~~~
        type struct_var_name struct {
            member type
            ……
        }

        // 示例
        type Book struct {
            id int
            title string
            author string
            price float32
        }
    ~~~
    - 访问结构体成员：使用.操作符号
    - 结构体作为函数的参数
    - 结构体指针

### sync包
- 互斥锁：sync.Mutex
- sync.WaitGroup
- select和channel

### 构建可复用的模块
- 包定义
- 包名和目录名的关系
- 包名和构建后的文件关系
- go依赖安装
    - go get命令：导入github上的包的时候，使用go get命令，不需要加schema和.git文件，下载的包会被放在GOPATH所在的路径，加上~~~-u~~~参数，可以不使用本地GOPATH已经存在的包，强制使用最新的包，例如~~~go get github.com/easierway/concurrent_map~~~
