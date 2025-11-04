# 🎨 CampusLink 前端项目

> 基于 Vue 3 + Vite + Element Plus 的校园互联平台前端

## 📦 技术栈

- **框架**: Vue 3 (Composition API)
- **构建工具**: Vite 5
- **UI 组件库**: Element Plus
- **路由**: Vue Router 4
- **状态管理**: Pinia
- **HTTP 客户端**: Axios
- **图标**: Element Plus Icons

## 🚀 快速开始

### 安装依赖

```bash
npm install
```

### 开发模式

```bash
npm run dev
```

访问: http://localhost:5173

### 生产构建

```bash
npm run build
```

### 预览构建结果

```bash
npm run preview
```

## 📁 项目结构

```
frontend/
├── public/                  # 静态资源
├── src/
│   ├── assets/             # 样式和静态资源
│   │   └── campus-theme.css # 校园主题样式
│   ├── components/         # 公共组件
│   │   ├── NavBar.vue     # 导航栏
│   │   └── Footer.vue     # 页脚
│   ├── views/              # 页面组件
│   │   ├── Home.vue       # 首页
│   │   ├── Products.vue   # 二手市场
│   │   ├── Tasks.vue      # 跑腿代购
│   │   ├── Orders.vue     # 订单管理
│   │   ├── Profile.vue    # 个人中心
│   │   └── Login.vue      # 登录页
│   ├── router/             # 路由配置
│   │   └── index.js
│   ├── App.vue             # 根组件
│   └── main.js             # 入口文件
├── index.html              # HTML 模板
├── vite.config.js          # Vite 配置
└── package.json            # 依赖配置
```

## 🎨 设计风格

### 校园风格配色

- **主色调**: #1E88E5 (校园蓝)
- **辅助色**: #66BB6A (活力绿)
- **文字色**: #2C3E50 (深灰)
- **背景色**: #F7F9FC (浅灰)

### 设计特点

- ✨ 清新活泼的配色
- 📱 响应式设计
- 🎯 简洁明了的布局
- 🌈 流畅的动画效果

## 🌐 路由

| 路径 | 组件 | 说明 |
|------|------|------|
| `/` | Home | 首页 |
| `/products` | Products | 二手市场 |
| `/tasks` | Tasks | 跑腿代购 |
| `/orders` | Orders | 订单管理 |
| `/profile` | Profile | 个人中心 |
| `/login` | Login | 登录页 |

## 🔧 开发建议

### 推荐 IDE 配置

- [VS Code](https://code.visualstudio.com/)
- [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (禁用 Vetur)

### ESLint 配置

```bash
npm install -D eslint eslint-plugin-vue
```

### 代码格式化

```bash
npm install -D prettier
```

## 📝 开发规范

### 组件命名

- 使用 PascalCase 命名组件
- 组件文件名与组件名保持一致

### 样式规范

- 使用 scoped 样式
- 遵循 BEM 命名规范
- 优先使用 CSS 变量

### Git 提交规范

```
feat: 新功能
fix: 修复bug
docs: 文档更新
style: 代码格式调整
refactor: 代码重构
test: 测试相关
chore: 构建/工具链相关
```

## 🔗 API 接口

### 基础配置

```javascript
// src/api/config.js
export const BASE_URL = 'http://localhost:8000'
```

### 请求拦截器

```javascript
// 添加 Token
axios.interceptors.request.use(config => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})
```

## 📱 响应式断点

```css
/* 手机 */
@media (max-width: 768px) { }

/* 平板 */
@media (min-width: 769px) and (max-width: 1024px) { }

/* 桌面 */
@media (min-width: 1025px) { }
```

## 🎯 待办事项

- [ ] 集成后端 API
- [ ] 添加全局状态管理
- [ ] 实现图片上传功能
- [ ] 添加实时聊天功能
- [ ] 优化移动端体验
- [ ] 添加单元测试

## 📄 License

MIT

---

**Made with ❤️ by CampusLink Team**
