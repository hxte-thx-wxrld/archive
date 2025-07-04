import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  publicDir: false,
  
  server: {
    allowedHosts: ["172.16.0.0/12", "frontend"],
    watch: {
      usePolling: true,
      interval: 100,
    },
    cors: true
  },
})
