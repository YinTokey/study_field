"use strict";
exports.__esModule = true;
var socket_io_client_1 = require("socket.io-client");
var socket = socket_io_client_1.io("ws://localhost:8080/", {});
socket.on("connect", function () {
    console.log("connect " + socket.id);
});
socket.on("disconnect", function () {
    console.log("disconnect");
});
setInterval(function () {
    var start = Date.now();
    socket.emit("ping", function () {
        console.log("pong (latency: " + (Date.now() - start) + " ms)");
    });
}, 1000);
