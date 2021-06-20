import {EventEmitter} from 'events';

const timerEventEmitter = new EventEmitter();

let currentTime = 0;

setInterval(()=> {
    currentTime++;
    timerEventEmitter.emit('update',currentTime);
}, 1000);

// timerEventEmitter.on('update',time=>{
//     console.log('从发布者收到的消息：');
//     console.log(`程序已经运行了 ${time} 秒`);
// });

timerEventEmitter.once('update', (time) => {
    console.log('从发布者收到的消息：');
    console.log(`程序已经运行了 ${time} 秒`);
});

const arr = [1,2,3];
function sum(x,y,z) {
    return x+y+z;
}

