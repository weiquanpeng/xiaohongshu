<template>
  <div class="app-container">
    <h1>{{ welcomeMessage }}</h1>
    <p>点击按钮看看会发生什么！</p>
    <button @click="changeMessage">改变消息</button>
    <p>{{ message }}</p>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const welcomeMessage = '欢迎来到 Vue 3 世界！';
const message = ref('这是一条初始消息。');

function changeMessage() {
  // 随机选择一个消息显示
  const messages = [
    'Vue 3 真的很有趣！',
    '你已经掌握了基础。',
    '继续探索 Vue 的世界！',
    '组件是 Vue 的核心。',
    '试试点击按钮看看会发生什么。'
  ];
  message.value = messages[Math.floor(Math.random() * messages.length)];
}
</script>

<style scoped>
.app-container {
  text-align: center;
  margin-top: 50px;
}

h1 {
  color: #42b983;
}

button {
  margin: 20px 0;
  padding: 10px 20px;
  font-size: 16px;
  cursor: pointer;
  background-color: #4fc08d;
  border: none;
  border-radius: 5px;
  color: white;
}

button:hover {
  background-color: #3aa373;
}
</style>