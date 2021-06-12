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

async function fanoutConsumer(): Promise<void> {

    // 创建链接对象
    const connection = await amqp.connect(host,'heartbeat=60');

    // 获取通道
    const channel = await connection.createChannel();

    // 声明参数
    const exchangeName = 'fanout_exchange_name';
    const queueName = 'fanout_queue';
    const routingKey = '';

    // 声明一个交换机
    await channel.assertExchange(exchangeName, 'fanout', { durable: true });

    // 声明一个队列
    await channel.assertQueue(queueName);

    // 绑定关系（队列、交换机、路由键）
    await channel.bindQueue(queueName, exchangeName, routingKey);

    // 消费
    await channel.consume(queueName, msg => {
        console.log('Consumer 1 fanout：', msg.content.toString());
        channel.ack(msg);
    });

    await channel.assertQueue('q2');

    // 绑定关系（队列、交换机、路由键）
    await channel.bindQueue('q2', exchangeName, routingKey);

    // 消费
    await channel.consume('q2', msg => {
        console.log('Consumer 2 fanout：', msg.content.toString());
        channel.ack(msg);
    });

    console.log('消费端启动成功！');
}

async function DLXConsumer(): Promise<void> {

    // 创建链接对象
    const connection = await amqp.connect(host,'heartbeat=60');

    // 获取通道
    const channel = await connection.createChannel();

    const testExchange = 'testEx';
    const testQueueDLX = 'testQueueDLX';
    const testExchangeDLX = 'testExDLX';
    const testRoutingKeyDLX = 'testRoutingKeyDLX';

    await channel.assertExchange(testExchangeDLX, 'direct', { durable: true });
    const queueResult = await channel.assertQueue(testQueueDLX, {
        exclusive: false,
    });

    await channel.bindQueue(queueResult.queue, testExchangeDLX, testRoutingKeyDLX);
    await channel.consume(queueResult.queue, msg => {
        console.log('consumer msg：dlx ', msg.content.toString());
    }, { noAck: true });

    await channel.get
}

// directConsumer();
// topicConsumer();
// fanoutConsumer();
DLXConsumer();
