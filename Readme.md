## 安装
```bash
sudo apt-get install libcephfs-dev librbd-dev librados-dev
make cephweb
```
## 运行
```bash
./cephweb –help
./cephweb <command> --help
```
默认情况下，需要创建名为"mytest"的ceph存储池，然后在项目路径下运行
```
./cephweb init
./cephweb start <ipaddress or host>
```
就可以启动服务器。

运行
```bash
./ceph video --help
./ceph video add --help
```
查看如何添加视频。

## ceph 部署
### ceph-deploy 安装
```bash
wget -q -O- 'https://download.ceph.com/keys/release.asc' | sudo apt-key add -
echo deb https://download.ceph.com/debian-{ceph-stable-release}/ $(lsb_release -sc) main | sudo tee /etc/apt/sources.list.d/ceph.list
sudo apt update
sudo apt install ceph-deploy
```
### 其他环境安装
```bash
sudo apt install ntpsec
sudo apt install openssh-server
```
### 网络配置
修改hosts文件为
```bash
192.168.92.128 admin
192.168.92.129 node0
192.168.92.130 node1
```
创建ceph专用用户
```bash
sudo useradd -d /home/{username} -m {username}
sudo passwd {username}
echo "{username} ALL = (root) NOPASSWD:ALL" | sudo tee /etc/sudoers.d/{username}
sudo chmod 0440 /etc/sudoers.d/{username}
```
同时需要保证每个机器的hostname与这里设置的一致，例如在node0可以通过执行
```bash
hostnamectl set-hostname node0
#xxxxx
```
配置ssh无密码访问
```bash
ssh-keygen
ssh-copy-id {username}@node0
ssh-copy-id {username}@node1
```
修改 ~/.ssh/config 从而指定各个节点username
```bash
Host node1
   Hostname node1
   User {username}
Host node2
   Hostname node2
   User {username}
Host node3
   Hostname node3
   User {username}
```
创建cluster目录
```bash
mkdir my-cluster
cd my-cluster
ceph-deploy new node0
ceph-deploy install node0 node1
ceph-deploy mon create-initial
ceph-deploy admin node0 node1
```

将admin key发送到所有节点之后，需要对key文件添加权限才能正常使用
```bash
sudo chmod +r /etc/ceph/ceph.client.admin.keyring
ceph-deploy mgr create node0
ceph-deploy osd create --data /dev/adb node0
ceph-deploy osd create --data /dev/adb node1
ceph-deploy mon add node0
ceph-deploy rgw create node0
ceph-deploy rgw create node1
```

获取数据
```bash
rados -p <存储池名称> get <对象名> <文件路径>
```
创建存储池
```bash
ceph osd pool create <存储池名称> <备份数量>
```
列出存储池中所有文件块
```bash
rados -p <存储池名称> ls
```
添加对象
```bash
rados put <对象名> <文件路径> --pool=<存储池名称>
```

删除对象
```bash
rados -p <存储池> rm <对象名> //或者加--force-full时强制删除一个对象，不在乎对象此时状态
```