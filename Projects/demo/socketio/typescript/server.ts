import { Server } from "socket.io";

const io = new Server(8080);

io.on("connect", (socket) => {
    console.log(`connect ${socket.id}`);

    socket.on("ping", (cb) => {
        console.log("ping");
        cb();
    });

    socket.on("event" ,(cb) => {
        console.log("event");
        cb();
    });

    socket.on("message", function (msg) {
        io.emit("message", `我是服务端，重复 : ${msg}`) //将新消息广播出去
    });

    socket.on("disconnect", () => {
        console.log(`disconnect ${socket.id}`);
    });
});
