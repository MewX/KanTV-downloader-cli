# 看TV 命令行版 视频（电影/电视剧）下载器

[![Build Status](https://travis-ci.com/MewX/KanTV-downloader-cli.svg?branch=master)](https://travis-ci.com/MewX/KanTV-downloader-cli)

本版本封装的是 看TV Android 客户端的API。

## 使用方式

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
```

如果没有网址，你也可以通过指定 TVID 来下载视频。TVID可以通过搜索功能获取，也可以通过网址提取。

命令如下：

```
$ kantv download --tvid <TVID>

# 例子：
$ kantv download --tvid 302002655075001
```