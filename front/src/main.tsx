import React from 'react'
import ReactDOM from 'react-dom/client'
import "./assets/style/tailwind.css"
import './index.css'
import {RouterProvider} from "react-router-dom";
import router from "@/router";
import {ConfigProvider} from "antd";
import zhCN from 'antd/locale/zh_CN';
import { StyleProvider } from '@ant-design/cssinjs';
ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
      <ConfigProvider locale={zhCN}>
          <StyleProvider hashPriority="high">
              <RouterProvider router={router} />
          </StyleProvider>
      </ConfigProvider>
  </React.StrictMode>,
)
