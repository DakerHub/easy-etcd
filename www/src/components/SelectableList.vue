<template>
  <div class="selectable-list">
    <div
      :class="['selectable-list__item', selectedKey === item.key && 'is-selected']"
      v-for="item in list"
      :key="item.key"
      @click="handleSelect(item)"
    >
      <slot name="title" :item="item">
        <p style="font-weight: bold">
          {{ item.name }}
        </p>
      </slot>
      <p v-if="item.desc" class="selectable-list__desc">{{ item.desc }}</p>
      <a-button
        class="selectable-list__close"
        type="danger"
        shape="circle"
        icon="close"
        size="small"
        @click.stop="$emit('close', item)"
      ></a-button>
      <!-- <a-icon type="close-circle" theme="filled"></a-icon> -->
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
  position: relative;
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

.selectable-list__item p {
  margin: 0;
}
.selectable-list__item:hover .selectable-list__close {
  opacity: 1;
}

.selectable-list__desc {
  font-size: 12px;
  white-space: nowrap;
  text-overflow: ellipsis;
  overflow: hidden;
}
.selectable-list__close {
  opacity: 0;
  position: absolute;
  right: 5px;
  top: 5px;
  transition: all 0.3s;
}
</style>
