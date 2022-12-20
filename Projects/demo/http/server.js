const http = require('http');

const server = http.createServer(function (req, res) {
    res.writeHead(200, {
        "Content-Type": "text/html;charset=UTF-8"
    })
    res.end("欢迎来到推啊！！！")
})

server.listen(3000, function () {
    console.log('listening port 3000')
})
