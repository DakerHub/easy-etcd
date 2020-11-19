import Vue from "vue";
import Antd, { message } from "ant-design-vue";
import "ant-design-vue/dist/antd.css";
Vue.use(Antd);

Vue.prototype.$alertSuccess = content => {
  message.success(content);
};

Vue.prototype.$alertError = content => {
  message.error(content);
};
