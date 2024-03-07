import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from  'tailwindcss'
import autoprefixer from 'autoprefixer'
import { resolve } from "path";
// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
      react(),
    ],
    resolve: {
        alias: {
            "@": resolve(__dirname, './src'), // 这里是将src目录配置别名为 @ 方便在项目中导入src目录下的文件
        }
    },
    css:{
      postcss:{
          plugins:[
              tailwindcss,
              autoprefixer,
          ]
      }
    }
})
