## Install Ceph
Use cephadm to manage ceph. Follow the installation document from [cephadm](https://docs.ceph.com/docs/master/cephadm/install/)

Note that cephadm is running in docker.

## 前后端交互
使用go编写后端：
- AJAX + gin实现前后端交互
- go-ceph 实现与ceph交互

[GOlang 实现MP4视频文件服务器](https://blog.csdn.net/wangshubo1989/article/details/78053856)
[用Golang搭建网站](https://studygolang.com/articles/20362?fr=sidebar)


## ceph 部署
修改hosts文件为
```bash
192.168.92.128 admin
192.168.92.129 node0
192.168.92.130 node1
```
同时需要保证每个机器的hostname与这里设置的一致，例如在node0可以通过执行
```bash
hostnamectl set-hostname node0
```
做到这一点

将admin key发送到所有节点之后，需要对key文件添加权限才能正常使用
```bash
sudo chmod +r /etc/ceph/ceph.client.admin.keyring
```

获取数据
```bash
rados -p <pool> <object> <file path>
```