import vue from '@vitejs/plugin-vue'
import { defineConfig } from 'vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],

  envDir: process.env.CONFIG_PATH || ".",
  mode: process.env.MODE,

  server: {
    port: 3001
  }
})
