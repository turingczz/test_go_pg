# sql发布管理工具 gopg-migrations

## 1. gopg-migrations 是什么？
[gopg-migrations](github.com/go-pg/migrations/v7) 在我们使用数据库时，每次发布都要改一些表结构，或者修改数据之类的操作。

这些操作，手动执行既危险，又容易忘记；当新人看到奇怪的表结构，感觉到很诧异。如果我们能将所有操作管理起来，本地测试好之后再发布上去，那么就不会有上述这些问题。

因此，gopg-migrations这个库，给我们带来非常大的便利。当然，这里是对postgres数据库的支持。

## 2. 怎么使用

下载扩展库 `github.com/go-pg/migrations/v7` 即可使用。

2.1 假如我们先创建一个数据库 students，然后调用 migrations 去执行代码。

   步骤如下：
- 编写创建表的sql，使用migrations目录管理起来，num_xxx 表示按照num顺序执行
- 主逻辑
   
直接运行 go run main.go，此时代码运行，会生成students数据表：

sql如下：
![alt 条形图](https://swarm-gateways.net/bzz:/5aac64589b094632861794b07fc658ac9bef554d2956863798666f791dfa6f19/pg1.png)


执行后，会创建号数据表： students 和 gopg-migrations。
![alt 条形图](https://swarm-gateways.net/bzz:/a17cb08722a105678eb03210ebb785e7131e5d21316c69e3600cd47aeb2479e5/pg6.png)

查看 gopg-migrations如下：

![alt 条形图](https://swarm-gateways.net/bzz:/50e130d4a8337e545536cf249b435c83b527c38f0c3003808e79a7c7704e3bcf/pg2.png)



2.2 假设我们要增加sql语句，怎么做呢？

- 直接添加sql
- 重新运行
   
sql如下：
![alt 条形图](https://swarm-gateways.net/bzz:/290f87a7d93dd43651b5f3842d8059e763237551b5ca2fa44135873a7be25ecb/pg3.png)

执行如下：
![alt 条形图](https://swarm-gateways.net/bzz:/84ee116ba02b6480f63c9538571f881a98dca0b0cd7f40c3fdf0b64cb8963abb/pg4.png)

查看students数据表，发现1，2操作都已经执行了：
![alt 条形图](https://swarm-gateways.net/bzz:/3754e6aea9b6e359d4ed7dcb4360f16bd23e1abf01d5e29859b30c76abe0a6c4/pg5.png)

2.3.管理sql语句
类似1，2中的方式，我们可以添加更多的sql语句，把sql管理起来。



## 总结

[gopg-migrations](github.com/go-pg/migrations/v7) 这个库，能够把数据库所有的操作sql全部管理起来，一方面可以让我们清晰的看到sql项目的发展历程，另一方面本地测试后发布更加安全！

