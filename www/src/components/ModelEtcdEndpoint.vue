<template>
  <div>
    <slot :open="showModal"></slot>
    <a-modal v-model="visible" :title="title" @ok="handleOk">
      <a-input v-model="form.name" autoFocus placeholder="请输入连接名称，如：我的ETCD"></a-input>
      <a-input
        v-model="form.url"
        placeholder="请输入连接地址，如：http://localhost:2379"
        @pressEnter="handleOk"
      ></a-input>
    </a-modal>
  </div>
</template>

<script>
export default {
  name: "ModelEtcdEndpoint",
  props: {
    title: {
      type: String,
      default: "",
    },
  },
  data() {
    return {
      form: {
        name: "",
        url: "",
      },
      visible: false,
    };
  },
  methods: {
    showModal() {
      this.visible = true;
    },
    handleOk() {
      if (!(this.form.name && this.form.url)) {
        return this.$alertError("请填写全部字段");
      }

      this.$emit("change", this.form);
      this.form.name = "";
      this.form.url = "";
      this.visible = false;
    },
  },
};
</script>

<style  scoped>
.ant-input {
  width: 320px;
  margin-bottom: 10px;
}
</style>
