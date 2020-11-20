<template>
  <div class="home">
    <div class="etcd-box">
      <div class="etcd-host">
        <EndpointList :selectedKey.sync="connectAddr" :tokens="tokens"></EndpointList>
        <!-- <SelectableList :list="addrList" :selectedKey.sync="connectAddr"></SelectableList> -->
      </div>
      <EtcdEdit v-if="connected" :connectAddr="connectAddr"></EtcdEdit>
      <div v-else class="tobe-connected">
        <div v-if="needConnect">
          <h2>{{ connectAddr }}</h2>
          <p>填写当前连接地址用户信息然后点击”连接“</p>
          <div class="etcd-auth">
            <a-input placeholder="用户名，无可填空" v-model="user.username"></a-input>
            <a-input placeholder="密码，无可填空" type="password" v-model="user.password"></a-input>
            <a-button :loading="connectting" type="primary" @click="readyToConnect">连接</a-button>
            <a-alert v-if="connectError" style="margin-top: 10px" type="error" :message="connectError" banner />
          </div>
        </div>
        <p v-else>👈 选择一个地址连接管理Etcd配置</p>
      </div>
    </div>
  </div>
</template>

<script>
import EtcdEdit from "./../components/EtcdEdit";
import EndpointList from "./../components/EndpointList";
import { getTokens, setTokens } from "./../util/storage.js";
import { testConnect } from "./../api/auth.js";

export default {
  name: "Home",
  components: {
    EtcdEdit,
    EndpointList,
  },
  data() {
    return {
      connectAddr: "",
      tokens: getTokens(),
      connected: false,
      connectting: false,
      connectError: "",
      user: {
        username: "",
        password: "",
      },
    };
  },
  computed: {
    needConnect() {
      return !this.connected && this.connectAddr;
    },
  },
  watch: {
    connectAddr: {
      immediate: true,
      handler() {
        this.checkConnectStatus();
      },
    },
  },
  methods: {
    checkConnectStatus() {
      this.connected = Object.hasOwnProperty.call(this.tokens, this.connectAddr);
    },
    async readyToConnect() {
      this.connectting = true;
      this.connectError = "";
      try {
        setTokens(this.connectAddr, this.user);
        await testConnect(this.connectAddr);
        this.connected = true;
        this.updateTokens();
      } catch (error) {
        console.error(error);
        setTokens(this.connectAddr, undefined);
        this.connectError = error;
      } finally {
        this.connectting = false;
      }
    },
    updateTokens() {
      this.tokens = getTokens();
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
  padding: 10px 0;
  margin-right: 20px;
  width: 200px;
}
.tobe-connected {
  height: calc(100vh - 120px);
  width: 100%;
  padding: 20px;
  background-color: var(--bg-color-1);
  border-radius: 4px;
  box-shadow: 0 0 15px 1px #31354b;
}
.tobe-connected h2 {
  color: var(--color-1);
}
.etcd-auth {
  width: 400px;
}
.etcd-auth .ant-input {
  margin-bottom: 20px;
}
.ant-select >>> .ant-select-selection {
  background-color: rgba(255, 255, 255, 0.7) !important;
}
.ant-select >>> .ant-select-selection__placeholder {
  color: #7e7e7e;
}
</style>
