/// <reference types="node" />
/// <reference types="mongoose" />
import { EggAppInfo, PowerPartial } from 'egg';
declare const _default: (appInfo: EggAppInfo) => {
    sourceUrl: string;
    multipart?: PowerPartial<{
        mode?: string;
        fileModeMatch?: string | RegExp | ((ctx: import("egg").Context<any>) => boolean) | import("egg").MatchItem[];
        autoFields?: boolean;
        defaultCharset?: string;
        fieldNameSize?: number;
        fieldSize?: string | number;
        fields?: number;
        fileSize?: string | number;
        files?: number;
        whitelist?: string[] | ((filename: string) => boolean);
        fileExtensions?: string[];
        tmpdir?: string;
        cleanSchedule?: import("egg-multipart").ScheduleOptions;
    }>;
    view?: PowerPartial<{
        root: string;
        cache: boolean;
        defaultExtension: string;
        defaultViewEngine: string;
        mapping: import("egg").PlainObject<string>;
    }>;
    workerStartTimeout?: number;
    baseDir?: string;
    middleware?: string[];
    bodyParser?: PowerPartial<{
        enable: boolean;
        encoding: string;
        formLimit: string;
        jsonLimit: string;
        textLimit: string;
        strict: boolean;
        queryString: {
            arrayLimit: number;
            depth: number;
            parameterLimit: number;
        };
        ignore: import("egg").IgnoreOrMatch;
        match: import("egg").IgnoreOrMatch;
        enableTypes: string[];
        extendTypes: {
            json: string[];
            form: string[];
            text: string[];
        };
    }>;
    logger?: PowerPartial<import("egg").EggLoggerConfig>;
    customLogger?: PowerPartial<{
        [key: string]: import("egg-logger").EggLoggerOptions;
    }>;
    httpclient?: PowerPartial<import("egg").HttpClientConfig>;
    development?: PowerPartial<{
        watchDirs: string[];
        ignoreDirs: string[];
        fastReady: boolean;
        reloadOnDebug: boolean;
        overrideDefault: boolean;
        overrideIgnore: boolean;
        reloadPattern: string | string[];
    }>;
    customLoader?: PowerPartial<{
        [key: string]: import("egg").CustomLoaderConfig;
    }>;
    dump?: PowerPartial<{
        ignore: Set<string>;
    }>;
    env?: string;
    HOME?: string;
    hostHeaders?: string;
    i18n?: PowerPartial<{
        defaultLocale: string;
        dir: string;
        queryField: string;
        cookieField: string;
        cookieMaxAge: string | number;
    }>;
    ipHeaders?: string;
    jsonp?: PowerPartial<{
        limit: number;
        callback: string;
        csrf: boolean;
        whiteList: import("express-serve-static-core").PathParams;
    }>;
    keys?: string;
    name?: string;
    pkg?: any;
    rundir?: string;
    security?: PowerPartial<{
        domainWhiteList: string[];
        protocolWhiteList: string[];
        defaultMiddleware: string;
        csrf: any;
        xframe: {
            enable: boolean;
            value: "SAMEORIGIN" | "DENY" | "ALLOW-FROM";
        };
        hsts: any;
        methodnoallow: {
            enable: boolean;
        };
        noopen: {
            enable: boolean;
        };
        xssProtection: any;
        csp: any;
    }>;
    siteFile?: PowerPartial<import("egg").PlainObject<string | Buffer>>;
    watcher?: PowerPartial<import("egg").PlainObject<any>>;
    onClientError?: PowerPartial<(err: Error, socket: import("net").Socket, app: import("egg").EggApplication) => import("egg").ClientErrorResponse | Promise<import("egg").ClientErrorResponse>>;
    serverTimeout?: number;
    mongoose?: PowerPartial<{
        url?: string;
        options?: import("mongoose").ConnectOptions;
        client?: import("egg").MongooseConfig;
        clients?: {
            [key: string]: import("egg").MongooseConfig;
        };
    }>;
    redis?: PowerPartial<import("egg-redis").EggRedisOptions>;
};
export default _default;
