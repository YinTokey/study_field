import { io } from "socket.io-client";

// 调试时，这里替换成相应的地址
const socket = io("ws://localhost:8080/", {});

socket.on("connect", () => {
    console.log(`connect ${socket.id}`);
});

socket.on("disconnect", () => {
    console.log(`disconnect`);
});

socket.on("message", (msg) => {
    console.log('收到服务端的消息： --- ' + msg);
});

setInterval(() => {
    const start = Date.now();
    socket.emit("ping", () => {
        console.log(`pong (latency: ${Date.now() - start} ms)`);
    });

    socket.emit("message", "xxx");

}, 1000);
