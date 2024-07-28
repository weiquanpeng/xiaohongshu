import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'


export default defineConfig({
  plugins: [
    vue(),
      AutoImport({
          imports: ['vue', 'pinia'],
          resolvers: [ElementPlusResolver()],
      }),
      Components({
          resolvers: [ElementPlusResolver()],
      }),
  ],
  server:{
    port:9999
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
})
