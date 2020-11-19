<template>
  <div class="ace-editor"></div>
</template>

<script>
export default {
  name: "AceEditor",
  props: {
    value: {
      type: String,
      default: "",
    },
    session: {
      type: String,
      default: "",
    },
    mode: {
      type: String,
      default: "plain_text",
    },
  },
  mounted() {
    this.initEditor();
  },
  data() {
    return {
      editor: null,
    };
  },
  watch: {
    session() {
      if (!this.editor) return;
      this.editor.setValue(this.value, 1);
    },
    value(value) {
      if (!this.editor) return;
      if (value !== this.editor.getValue()) {
        this.editor.setValue(this.value, 1);
      }
    },
    mode(mode) {
      if (!this.editor) return;
      this.editor.session.setMode(`ace/mode/${mode}`);
    },
  },
  methods: {
    initEditor() {
      /* eslint-disable no-undef */
      const editor = ace.edit(this.$el);
      this.editor = editor;
      editor.setValue(this.value, 1);
      editor.on("change", (e) => {
        const val = editor.getValue();
        this.$emit("input", val);
        this.$emit("change", val);
      });
      editor.setTheme("ace/theme/idle_fingers");
      editor.session.setMode(`ace/mode/${this.mode}`);
      this.$once("hook:destroyed", () => {
        editor.destroy();
      });
    },
  },
};
</script>

<style scoped>
.ace-editor {
  min-height: 300px;
}
</style>
