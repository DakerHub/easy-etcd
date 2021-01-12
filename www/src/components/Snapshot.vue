<template>
  <div class="snapshot">
    <a-tooltip title="备份与恢复">
      <a-button type="link" icon="history" @click="open" ghost size="small"></a-button>
    </a-tooltip>

    <a-modal v-model="visible" title="备份与恢复" :mask="false" :footer="null">
      <p class="tip">备份为浅备份，只记录当前的键值，不会记录版本及其他信息</p>
      <a-list :loading="loading" item-layout="horizontal" :data-source="data">
        <a-list-item slot="renderItem" slot-scope="item, index">
          <a slot="actions">
            <a-popconfirm title="确认删除？" @confirm="deleteSnapshot(item)">
              <a-button :loading="item._submitting" type="link" size="small"> 删除 </a-button>
            </a-popconfirm>
          </a>
          <a slot="actions">
            <a-popconfirm title="确认恢复该备份？" @confirm="restoreFrom(item)">
              <a-button :loading="item._submitting" type="link" size="small"> 恢复 </a-button>
            </a-popconfirm>
          </a>
          <div class="snapshot-item">
            <span>{{ index + 1 }}. <span class="light-text">创建于 </span>{{ timeFormat(item.createAt) }}</span>
            <span>文档数：{{ item.docCount }}</span>
          </div>
        </a-list-item>
      </a-list>
      <a-button icon="plus" :loading="submitting" type="primary" @click="createSnapshot">备份当前</a-button>
    </a-modal>
  </div>
</template>

<script>
import { getSnapshots, createSnapshot, restore, deleteSnapshot } from "./../api/snapshot.js";
import { timeFormat } from "./../util/date.js";

export default {
  name: "Snapshot",
  props: {
    connectAddr: {
      type: String,
      default: "",
    },
  },
  data() {
    return {
      visible: false,
      loading: true,
      data: [],
      submitting: false,
    };
  },
  methods: {
    timeFormat,
    open() {
      this.visible = true;
      this.fetchData();
    },
    async fetchData() {
      this.loading = true;
      try {
        const res = await getSnapshots(this.connectAddr);
        res.data.sort((a, b) => {
          return new Date(b.createAt) - new Date(a.createAt);
        });
        this.data = res.data;
      } catch (error) {
        console.error(error);
      } finally {
        this.loading = false;
      }
    },
    async deleteSnapshot(item) {
      this.$set(item, "_submitting", true);
      try {
        await deleteSnapshot(this.connectAddr, item.id);
        this.fetchData();
        this.$message.success("删除成功");
      } catch (error) {
        console.error(error);
      } finally {
        this.$set(item, "_submitting", false);
      }
    },
    async restoreFrom(item) {
      this.$set(item, "_submitting", true);
      try {
        await restore(this.connectAddr, item.id);
        this.$emit("restore");
        this.$message.success("恢复成功");
        this.visible = false;
      } catch (error) {
        console.error(error);
      } finally {
        this.$set(item, "_submitting", false);
      }
    },
    async createSnapshot() {
      this.submitting = true;
      try {
        await createSnapshot(this.connectAddr);
        this.fetchData();
        this.$message.success("备份成功");
      } catch (error) {
        console.error(error);
      } finally {
        this.submitting = false;
      }
    },
  },
};
</script>

<style scoped>
.snapshot {
  display: inline-block;
}
.snapshot-item {
  width: 100%;
  display: flex;
  justify-content: space-between;
}
.light-text {
  color: #999;
}
.tip {
  color: #999;
  font-size: 14px;
}
.ant-list {
  padding: 20px 0;
}
</style>
