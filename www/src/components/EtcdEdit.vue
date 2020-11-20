<template>
  <div class="etcd-edit">
    <!-- 目录 begin -->
    <div class="tree-view">
      <template v-if="treeData.length">
        <div class="tree-view-tool">
          <ModelInputKey
            title="新建键值对"
            placeholder="请输入建名，以“/”作为分隔符。比如：/config/app/user"
            @change="checkKeyExist"
          >
            <template v-slot="{ open }">
              <!-- <a-icon type="plus" @click.native="open"></a-icon> -->
              <a-button type="link" icon="file-add" @click="open" ghost size="small"></a-button>
            </template>
          </ModelInputKey>
          <div>
            <a-radio-group v-model="listMode" size="small" button-style="solid">
              <a-radio-button value="list"> <a-icon type="ordered-list"></a-icon> </a-radio-button>
              <a-radio-button value="tree"> <a-icon type="folder"></a-icon> </a-radio-button>
            </a-radio-group>
          </div>
        </div>
        <div class="tree-view-main">
          <a-directory-tree
            v-if="listMode === 'tree'"
            showLine
            :tree-data="treeData"
            :selectedKeys.sync="currentSelect"
            :expandedKeys.sync="expandedKeys"
            :replaceFields="replaceFields"
            :showIcon="false"
          >
            <span slot="title">TITLE</span>
          </a-directory-tree>
          <a-tree
            v-if="listMode === 'list'"
            class="list-view"
            blockNode
            showIcon
            :tree-data="renderList"
            :selectedKeys.sync="currentSelect"
            :replaceFields="{
              title: 'key',
              key: 'key',
            }"
          >
            <a-icon slot="file" type="file" />

            <template slot="tree-title" slot-scope="{ key, isNew }">
              <a-icon v-if="isNew" type="file-exclamation" theme="filled" style="color: cyan; margin-right: 2px" />
              <span :title="key">{{ key }}</span>
            </template>
          </a-tree>
        </div>
        <div class="tree-view-footer">
          <span>共{{ rawList.length }}条记录</span>
        </div>
      </template>

      <div v-else-if="loading">加载中...</div>

      <div v-else class="etcd-empty">
        <p>没有任何配置 <a-icon type="meh" /></p>
        <ModelInputKey title="新建键值对" @change="checkKeyExist">
          <template v-slot="{ open }">
            <a-button type="primary" icon="plus" @click="open">新建一条配置</a-button>
          </template>
        </ModelInputKey>
      </div>
    </div>
    <!-- 目录 end -->

    <!-- 详情 begin -->
    <div v-if="hasSelectKey" class="edit-view">
      <div class="edit-head">
        <div>
          <span class="edit-title">{{ currentEdit.key }}</span>
          <span class="edit-title-tip" v-if="currentEdit.isNew">新建</span>
          <!-- 编辑key需要新增一天记录，然后删除该记录 -->
          <!-- <a-button type="link" ghost>
            <a-icon type="edit" />
          </a-button> -->
        </div>
        <div>
          <a-button type="danger" size="small" ghost @click="confirmDelete">删除</a-button>
        </div>
      </div>
      <div class="edit-content">
        <AceEditor v-model="currentEdit.value" :session="currentEdit.key" :mode="format"></AceEditor>
      </div>
      <div class="edit-content__footer">
        <a-button :loading="saving" type="primary" icon="sync" @click="save">保存</a-button>
        <div>
          <a-button v-if="format === 'json'" ghost style="margin-right: 5px" @click="formatJSON">格式化</a-button>
          <a-select style="width: 120px" v-model="format" @change="handleChange">
            <a-select-option value="plain_text">纯文本</a-select-option>
            <a-select-option value="json">JSON字符串</a-select-option>
          </a-select>
        </div>
      </div>
    </div>
    <div v-else class="etcd-empty">
      <a-empty description="👈  在左侧选择叶子节点进行编辑" />
    </div>
    <!-- 详情 end -->
  </div>
</template>

<script>
import Vue from "vue";
import { getAllKvs, putKv, deleteKv } from "./../api/kv.js";
import { listToTree, removeTreeNode, insertNode } from "./helper.etcdEdit.js";
import AceEditor from "./../components/AceEditor";
import ModelInputKey from "./ModelInputKey";

export default {
  name: "EtcdEdit",
  components: {
    AceEditor,
    ModelInputKey,
  },
  props: {
    connectAddr: {
      type: String,
      default: "",
    },
  },
  data() {
    return {
      treeData: [],
      rawList: [],
      expandedKeys: [],
      listMode: "tree",
      replaceFields: {
        title: "name",
      },
      format: "plain_text",
      currentSelect: [],
      currentEdit: {},
      loading: true,
      saving: false,
    };
  },
  computed: {
    hasSelectKey() {
      return this.currentSelect.length > 0;
    },
    renderList() {
      return this.rawList.map((i) => {
        return Object.assign(
          {
            slots: { icon: "file" },
            scopedSlots: { title: "tree-title" },
            isLeaf: true,
          },
          i
        );
      });
    },
  },
  watch: {
    currentSelect(value) {
      const selectKey = value[0];
      if (selectKey !== undefined) {
        const target = this.rawList.find((el) => el.key === selectKey);
        this.currentEdit = Object.assign({}, target || {});

        // 格式化JSON展示
        /* eslint-disable no-empty */
        if (this.format === "json") {
          this.formatJSON();
        }
      } else {
        this.currentEdit = {};
      }
    },
    connectAddr: {
      immediate: true,
      handler() {
        this.currentSelect = [];
        this.fetchKvs();
      },
    },
  },
  methods: {
    async fetchKvs() {
      this.loading = true;
      try {
        const res = await getAllKvs(this.connectAddr);
        this.rawList = res.kvs;
        this.listToTree();
      } catch (error) {
        console.error(error);
      } finally {
        this.loading = false;
      }
    },
    async save() {
      this.saving = true;

      try {
        let { key, value } = this.currentEdit;

        if (this.format === "json") {
          // 存储时还原JSON格式
          /* eslint-disable no-empty */
          try {
            value = JSON.stringify(JSON.parse(value), {}, 0);
          } catch (error) {}
        }

        await putKv(this.connectAddr, key, value);
        this.$alertSuccess("保存成功");
        this.updateList(key, { value, isNew: false });
        this.currentEdit.isNew = false;
      } catch (error) {
        console.error(error);
        this.$alertSuccess("保存失败");
      } finally {
        this.saving = false;
      }
    },
    confirmDelete() {
      const { key, isNew } = this.currentEdit;

      this.$confirm({
        title: `确认删除 "${key}"？`,
        content: "",
        onOk: async () => {
          try {
            if (!isNew) {
              await deleteKv(this.connectAddr, key);
            }

            this.$alertSuccess("删除成功");
            this.removeItem(key);
            this.currentSelect = [];
          } catch (error) {
            this.$alertError("删除失败");
          }
        },
        onCancel() {},
      });
    },
    listToTree() {
      const tree = listToTree(this.rawList, "key", (node) => {
        node.selectable = !!node.isLeaf;
        return node;
      });
      this.treeData = tree;
    },
    checkKeyExist(key) {
      if (!this.rawList.find((i) => i.key === key)) {
        return this.newLocalKv(key);
      }

      this.$alertError("该键值对已经存在");
      this.currentSelect = [key];
    },
    newLocalKv(key) {
      const kv = {
        key,
        value: "",
        isNew: true,
      };
      this.rawList.push(kv);

      if (!key.startsWith("/")) {
        key = "/" + key;
      }

      const node = Object.assign({ path: key }, kv);
      insertNode(this.treeData, node, "/", (node) => {
        node.selectable = !!node.isLeaf;
        return node;
      });

      this.currentSelect = [key];
    },
    updateList(key, newItem) {
      const target = this.rawList.find((el) => el.key === key);
      if (!target) {
        return this.$alertError("更新本地列表失败：" + key);
      }

      Object.assign(target, newItem);
    },
    removeItem(key) {
      removeTreeNode(this.treeData, key);
      const idx = this.rawList.findIndex((el) => el.key === key);
      if (idx != -1) {
        this.rawList.splice(idx, 1);
      }
    },
    handleChange() {
      if (this.format === "json") {
        this.formatJSON();
      }
    },
    formatJSON(indent = 2) {
      /* eslint-disable no-empty */
      try {
        this.currentEdit.value = JSON.stringify(JSON.parse(this.currentEdit.value), {}, indent);
      } catch (error) {}
    },
  },
};
</script>

<style scoped>
.etcd-edit {
  display: flex;
  height: calc(100vh - 120px);
  padding: 20px;
  background-color: var(--bg-color-1);
  border-radius: 4px;
  box-shadow: 0 0 15px 1px #31354b;
}
.tree-view {
  display: flex;
  flex-direction: column;
  width: 200px;
  flex-shrink: 0;
  padding-right: 10px;
  border-right: thin solid var(--color-1);
}
.tree-view-tool {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: var(--bg-color-3);
  padding: 5px;
  border-left: 2px solid #ccc;
}
.tree-view-main {
  height: 100%;
  overflow: auto;
}
.tree-view-tool .ant-btn-background-ghost:hover {
  background-color: var(--bg-color-2) !important;
}
.ant-radio-group >>> .ant-radio-button-wrapper {
  border-radius: 0 !important;
  border: none !important;
  background: rgba(255, 255, 255, 0.9);
}
.ant-radio-group >>> .ant-radio-button-wrapper-checked:not(.ant-radio-button-wrapper-disabled){
  background: #1890ff;
}

.edit-view {
  width: 100%;
  margin-left: 20px;
  display: flex;
  flex-direction: column;
}
.tree-view-footer {
  height: 20px;
  font-size: 12px;
  line-height: 20px;
  text-align: right;
  /* border-top: thin solid var(--border-color); */
}
.edit-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: var(--bg-color-3);
  padding: 10px 20px;
  border-radius: 4px;
  box-shadow: 0 0 9px 2px rgba(40, 43, 56, 0.369);
}
.edit-title {
  font-weight: bold;
  font-size: 16px;
  color: var(--color-0);
}
.edit-title-tip {
  margin-left: 5px;
  color: #4caf50;
  font-weight: bold;
  font-size: 12px;
}
.edit-content {
  height: 100%;
  background-color: var(--bg-color-3);
  padding: 10px 20px;
  border-radius: 4px;
  margin-top: 20px;
  box-shadow: 0 0 9px 2px rgba(40, 43, 56, 0.369);
}
.edit-content__footer {
  margin-top: 20px;
  display: flex;
  justify-content: space-between;
}
.ant-select >>> .ant-select-selection {
  background-color: rgba(255, 255, 255, 0.7) !important;
}

.etcd-empty {
  height: 100%;
  width: 100%;
  margin-top: 40px;
  display: flex;
  flex-direction: column;
  align-items: center;
}
.ace-editor {
  height: 100%;
  overflow: auto;
}

/* 定制化 */
.ant-tree {
  color: var(--color-1);
}
.ant-tree.ant-tree-show-line >>> li span.ant-tree-switcher {
  color: var(--color-1);
  background-color: transparent;
}
.ant-tree >>> li .ant-tree-node-content-wrapper {
  color: var(--color-1);
}
.ant-tree.ant-tree-directory >>> li.ant-tree-treenode-selected > span.ant-tree-node-content-wrapper::before {
  background: #1890ff !important;
}
.ant-tree.ant-tree-directory >>> li span.ant-tree-node-content-wrapper:hover::before,
.ant-tree.ant-tree-directory >>> .ant-tree-child-tree li span.ant-tree-node-content-wrapper:hover::before {
  background: var(--bg-color-0);
  color: var(--color-1);
}

.ant-tree >>> li.ant-tree-treenode-selected > span.ant-tree-node-content-wrapper {
  background: #1890ff !important;
  color: var(--color-0) !important;
}
.ant-tree >>> li span.ant-tree-node-content-wrapper:hover,
.ant-tree >>> .ant-tree-child-tree li span.ant-tree-node-content-wrapper:hover {
  background: var(--bg-color-0);
  color: var(--color-1);
}
.ant-tree.list-view >>> li span.ant-tree-switcher {
  display: none;
}
.ant-tree.list-view >>> li .ant-tree-node-content-wrapper {
  width: 100%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
