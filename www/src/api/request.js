import axios from "axios";

const instance = axios.create({
  baseURL: "http://localhost:9600/api/"
});

instance.interceptors.response.use(res => {
  return Promise.resolve(res.data);
});

export default instance;
