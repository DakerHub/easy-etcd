import axios from "axios";
import { getTokens } from "./../util/storage";
import { message } from "ant-design-vue";

const instance = axios.create({
  baseURL: "http://localhost:9600/api/"
});

instance.interceptors.request.use(config => {
  const tokens = getTokens();
  const { nodes } = config.data;
  const node = nodes[0];

  const auth = tokens[node];

  if (!auth) {
    return Promise.reject("No Auth");
  }

  config.auth = auth;

  return config;
});

instance.interceptors.response.use(
  res => {
    return Promise.resolve(res.data);
  },
  err => {
    message.error("服务器异常");
    const msg = err.response?.data?.message || "未知错误";
    return Promise.reject(msg);
  }
);

export default instance;
