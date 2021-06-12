import * as amqp from 'amqplib';

const host = 'amqp://localhost';


async function directConsumer(): Promise<void> {

    // 创建链接对象
    const connection = await amqp.connect(host,'heartbeat=60');

    // 获取通道
    const channel = await connection.createChannel();

    // 声明参数
    const exchangeName = 'direct_exchange_name';
    const queueName = 'direct_queue';
    const routingKey = 'X1';

    // 声明一个交换机
    await channel.assertExchange(exchangeName, 'direct', { durable: true });

    // 声明一个队列
    await channel.assertQueue(queueName);

    // 绑定关系（队列、交换机、路由键）
    await channel.bindQueue(queueName, exchangeName, routingKey);

    // 消费
    await channel.consume(queueName, msg => {
        console.log('Consumer：', msg.content.toString());
        channel.ack(msg);
    });

    console.log('消费端启动成功！');
}

async function topicConsumer() :Promise<void> {

    // 创建链接对象
    const connection = await amqp.connect(host,'heartbeat=60');

    // 获取通道
    const channel = await connection.createChannel();

    // 声明参数
    const exchangeName = 'topic_exchange_name';
    const queueName = 'topic_queue';
    const routingKey = 'topic_routingKey.#';

    // 声明一个交换机
    await channel.assertExchange(exchangeName, 'topic', { durable: true });

    // 声明一个队列
    await channel.assertQueue(queueName);

    // 绑定关系（队列、交换机、路由键）
    await channel.bindQueue(queueName, exchangeName, routingKey);

    // 消费
    await channel.consume(queueName, msg => {
        console.log('Consumer：', msg.content.toString());
        channel.ack(msg);
    });

    console.log('消费端启动成功！');
}

// directConsumer();
topicConsumer();
