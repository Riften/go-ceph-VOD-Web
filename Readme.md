## Install Ceph
Use cephadm to manage ceph. Follow the installation document from [cephadm](https://docs.ceph.com/docs/master/cephadm/install/)

Note that cephadm is running in docker.

## 前后端交互
使用go编写后端：
- AJAX + gin实现前后端交互
- go-ceph 实现与ceph交互

[GOlang 实现MP4视频文件服务器](https://blog.csdn.net/wangshubo1989/article/details/78053856)


<!--
## 跑通ceph-video-web
### build vue
在vue-front目录
```
yarn config set registry https://registry.npm.taobao.org
yarn install
```

<b>Current existing ChromeDriver binary is unavailable</b>解决方法 <!--code>npm install chromedriver --chromedriver_cdnurl=http://cdn.npm.taobao.org/dist/chromedriver</code-->
把命令行提示信息里的chromedriver下下来，解压后，放到项目目录下node_modules/chromedriver/li
<!--把提示中下载的/tmp/2.46/chromedriver/chromedriver挪到node_modules/chromedriver-->
-->