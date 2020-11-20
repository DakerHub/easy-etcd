<template>
  <a-config-provider :locale="zh_CN">
    <a-layout id="app">
      <a-layout-header class="layout-header">
        <h1>Easy Etcd</h1>
        <!-- <a-button type="link" ghost>登出</a-button> -->
      </a-layout-header>
      <a-layout-content>
        <router-view></router-view>
      </a-layout-content>
      <!-- <a-layout-footer style="text-align: center">
        Easy Etcd by
        <a href="https://github.com/DakerHub" target="__blank" ref="noopener">Daker</a>
      </a-layout-footer> -->
    </a-layout>
  </a-config-provider>
</template>

<script>
import zh_CN from "ant-design-vue/lib/locale-provider/zh_CN";

export default {
  name: "app",
  data() {
    return {
      zh_CN,
    };
  },
};
</script>

<style scoped>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  min-height: 100vh;
}

.ant-layout,
.ant-layout-footer {
  background-color: var(--bg-color-2);
  color: var(--color-1);
}
.ant-layout-header {
  background-color: var(--bg-color-1);
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.ant-layout-header h1 {
  color: var(--color-1);
  margin: 0;
}
.ant-layout-footer {
  height: 40px;
  padding: 0;
  line-height: 40px;
}
</style>

<style >
.ant-input {
  background-color: transparent;
  color: var(--color-1);
}
</style>
