const http = require('http');
const axios = require('axios')
// const server = http.createServer(function (req, res) {
//
//     res.writeHead(200, {
//         "Content-Type": "text/html;charset=UTF-8"
//     })
//     res.end("hello client!")
// })
//
// server.listen(3000, function () {
//     console.log('listening port 3000')
// })

async function get(url) {
    const content = await axios.get(url)
    console.log(content)
}

const url = 'https://event.on24.com/eventRegistration/console/EventConsoleApollo.jsp?simulive=y&eventid=2077058&sessionid=1&username=&partnerref=&format=fhvideo1&mobile=&flashsupportedmobiledevice=&helpcenter=&key=5203927B9EC218A69D0B52B5EF25EED3&newConsole=true&nxChe=true&newTabCon=true&consoleEarEventConsole=false&text_language_id=en&playerwidth=748&playerheight=526&eventuserid=464633638&contenttype=A&mediametricsessionid=400881544&mediametricid=2929477&usercd=464633638&mode=launch'
get(url).then()





