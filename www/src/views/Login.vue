<template>
  <div class="login">
    <div class="login-form">
      <a-form-model ref="ruleForm" :model="form" :rules="rules" v-bind="layout">
        <a-form-model-item prop="username">
          <a-input name="username" v-model="form.username" placeholder="请输入用户名" autocomplete="off" />
        </a-form-model-item>
        <a-form-model-item prop="password">
          <a-input
            name="password"
            v-model="form.password"
            type="password"
            placeholder="请输入密码"
            autocomplete="off"
          />
        </a-form-model-item>
        <a-form-model-item :wrapper-col="{ span: 24, offset: 0 }">
          <a-button type="primary" block @click="submitForm('ruleForm')"> 登录 </a-button>
        </a-form-model-item>
      </a-form-model>
    </div>
  </div>
</template>

<script>
import { login } from "./../api/auth.js";

export default {
  name: "Login",
  data() {
    return {
      form: {
        username: "",
        password: "",
      },
      submitting: false,
      rules: {
        username: [{ required: true, message: "必填" }],
        password: [{ required: true, message: "必填" }],
      },
      layout: {
        labelCol: { span: 0 },
        wrapperCol: { span: 24 },
      },
    };
  },
  methods: {
    async submitForm(formName) {
      try {
        await this.$refs[formName].validate();
      } catch (error) {
        return error;
      }

      this.submitting = true;
      try {
        await login(this.form.username, this.form.password);
        this.$message.success("登录成功");
        this.$router.push("/");
      } catch (error) {
        console.error(error);
      } finally {
        this.submitting = false;
      }
    },
  },
};
</script>

<style  scoped>
.login {
  margin-top: 80px;
  display: flex;
  justify-content: center;
  align-items: center;
}
.login-form {
  width: 420px;
  padding: 80px 40px 20px;
  border-radius: 6px;
  background-color: #fff;
  box-shadow: 0px 6px 17px 2px rgba(10, 10, 10, 0.161);
}
.login-form .ant-input {
  font-size: 16px;
}
</style>
