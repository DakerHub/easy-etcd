<template>
  <div class="home">
    <div class="etcd-box">
      <div class="etcd-host">
        <SelectableList :list="addrList" :selectedKey.sync="connectAddr"></SelectableList>
      </div>
      <EtcdEdit v-if="connected" :connectAddr="connectAddr"></EtcdEdit>
      <div v-else class="no-addr">
        <div v-if="needConnect">
          <p>填写当前连接地址用户信息然后点击”连接“</p>
          <div class="etcd-auth">
            <a-input placeholder="用户名，无可填空"></a-input>
            <a-input placeholder="密码，无可填空" type="password"></a-input>
            <a-button type="primary">连接</a-button>
          </div>
        </div>
        <p v-else>👆 选择一个地址连接管理Etcd配置</p>
      </div>
    </div>
  </div>
</template>

<script>
import EtcdEdit from "./../components/EtcdEdit";
import SelectableList from "./../components/SelectableList";

export default {
  name: "Home",
  components: {
    EtcdEdit,
    SelectableList,
  },
  data() {
    return {
      addrList: [
        {
          name: "开发环境",
          key: "http://123.207.16.31:2379",
        },
        {
          name: "测试环境",
          key: "http://localhost:2380",
        },
      ],
      connectAddr: "http://123.207.16.31:2379",
      connected: true,
    };
  },
  computed: {
    needConnect() {
      return !this.connected && this.connectAddr;
    },
  },
  watch: {
    connectAddr() {
      this.checkConnectStatus();
    },
  },
  methods: {
    checkConnectStatus() {
      this.connected = false;
    },
  },
};
</script>

<style scoped>
.etcd-box {
  display: flex;
  max-width: 1200px;
  margin: 20px auto;
}
.etcd-edit {
  width: 100%;
}
.home {
  height: 100%;
}
.etcd-host {
  padding: 10px 20px;
}
.no-addr {
  height: calc(100vh - 200px);
  padding: 20px;
  background-color: var(--bg-color-1);
  border-radius: 4px;
  box-shadow: 0 0 15px 1px #31354b;
}
.etcd-auth {
  width: 400px;
}
.etcd-auth .ant-input {
  margin-bottom: 20px;
}
.selectable-list {
  width: 200px;
}
.ant-select >>> .ant-select-selection {
  background-color: rgba(255, 255, 255, 0.7) !important;
}
.ant-select >>> .ant-select-selection__placeholder {
  color: #7e7e7e;
}
</style>
