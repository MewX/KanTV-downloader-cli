# 看TV 命令行版 视频（电影/电视剧）下载器

[![Build Status](https://travis-ci.com/MewX/KanTV-downloader-cli.svg?branch=master)](https://travis-ci.com/MewX/KanTV-downloader-cli)

本版本封装的是 看TV Android 客户端的API。

## 编译方式

请事先下载好最新版 Bazel 编译套件：
https://docs.bazel.build/versions/master/install.html

然后执行：

```
$ bazel build :kantv
```

## 使用方式

如果不想编译，请直接到 [Release区](https://github.com/MewX/KanTV-downloader-cli/releases) 下载最新版使用。

查看完整用法的命令：

```
$ kantv --help
```

### 获取支持的国家列表

```
$ kantv country
```

### 下载视频

指定网址下载视频：

```
$ kantv download --url <video_URL>

# 例子：
$ kantv download --url https://www.wekan.tv/movie/302002655075001
$ bazel run :kantv -- download --url https://www.wekan.tv/tvdrama/301930368997001-161930368997002 --outdir ~/downloads
```

指定 `--outdir ...` 可以自定义下载位置。

如果没有网址，你也可以通过指定 TVID (+Part ID) 来下载视频。TVID可以通过搜索功能获取，也可以通过网址提取。

命令如下：

```
$ kantv download --tvid <TVID>

# 例子：
$ kantv download --tvid 302002655075001
```

或者：

```
$ kantv download --tvid <TVID> --partid <Part ID>

# 例子：
$ kantv download --tvid 301930368997001 --partid 161930368997002 --outdir ~/Downloads
```

目前版本视频没有合并，有的播放器可以通过m3u8播放。 下一个版本将合并成一个ts文件。
