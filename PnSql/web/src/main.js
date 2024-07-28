import './assets/base.css'
import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import Element from 'element-plus';
import 'element-plus/dist/index.css'

const app = createApp(App)

app
    .use(createPinia)
    .use(router)
    .use(Element)

app.mount('#app')
