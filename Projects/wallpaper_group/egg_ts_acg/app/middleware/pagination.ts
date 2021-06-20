/*
 * import {Context, Application} from 'egg';
 *
 * export default function paginationMiddleWare(app: Application, ctx:Context) :any {
 *     return async (ctx:Context, next:()=> Promise<any>) => {
 *
 *         if (!ctx.pagination) {
 *
 *             const query = ctx.query;
 *             const config = ctx.app.config;
 *             const pagination = {};
 *
 *             pagination.limit = Math.min(100, parseInt(query.limit || config.default_limit, 10));
 *             const page = Math.max(1, parseInt(query.page || config.default_page, 10));
 *             pagination.skip = (page - 1) * pagination.limit;
 *
 *             console.log('执行分页' + query + pagination.limit);
 *
 *             ctx.pagination = pagination;
 *         }
 *         await next();
 *
 *     };
 *
 *
 * }
 */