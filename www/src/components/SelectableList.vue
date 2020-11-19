<template>
  <div class="selectable-list">
    <div
      :class="['selectable-list__item', selectedKey === item.key && 'is-selected']"
      v-for="item in list"
      :key="item.key"
      @click="handleSelect(item)"
    >
      {{ item.name }}
    </div>
  </div>
</template>

<script>
export default {
  name: "SelectableList",
  props: {
    list: {
      type: Array,
      default: () => [],
    },
    selectedKey: {
      type: String,
      default: "",
    },
  },
  methods: {
    handleSelect(item) {
      this.$emit("update:selectedKey", item.key);
    },
  },
};
</script>

<style scoped>
.selectable-list__item {
  padding: 10px 20px;
  border-radius: 2px;
  cursor: pointer;
  transition: all 0.3s;
}
.selectable-list__item.is-selected {
  background-color: var(--active-color-bg) !important;
  color: var(--active-color) !important;
}
.selectable-list__item + .selectable-list__item {
  margin-top: 10px;
}
.selectable-list__item:hover {
  background-color: var(--bg-color-0);
}
</style>
