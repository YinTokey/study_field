import { Application } from 'egg';

export default (app: Application) => {
    const { controller, router } = app;

    router.get('/', controller.home.index);
    router.get('/list', controller.list.index);
    router.get('/random', controller.random.index);
    router.get('/tags', controller.tag.index);

};
