import axios from 'axios';
// 创建axios实例
export const axiosRequest = axios.create({
    baseURL: '', // api 的 base_url
    timeout: 20000 // 请求超时时间
});

// request拦截器
axiosRequest.interceptors.request.use(
    config => {

        /*
         *  config.headers.cookie = '__cfduid=d79de951597bfc71c046ceb79cb78b6891608887954; XSRF-TOKEN=eyJpdiI6InJVRVRaV3U5a2RYQWdNcW1xNUhpXC93PT0iLCJ2YWx1ZSI6IlB4N0xLTnhqcHJDRnZ2TElsOGRQWlBNYTE0RExqUHR0SnJoMWxGMjNocU9IMnJscUc3N3I3bEIwVjFNRVhvUFoiLCJtYWMiOiJhZGIwZTc5MTNhYWEzODQ1NDA1MTNkMTk5NWQzYWVhMWY1YjM5ZGFiZTRmZDkzODNiZGI5MDQ2OWY1OThkY2YzIn0%3D; wallhaven_session=eyJpdiI6ImRhcHRkN1FaNlJFOEtSNTVwYVBNZ0E9PSIsInZhbHVlIjoic3d0Sm53QkxQNWJJQXI0WWo2MVh2TDhqZVRzSjVnYkl1bjFzcTVKaU52b3dacjlaXC8zdDc0bGVNWXo2dnhyMGUiLCJtYWMiOiJkNjA5ZjlmY2FjOGY3ZjcwYzEyODRjM2IwMWRjMzlmOTc5MGJiODI2MTNiYzUyZDkwM2VmN2I2YjgzNzg4MWJmIn0%3D'
         *  config.headers['user-agent'] = 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36'
         */

        return config;
    },
    error => {
        // Do something with request error
        console.log(error); // for debug
        Promise.reject(error);
    }
);

// 响应拦截器
axiosRequest.interceptors.response.use(
    response => {
        return response.data;
    },
    error => {
        return Promise.reject(error);
    }
);

