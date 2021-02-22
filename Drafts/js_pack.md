# javascript 打包工具汇总
## 概述
前端打包综合来看，最建议的选择还是webpack。
后端则根据情况，选择webpack, pkg, ncc。

## pkg

把 node 打包成各种平台的二进制可执行文件。

## ncc

是pkg的作者又开的项目，把node.js 工程打包成仅一个 js 文件。

## webpack

SPA，目前最强大的打包工具，但是过于臃肿，如何单纯打包js不推荐。

在webpack里一切皆模块（包括图片文件等）, 通过 loader 转换文件，通过Plugin 注入钩子，最后输出由多个模块组合成的文件。Webpack 专注于构建模块化项目。就像他官网所介绍。

![](https://webpack.github.io/assets/what-is-webpack.png)

它可以做各种事，比如

- [打包发布到Npm上的库](http://webpack.wuhaolin.cn/3实战/3-13构建Npm模块.html)
- [构建Node.js应用(同构应用)](http://webpack.wuhaolin.cn/3实战/3-11构建同构应用.html)
- [构建Electron应用](http://webpack.wuhaolin.cn/3实战/3-12构建Electron应用.html)
- [构建离线应用(ServiceWorkers)](http://webpack.wuhaolin.cn/3实战/3-14构建离线应用.html)



最简配置 https://blog.csdn.net/liuxiao723846/article/details/106717821

`一个坑` 它没法把Node.js 依赖的第三方库一起打进去，必须是 external 外部依赖形式。

 https://blog.csdn.net/huzhenv5/article/details/103925804



## Rollup

相当于轻量化 webpack，打出来的包更小。但是功能不够完善

## browserify

## 


## grunt

最老牌的打包工具，它运用配置的思想来写打包脚本，一切皆配置，所以会出现比较多的配置项，诸如option,src,dest等等。而且不同的插件可能会有自己扩展字段，认知成本高，运用的时候需要明白各种插件的配置规则

现在基本没人用。

## gulp

用代码方式来写打包脚本，并且代码采用流式的写法，只抽象出了gulp.src, gulp.pipe, gulp.dest, gulp.watch 接口，运用相当简单。更易于学习和使用，使用gulp的代码量能比grunt少一半左右。

`缺点`：和 grunt 一样要写一大堆配置文件。


## Parcel

它利用多核处理提供了极快的速度，并且不需要任何配置。

Parcel的优点：

- 极速打包。Parcel 使用 worker 进程去启用多核编译。同时有文件系统缓存，即使在重启构建后也能快速再编译。
- 开箱即用。对 JS, CSS, HTML, 文件 及更多的支持，而且不需要插件。
- 自动转换。如若有需要，Babel, PostCSS, 和PostHTML甚至 node_modules 包会被用于自动转换代码.
- 热模块替换。Parcel 无需配置，在开发环境的时候会自动在浏览器内随着你的代码更改而去更新模块。
- 友好的错误日志。当遇到错误时，Parcel 会输出 语法高亮的代码片段，帮助你定位问题。

缺点：

- 不支持SourceMap：在开发模式下，Parcel也不会输出SourceMap，目前只能去调试可读性极低的代码；
- 不支持剔除无效代码(TreeShaking)：很多时候我们只用到了库中的一个函数，结果Parcel把整个库都打包了进来；
- 一些依赖会让Parcel出错：当你的项目依赖了一些Npm上的模块时，有些Npm模块会让Parcel运行错误；
- !!!!! Parcel使用场景受限。目前Parcel只能用来构建用于运行在浏览器中的网页，这也是他的出发点和专注点


#  参考链接

https://segmentfault.com/a/1190000022695840

https://xinyufeng.net/2019/09/15/%E5%89%8D%E7%AB%AF%E6%9E%84%E5%BB%BA%E5%B7%A5%E5%85%B7%E5%8F%91%E5%B1%95%E5%8F%8A%E5%85%B6%E6%AF%94%E8%BE%83/

