## 1. 如何入手选择开源项目并贡献代码？

关于如何入手开源项目，可以从 [opensource.guide](https://opensource.guide/how-to-contribute/) 中找到最详尽的解答，这里简要概括一下所涉及的重点，并结合自己的理解陈述一下如何找到贡献目标。

### 1.1 如何找到合适的项目

首先从自己项目中所涉及的开源轮子以及自己准备使用的轮子中找。关注一下项目的 README 中是否有待完成事项，关注一下 issue 列表中别人反馈的 bug，尝试看能否修复个别简单 bug，就像新入职一家公司，很多时候也会从修简单的 bug 作为上手项目的方法。除此之外，关注项目的文档是否足够完善，如果不够发现有缺陷（甚至是一些错别字），都是可以方便上手修改的。如果是国外的项目，看看是否有中文版文档，如果没有可尝试将其 README 翻译出一个中文版本。

如果确实没有目标，那么可以根据自己擅长的编程语言，从这里找找看有没有可以贡献的项目

- [GitHub Explore](https://github.com/explore/)
- [Open Source Friday](https://opensourcefriday.com/)
- [First Timers Only](https://www.firsttimersonly.com/)
- [CodeTriage](https://www.codetriage.com/)
- [24 Pull Requests](https://24pullrequests.com/)
- [Up For Grabs](https://up-for-grabs.net/)
- [Contributor-ninja](https://contributor.ninja/)
- [First Contributions](https://firstcontributions.github.io/)
- [SourceSort](https://web.archive.org/web/20201111233803/https://www.sourcesort.com/)

在选择一个开源项目之前，还有一些注意事项：

- 看看这个项目活跃程度怎么样？作者是否经常维护，如果好几个月没更新过的，就没必要浪费精力了。
- 看看这个项目的贡献者分布情况，如果绝大部分贡献都只有作者本人，那么这个项目不适合参与。
- 看看 issue 列表和 pull request 列表，看别人反馈的问题，以及作者对问题的态度是否积极。
- 看看项目的 start 数。

以上就是关于关于如何选择贡献目标的简单介绍，更详细的内容可以  [opensource.guide](https://opensource.guide/how-to-contribute/) 中仔细阅读。

## 2. MimeType 项目贡献过程

###  2.1 为什么选择 mimetype

第一次开源贡献，作者根据自身情况选择使用 Go 语言并且明确此次贡献目的在于练手，基于上述原则和网站所列举项目中选择了 [mimetype](https://github.com/gabriel-vasile/mimetype) .

**什么是 mimetype **? 

> **媒体类型**（通常称为 **Multipurpose Internet Mail Extensions** 或 **MIME** 类型 ）是一种标准，用来表示文档、文件或字节流的性质和格式。它在[IETF RFC 6838](https://tools.ietf.org/html/rfc6838)中进行了定义和标准化。
>
> 互联网号码分配机构（[IANA](https://www.iana.org/)）是负责跟踪所有官方MIME类型的官方机构，您可以在[媒体类型](https://www.iana.org/assignments/media-types/media-types.xhtml)页面中找到最新的完整列表。

**[mimetype](https://github.com/gabriel-vasile/mimetype) 项目的工作原理是什么？**

通过建立 magic number 与文件类型名称的映射关系，来判断文件的类型（比如 mp4, png, jpeg 等 ）。

什么是文件的 magic number？某些类型文件会在开头前几位使用固定的字符，以此来表明自己的身份，那么对于这些文件，只要读取文件前几位，就能知道它是什么类型的文件了。但是并非所有类型的文件都有 magic number，这不是一个 100% 可靠的失败文件方式，因此首先要先确保目标文件类型支持 magic number。

基于以上原因，只要找到 README 中尚未被收录且支持 magic number 的文件类型，就可以为该开源项目贡献代码了。

### 2.2 我如何为minetype 贡献

首先可以通过各种渠道找到支持 magic number 的文件列表，对比 README，确定还没被收录的文件类型。这里参考了 [这个列表](https://gist.github.com/leommoore/f9e57ba2aa4bf197ebc5)，确定了 **xpm**。

其次就是 fork 项目，贡献代码，然后提交 pull request。至于这部分如何操作，可以参考[官方指引](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request-from-a-fork)。

这里本人的代码存在一些缺陷，被项目作者犀利指出。经过一点细节修改和讨论，再次提交 commit。

![](https://tva1.sinaimg.cn/large/008i3skNgy1gqkcvbkhnxj31780u0jzw.jpg)

最终如图中所见，代码被合并到主干，即完成了本次开源代码贡献。

### 3. 总结

对于开源项目参与，可先从最简单的入手，最重要的跨出第一步。

### 4. 参考资料

[How to Contribute to Open Source | Open Source Guides](https://www.notion.so/How-to-Contribute-to-Open-Source-Open-Source-Guides-62e7bd9c70344756b6a3b28123493ce8)

[List of file signatures - Wikipedia](https://www.notion.so/List-of-file-signatures-Wikipedia-94ac36b4311d4107b62b8d98cdc774ef)

[Detect file mime type using magic numbers and JavaScript | by Andreas Kihlberg | The everyday developer | Medium](https://www.notion.so/Detect-file-mime-type-using-magic-numbers-and-JavaScript-by-Andreas-Kihlberg-The-everyday-develo-c41061ae14194634a7d2e37fe84450c3)

[MIME 类型 - HTTP | MDN](https://www.notion.so/MIME-HTTP-MDN-d39151e9d92c42d5bf3cef61c4440603)

[zpage.mime.types.title](https://www.notion.so/zpage-mime-types-title-57b22827138f4d7ba1e0b28af455f225)

[File Magic Numbers](https://www.notion.so/File-Magic-Numbers-375dd1d1145d479a91736a89b5389fff)

[File Converter - video converter, audio converter, image converter, eBook converter](https://www.notion.so/File-Converter-video-converter-audio-converter-image-converter-eBook-converter-2f648b374471433dba613f2bac6911d8)

[gabriel-vasile/mimetype: A fast golang library for MIME type and file extension detection, based on magic numbers](https://www.notion.so/gabriel-vasile-mimetype-A-fast-golang-library-for-MIME-type-and-file-extension-detection-based-on--5fb5775ea6c446aa97b9814fead66dc0)





