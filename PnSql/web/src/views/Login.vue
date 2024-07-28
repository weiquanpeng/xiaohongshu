<template>
  <div class="login-box">
    <div class="left">
      <el-form :model="formInline" class="demo-form-inline">
        <el-form-item label="用户:" >
          <el-input placeholder="Username" v-model="loginUser.username" clearable />
        </el-form-item>
        <el-form-item label="密码:">
          <el-input placeholder="Password" v-model="loginUser.password" clearable />
        </el-form-item>
        <el-form-item label="验证:">
          <el-input placeholder="Code" v-model="loginUser.code" clearable />
        </el-form-item>
        <el-form-item>
          <img v-if="codeURL" :src="codeURL" alt="" class="captcha-image-background" @click="getCaptcha">
          <el-button type="primary" @click="onSubmit" style="float: right; margin-left: 30px;">登录</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="right">
      <img src="@/assets/images/dashboard.png" alt="">
    </div>
  </div>
</template>

<script setup>
import axios from "axios";
const codeURL=ref("")
const loginUser=reactive({
  username:"",
  password:"",
  id:"",
  code:""
})

const getCaptcha = async () =>{
  const resp = await axios.get("http://127.0.0.1:2379/api/v1/public/login/get")
  const {baseURL,id} = resp.data.data
  codeURL.value=baseURL
  loginUser.id=id
}
const onSubmit = async () =>{
  const resp = await axios.post("http://127.0.0.1:2379/api/v1/public/login/post",loginUser)
  if (resp.data.code ===200){
    alert(resp.data.data)
    // window.location.href="/"
  }else {
    if (resp.data.code !==200){
      alert(resp.data.data)
      getCaptcha()
    }
  }
}

onMounted(()=>{
  getCaptcha()
    })
</script>

<style scoped lang="scss">
.login-box{
  display: flex;
  flex-wrap: wrap;
  background: url("@/assets/images/login_background.jpg") no-repeat center center / cover;
  .left{
    width: 50%;
    height: 100vh;
    display: flex; /* 使.left成为flex容器 */
    justify-content: center; /* 水平居中 */
    align-items: center; /* 垂直居中 */
  }
  .right{
    width: 50%;
    height: 100vh;
    display: flex;
    justify-content: center;  /* 水平居中 */
    align-items: center;      /* 垂直居中 */
  }
  .right img{
    max-width: 80%;    /* 图片最大宽度为容器宽度的100% */
    max-height: 80%;   /* 图片最大高度为容器高度的100% */

  }
  .demo-form-inline .el-input {
    --el-input-width: 220px;
  }

  .demo-form-inline .el-select {
    --el-select-width: 220px;
  }
}
.captcha-image-background {
  display: inline-block;
  margin-left: 60px;
  width: 100px;
  height: 30px;
  background-size: contain; /* 保持宽高比 */
}

</style>