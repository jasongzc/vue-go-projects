一、项目结构
vue-go-projects
├── backend/          # Go后端
├── frontend/         # Vue前端
└── README.md

1. 初始化后端项目
bash
mkdir -p my-web-app/backend
cd my-web-app/backend
go mod init webapp-backend
2. 安装依赖
bash
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/dgrijalva/jwt-go
go get -u golang.org/x/crypto/bcrypt
3. 项目结构
backend/
├── main.go
├── go.mod
├── go.sum
├── config/
│   └── config.go
├── models/
│   ├── user.go
│   └── database.go
├── routes/
│   └── routes.go
├── controllers/
│   └── auth.go
├── middleware/
│   └── auth.go
└── utils/
    └── jwt.go

三、前端实现（Vue3 + Vite + Element Plus）
1. 初始化前端项目
bash
cd my-web-app
npm create vue@latest frontend
# 选择以下配置：
# √ Add TypeScript? No
# √ Add JSX Support? No
# √ Add Vue Router for Single Page Application development? Yes
# √ Add Pinia for state management? Yes
# √ Add Vitest for Unit Testing? No
# √ Add an End-to-End Testing Solution? No
# √ Add ESLint for code quality? No
2. 安装Element Plus
bash
cd frontend
npm install element-plus @element-plus/icons-vue
npm install axios
3. 项目结构
frontend/
├── src/
│   ├── api/          # API接口
│   ├── router/       # 路由
│   ├── store/        # Pinia状态管理
│   ├── views/        # 页面组件
│   ├── components/   # 公共组件
│   ├── utils/        # 工具函数
│   └── App.vue   


1. 启动后端
bash
cd backend
go run main.go

2. 启动前端
bash
cd frontend
npm install
npm run dev


3. 安装推荐扩展
Go: Go语言支持
Vue Language Features (Volar): Vue3开发支持
Element Plus Snippets: Element Plus代码片段
MySQL: 数据库管理

这个完整的全栈应用包含了：
后端: Go + Gin + JWT认证 + MySQL
前端: Vue3 + Vite + Element Plus + Pinia状态管理
功能: 用户注册、登录、个人信息管理
安全: 密码加密、JWT令牌、路由守卫

![Uploading image.png…]()

