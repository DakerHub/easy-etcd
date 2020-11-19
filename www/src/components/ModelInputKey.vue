<template>
  <div>
    <slot :open="showModal"></slot>
    <a-modal v-model="visible" :title="title" @ok="handleOk">
      <a-input v-model="value" autoFocus :placeholder="placeholder" @pressEnter="handleOk"></a-input>
    </a-modal>
  </div>
</template>

<script>
export default {
  name: "ModelInputKey",
  props: {
    title: {
      type: String,
      default: "",
    },
    placeholder: {
      type: String,
      default: "",
    },
  },
  data() {
    return {
      value: "/test/xxxx",
      visible: false,
    };
  },
  methods: {
    showModal() {
      this.visible = true;
    },
    handleOk() {
      this.$emit("change", this.value);
      this.value = "";
      this.visible = false;
    },
  },
};
</script>

<style  scoped>
</style>
