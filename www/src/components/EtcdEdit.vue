<template>
  <div class="etcd-edit">
    <!-- 目录 begin -->
    <div class="tree-view">
      <a-directory-tree
        v-if="treeData.length"
        :tree-data="treeData"
        :selectedKeys.sync="currentSelect"
        :replaceFields="replaceFields"
        showLine
        :showIcon="false"
      >
      </a-directory-tree>
      <div v-else class="etcd-empty">
        <p>没有任何配置 <a-icon type="meh" /></p>
        <a-button type="primary" icon="plus">新建一条配置</a-button>
      </div>
    </div>
    <!-- 目录 end -->

    <!-- 详情 begin -->
    <div v-if="hasSelectKey" class="edit-view">
      <div class="edit-head">
        <div>
          <span class="edit-title">{{ currentEdit.key }}</span>
          <!-- 编辑key需要新增一天记录，然后删除该记录 -->
          <!-- <a-button type="link" ghost>
            <a-icon type="edit" />
          </a-button> -->
        </div>
        <div>
          <a-button type="danger" size="small" ghost>删除</a-button>
        </div>
      </div>
      <div class="edit-content">
        <AceEditor v-model="currentEdit.value" :session="currentEdit.key" :mode="format"></AceEditor>
      </div>
      <div class="edit-content__footer">
        <a-button type="primary" icon="sync">保存</a-button>
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
import { getAllKvs } from "./../api/kv.js";
import { listToTree } from "./helper.etcdEdit.js";
import AceEditor from "./../components/AceEditor";

export default {
  name: "EtcdEdit",
  components: {
    AceEditor,
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
      replaceFields: {
        title: "name",
      },
      format: "plain_text",
      currentSelect: [],
      currentEdit: {},
    };
  },
  computed: {
    hasSelectKey() {
      return this.currentSelect.length > 0;
    },
  },
  watch: {
    currentSelect(value) {
      const selectKey = value[0];
      if (selectKey !== undefined) {
        const target = this.rawList.find((el) => el.key === selectKey);
        this.currentEdit = Object.assign({}, target || {});
      }
    },
  },
  created() {
    this.fetchKvs();
  },
  methods: {
    async fetchKvs() {
      const res = await getAllKvs(this.connectAddr);
      const tree = listToTree(res.kvs, "key", (node) => {
        node.selectable = !!node.isLeaf;
        return node;
      });
      this.rawList = res.kvs;
      this.treeData = tree;
    },
    handleChange() {
      if (this.format === "json") {
        this.formatJSON();
      }
    },
    formatJSON() {
      /* eslint-disable no-empty */
      try {
        this.currentEdit.value = JSON.stringify(JSON.parse(this.currentEdit.value), {}, 2);
      } catch (error) {}
    },
  },
};
</script>

<style scoped>
.etcd-edit {
  display: flex;
  height: calc(100vh - 200px);
  padding: 20px;
  background-color: var(--bg-color-1);
  border-radius: 4px;
  box-shadow: 0 0 15px 1px #31354b;
}
.tree-view {
  width: 200px;
  flex-shrink: 0;
  padding: 0 10px;
  border-right: thin solid var(--color-1);
}
.edit-view {
  width: 100%;
  margin-left: 20px;
  display: flex;
  flex-direction: column;
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
</style>
