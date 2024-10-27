README

## 简介

该系统主要用于为各种需要管理的小型平台提供框架，其包含 platform 前台设计与服务端搭建

## 开发

### 快速开始

```bash
# 克隆项目
$ git clone git@github.com:Lycoiref/Platform.git

# 安装rush (若已安装，请跳过此步骤)
$ npm install -g @microsoft/rush

# 进入项目目录
$ cd Platform

# 安装依赖
$ rush update

# 启动调试
$ cd packages/platform-core
$ rushx dev

# 启动服务端调试
$ rush dev:server
```

ps: 更多 rush 命令请参考 [rush 文档](https://rushjs.io/zh-cn/pages/intro/welcome/)

## 技术栈

-   Next
-   Golang , Java , Rust

# 架构设计

Monorepo 架构

WIP...

### Git Commit 规范

feat： 新增 feature

fix: 修复 bug

docs: 仅仅修改了文档，比如 README, CHANGELOG, CONTRIBUTE 等等

style: 仅仅修改了空格、格式缩进、逗号等等，不改变代码逻辑

refactor: 代码重构，没有加新功能或者修复 bug

perf: 优化相关，比如提升性能、体验

test: 测试用例，包括单元测试、集成测试等

chore: 改变构建流程、或者增加依赖库、工具等

revert: 回滚到上一个版本
