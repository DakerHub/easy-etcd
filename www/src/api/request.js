import axios from "axios";
import { getTokens } from "./../util/storage";
import { message } from "ant-design-vue";
import router from "./../router";

const isDev = process.env.NODE_ENV === "development";

const instance = axios.create({
  baseURL: isDev ? "http://localhost:9600/api/" : "/api/"
});

instance.interceptors.request.use(config => {
  const tokens = getTokens();
  const { nodes = [] } = config.data;
  const node = nodes[0];

  if (node) {
    const auth = tokens[node];
    if (!auth) {
      return Promise.reject("No Auth");
    }
    config.auth = auth;
  }

  return config;
});

instance.interceptors.response.use(
  res => {
    return Promise.resolve(res.data);
  },
  err => {
    if (err.response.status === 401) {
      router.push("/login");
    }
    const msg = err.response?.data?.message || "未知错误";
    message.error(msg);
    return Promise.reject(msg);
  }
);

export default instance;
