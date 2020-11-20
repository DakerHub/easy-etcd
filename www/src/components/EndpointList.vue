<template>
  <div>
    <SelectableList
      v-if="endpoints.length"
      :list="endpointsRender"
      :selectedKey.sync="connectAddr"
      @close="removeEndpoint"
    >
      <template #title="{ item }">
        <p>
          <a-icon v-if="item.connected" type="smile" theme="filled" style="color: #8bc34a" />
          <span style="font-weight: bold"> {{ item.name }}</span>
        </p>
      </template>
    </SelectableList>
    <div v-else>
      <!-- <a-empty></a-empty> -->
      <p>还没设置etcd服务器地址</p>
    </div>
    <div class="endpoint-action">
      <ModelEtcdEndpoint title="新增Etcd连接地址" @change="handleChange">
        <template v-slot="{ open }">
          <a-button block type="primary" @click="open" icon="plus" ghost>新增Etcd连接地址</a-button>
        </template>
      </ModelEtcdEndpoint>
    </div>
  </div>
</template>

<script>
import SelectableList from "./../components/SelectableList";
import ModelEtcdEndpoint from "./ModelEtcdEndpoint";
import { setEndpoint, getEndpoint, removeEndpoint } from "./../util/storage.js";

export default {
  name: "EndpointList",
  components: {
    SelectableList,
    ModelEtcdEndpoint,
  },
  props: {
    selectedKey: {
      type: String,
      default: "",
    },
    tokens: {
      type: Object,
      default: () => ({}),
    },
  },
  data() {
    return {
      endpoints: getEndpoint(),
      connectAddr: "",
    };
  },
  watch: {
    connectAddr(value) {
      this.$emit("update:selectedKey", value);
    },
    selectedKey: {
      immediate: true,
      handler() {
        this.connectAddr = this.selectedKey;
      },
    },
  },
  computed: {
    endpointsRender() {
      return this.endpoints.map((e) => ({
        name: e.name,
        desc: e.url,
        key: e.url,
        connected: !!this.tokens[e.url],
      }));
    },
  },
  methods: {
    handleChange(endpoint) {
      setEndpoint(endpoint);
      this.updateEndpoints();
    },
    updateEndpoints() {
      this.endpoints = getEndpoint();
    },
    removeEndpoint(p) {
      removeEndpoint(p.key);
      this.updateEndpoints();
      this.connectAddr = "";
    },
  },
};
</script>

<style scoped>
.endpoint-action {
  margin-top: 10px;
}
</style>
