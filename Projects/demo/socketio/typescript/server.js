"use strict";
exports.__esModule = true;
var socket_io_1 = require("socket.io");
var io = new socket_io_1.Server(8080);
io.on("connect", function (socket) {
    console.log("connect " + socket.id);
    socket.on("ping", function (cb) {
        console.log("ping");
        cb();
    });
    socket.on("disconnect", function () {
        console.log("disconnect " + socket.id);
    });
});
