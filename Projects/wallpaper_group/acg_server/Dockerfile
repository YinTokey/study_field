
# 设置基本镜像。 稍后添加的所有内容都将基于该镜像。 在此示例中，我们使用的是Node.js版本 10.0-alpine
FROM node:latest

# 环境设置
ARG NODE_ENV=development
ENV NODE_ENV=${NODE_ENV}

# 运行终端命令，创建文件夹
RUN mkdir -p /usr/src/app

# 此命令设置将用于COPY，RUN和CMD命令的工作目录（docker 容器）。
WORKDIR /usr/src/app

# 拷贝项目的包依赖描述 到工作目录中（docker 容器）
COPY package.json /usr/src/app/

# 在工作目录中安装项目所需要的依赖
RUN npm i --save

# RUN npm i --production --registry=https://registry.npm.taobao.org

# 把工程的文件拷贝到容器中
COPY . /usr/src/app

# 端口
EXPOSE 7002

# 在容器中运行程序
CMD npm run dev
