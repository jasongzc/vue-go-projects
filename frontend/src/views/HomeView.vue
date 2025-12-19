<template>
  <div class="home-container">
    <el-container class="main-container">
      <!-- 导航栏 -->
      <el-header class="header">
        <div class="nav-container">
          <div class="nav-left">
            <div class="logo">
              <el-icon size="28" color="#409EFF"><ElementPlus /></el-icon>
              <span class="logo-text">vue-golang框架</span>
            </div>
            <el-menu
              mode="horizontal"
              :default-active="activeIndex"
              :collapse="false"
              @select="handleSelect"
              class="nav-menu"
            >
              <el-menu-item index="1">首页</el-menu-item>
              <el-menu-item index="2">功能</el-menu-item>
              <el-menu-item index="3" style="display:none">文档</el-menu-item>
            </el-menu>
          </div>
          
          <div class="nav-right">
            <div v-if="authStore.isAuthenticated" class="user-section">
              <el-dropdown @command="handleCommand" trigger="click">
                <div class="user-dropdown-trigger">
                  <el-avatar :size="32" class="user-avatar">
                    {{ authStore.user?.username?.charAt(0)?.toUpperCase() || 'U' }}
                  </el-avatar>
                  <span class="username">{{ authStore.user?.username }}</span>
                  <el-icon class="dropdown-icon"><arrow-down /></el-icon>
                </div>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item command="profile">
                      <el-icon><User /></el-icon>
                      <span>个人中心</span>
                    </el-dropdown-item>
                    <el-dropdown-item command="settings">
                      <el-icon><Setting /></el-icon>
                      <span>修改密码</span>
                    </el-dropdown-item>
                    <el-dropdown-item divided command="logout">
                      <el-icon><SwitchButton /></el-icon>
                      <span>退出登录</span>
                    </el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </div>
            
            <div v-else class="auth-section">
              <el-button 
                type="text" 
                class="auth-btn"
                @click="router.push('/login')"
              >
                登录
              </el-button>
              <el-button 
                type="primary" 
                plain 
                size="small"
                @click="router.push('/register')"
              >
                注册
              </el-button>
            </div>
          </div>
        </div>
      </el-header>
      
      <!-- 主内容区域 -->
      <el-main class="main-content">
        <!-- 欢迎横幅 -->
        <div class="welcome-banner">
          <div class="banner-content">
            <h1 class="banner-title">
              欢迎使用 <span class="highlight"> web 应用</span>
            </h1>
            <p class="banner-subtitle">
              现代化的全栈应用解决方案
            </p>
            <div class="banner-actions" v-if="!authStore.isAuthenticated">
              <el-button 
                type="primary" 
                size="large" 
                @click="router.push('/login')"
                class="banner-btn"
              >
                立即开始
              </el-button>
              <el-button 
                size="large" 
                @click="router.push('/register')"
                class="banner-btn secondary"
              >
                免费注册
              </el-button>
            </div>
          </div>
          <div class="banner-illustration">
            <el-icon size="120" color="var(--el-color-primary-light-3)">
              <DataLine />
            </el-icon>
          </div>
        </div>

        <!-- 用户信息卡片（已登录时显示） -->
       <!--  <div v-if="authStore.isAuthenticated" class="user-info-card">
          <el-card shadow="hover" class="user-card">
            <template #header>
              <div class="user-card-header">
                <div class="user-avatar-large">
                  <el-avatar :size="64">
                    {{ authStore.user?.username?.charAt(0)?.toUpperCase() || 'U' }}
                  </el-avatar>
                </div>
                <div class="user-info">
                  <h3 class="user-name">{{ authStore.user?.username }}</h3>
                  <p class="user-email">{{ authStore.user?.email }}</p>
                </div>
              </div>
            </template>
            <div class="user-stats-grid">
              <div class="stat-item">
                <div class="stat-label">用户ID</div>
                <div class="stat-value">{{ authStore.user?.id }}</div>
              </div>
              <div class="stat-item">
                <div class="stat-label">注册时间</div>
                <div class="stat-value">{{ authStore.user?.createdAt }}</div>
              </div>
              <div class="stat-item">
                <div class="stat-label">账号状态</div>
                <div class="stat-value">
                  <el-tag type="success" size="small">正常</el-tag>
                </div>
              </div>
            </div>
            <div class="user-actions">
              <el-button type="primary" @click="router.push('/profile')" plain>
                编辑资料
              </el-button>
            </div>
          </el-card>
        </div> -->

        <!-- 功能介绍区域 -->
        <div class="features-section">
          <div class="section-header">
            <h2 class="section-title">核心功能</h2>
            <p class="section-subtitle">提供完整的应用解决方案</p>
          </div>
          
          <el-row :gutter="24" class="features-grid">
            <el-col :xs="24" :sm="12" :md="8" :lg="8">
              <el-card class="feature-card" shadow="hover">
                <div class="feature-icon">
                  <el-icon><UserFilled /></el-icon>
                </div>
                <h3 class="feature-title">用户管理</h3>
                <p class="feature-desc">
                  完整的用户认证系统，支持注册、登录、个人信息管理和权限控制
                </p>
              </el-card>
            </el-col>
            
            <el-col :xs="24" :sm="12" :md="8" :lg="8">
              <el-card class="feature-card" shadow="hover">
                <div class="feature-icon security">
                  <el-icon><Lock /></el-icon>
                </div>
                <h3 class="feature-title">安全认证</h3>
                <p class="feature-desc">
                  基于 JWT 的 Token 认证，密码加密存储，确保数据安全
                </p>
              </el-card>
            </el-col>
            
            <el-col :xs="24" :sm="12" :md="8" :lg="8">
              <el-card class="feature-card" shadow="hover">
                <div class="feature-icon tech">
                  <el-icon><Cpu /></el-icon>
                </div>
                <h3 class="feature-title">现代化技术栈</h3>
                <p class="feature-desc">
                  Vue3 + Vite + Element Plus + Go + Gin + MySQL，前后端分离架构
                </p>
              </el-card>
            </el-col>

            <el-col :xs="24" :sm="12" :md="8" :lg="8">
              <el-card class="feature-card" shadow="hover">
                <div class="feature-icon api">
                  <el-icon><Connection /></el-icon>
                </div>
                <h3 class="feature-title">RESTful API</h3>
                <p class="feature-desc">
                  标准的 RESTful 接口设计，支持跨域请求，易于集成和扩展
                </p>
              </el-card>
            </el-col>

            <el-col :xs="24" :sm="12" :md="8" :lg="8">
              <el-card class="feature-card" shadow="hover">
                <div class="feature-icon responsive">
                  <el-icon><Monitor /></el-icon>
                </div>
                <h3 class="feature-title">响应式设计</h3>
                <p class="feature-desc">
                  适配桌面端和移动端，提供一致的用户体验
                </p>
              </el-card>
            </el-col>

            <el-col :xs="24" :sm="12" :md="8" :lg="8">
              <el-card class="feature-card" shadow="hover">
                <div class="feature-icon deploy">
                  <el-icon><Cloudy /></el-icon>
                </div>
                <h3 class="feature-title">易于部署</h3>
                <p class="feature-desc">
                  提供 Docker 支持，一键部署到各种云平台
                </p>
              </el-card>
            </el-col>
          </el-row>
        </div>

        <!-- 技术栈展示 -->
        <div class="tech-stack-section">
          <div class="section-header">
            <h2 class="section-title">技术架构</h2>
            <p class="section-subtitle">基于现代化技术栈构建</p>
          </div>
          
          <el-card class="tech-stack-card" shadow="never">
            <div class="tech-stack">
              <div class="tech-item">
                <div class="tech-icon vue">
                  <el-icon><StarFilled /></el-icon>
                </div>
                <div class="tech-info">
                  <h4>Vue 3</h4>
                  <p>渐进式 JavaScript 框架</p>
                </div>
              </div>
              
              <div class="tech-item">
                <div class="tech-icon element">
                  <el-icon><ElementPlus /></el-icon>
                </div>
                <div class="tech-info">
                  <h4>Element Plus</h4>
                  <p>基于 Vue 3 的组件库</p>
                </div>
              </div>
              
              <div class="tech-item">
                <div class="tech-icon go">
                  <el-icon><Platform /></el-icon>
                </div>
                <div class="tech-info">
                  <h4>Go + Gin</h4>
                  <p>高性能后端服务</p>
                </div>
              </div>
              
              <div class="tech-item">
                <div class="tech-icon mysql">
                  <el-icon><DataAnalysis /></el-icon>
                </div>
                <div class="tech-info">
                  <h4>MySQL</h4>
                  <p>关系型数据库</p>
                </div>
              </div>
            </div>
          </el-card>
        </div>
      </el-main>
      
      <!-- 页脚 -->
     <!--  <el-footer class="footer">
        <div class="footer-content">
          <div class="footer-links">
            <el-link type="info" :underline="false">关于我们</el-link>
            <el-link type="info" :underline="false">文档</el-link>
            <el-link type="info" :underline="false">GitHub</el-link>
            <el-link type="info" :underline="false">联系我们</el-link>
          </div>
          <p class="copyright">
            © 2024 Web Application. All rights reserved.
          </p>
        </div>
      </el-footer> -->
    </el-container>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import { ElMessage, ElMessageBox } from 'element-plus'
import { 
  User, 
  Lock, 
  Setting, 
  ArrowDown, 
  SwitchButton,
  ElementPlus,
  DataLine,
  UserFilled,
  Cpu,
  Connection,
  Monitor,
  Cloudy,
  StarFilled,
  Platform,
  DataAnalysis
} from '@element-plus/icons-vue'

const router = useRouter()
const authStore = useAuthStore()

const activeIndex = ref('1')

const userCreatedAt = computed(() => {
  if (!authStore.user?.createdAt) return '未知'
  const date = new Date(authStore.user.createdAt)
  return date.toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
})

const handleSelect = (key) => {
  console.log('Select:', key)
}

const handleCommand = (command) => {
  if (command === 'profile') {
    router.push('/profile')
  } else if (command === 'settings') {
    router.push('/usersettings')
  } else if (command === 'logout') {
    handleLogout()
  }
}

const handleLogout = () => {
  ElMessageBox.confirm(
    '确定要退出登录吗？',
    '提示',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(() => {
    authStore.clearAuth()
    ElMessage.success('已退出登录')
    router.push('/login')
  })
}
</script>

<style scoped>
.home-container {
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
}

.main-container {
  min-height: 100vh;
}

/* 导航栏样式 */
.header {
  padding: 0;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  position: sticky;
  top: 0;
  z-index: 1000;
}

.nav-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 60px;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 24px;
}

.nav-left {
  display: flex;
  align-items: center;
  gap: 40px;
}

.logo {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
}

.logo-text {
  font-size: 20px;
  font-weight: 600;
  color: var(--el-color-primary);
}

.nav-menu {
  border-bottom: none;
}

.nav-right {
  display: flex;
  align-items: center;
}

.user-section {
  display: flex;
  align-items: center;
}

.user-dropdown-trigger {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 12px;
  border-radius: 6px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.user-dropdown-trigger:hover {
  background-color: var(--el-color-primary-light-9);
}

.user-avatar {
  background: var(--el-color-primary);
  color: white;
}

.username {
  font-weight: 500;
  color: var(--el-text-color-primary);
}

.dropdown-icon {
  font-size: 12px;
  color: var(--el-text-color-secondary);
}

.auth-section {
  display: flex;
  align-items: center;
  gap: 16px;
}

.auth-btn {
  padding: 8px 16px;
}

/* 主内容样式 */
.main-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 24px;
}

/* 欢迎横幅 */
.welcome-banner {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: linear-gradient(135deg, var(--el-color-primary) 0%, var(--el-color-primary-light-3) 100%);
  border-radius: 16px;
  padding: 48px;
  margin-bottom: 40px;
  color: white;
  min-height: 320px;
}

.banner-content {
  flex: 1;
  max-width: 600px;
}

.banner-title {
  font-size: 3rem;
  font-weight: 700;
  margin-bottom: 16px;
  line-height: 1.2;
}

.banner-title .highlight {
  color: #ffd04b;
}

.banner-subtitle {
  font-size: 1.25rem;
  margin-bottom: 32px;
  opacity: 0.9;
}

.banner-actions {
  display: flex;
  gap: 16px;
}

.banner-btn {
  padding: 12px 32px;
  font-size: 1rem;
}

.banner-btn.secondary {
  background: rgba(255, 255, 255, 0.2);
  border-color: rgba(255, 255, 255, 0.3);
  color: white;
}

.banner-btn.secondary:hover {
  background: rgba(255, 255, 255, 0.3);
}

.banner-illustration {
  flex: 0 0 auto;
}

/* 用户信息卡片 */
.user-info-card {
  margin-bottom: 40px;
}

.user-card {
  border-radius: 12px;
}

.user-card-header {
  display: flex;
  align-items: center;
  gap: 24px;
}

.user-avatar-large {
  flex: 0 0 auto;
}

.user-avatar-large .el-avatar {
  background: var(--el-color-primary);
  color: white;
  font-size: 24px;
  font-weight: bold;
}

.user-info {
  flex: 1;
}

.user-name {
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 8px;
  color: var(--el-text-color-primary);
}

.user-email {
  color: var(--el-text-color-secondary);
  margin: 0;
}

.user-stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 24px;
  margin: 24px 0;
}

.stat-item {
  text-align: center;
}

.stat-label {
  font-size: 0.875rem;
  color: var(--el-text-color-secondary);
  margin-bottom: 8px;
}

.stat-value {
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--el-text-color-primary);
}

.user-actions {
  display: flex;
  justify-content: center;
  margin-top: 24px;
}

/* 功能区域 */
.section-header {
  text-align: center;
  margin-bottom: 48px;
}

.section-title {
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 16px;
  color: var(--el-text-color-primary);
}

.section-subtitle {
  font-size: 1.125rem;
  color: var(--el-text-color-secondary);
}

.features-grid {
  margin-bottom: 60px;
}

.feature-card {
  height: 100%;
  border-radius: 12px;
  transition: transform 0.3s;
  text-align: center;
  padding: 32px 24px;
}

.feature-card:hover {
  transform: translateY(-4px);
}

.feature-icon {
  width: 64px;
  height: 64px;
  margin: 0 auto 24px;
  background: var(--el-color-primary-light-9);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.feature-icon .el-icon {
  font-size: 28px;
  color: var(--el-color-primary);
}

.feature-icon.security {
  background: var(--el-color-success-light-9);
}

.feature-icon.security .el-icon {
  color: var(--el-color-success);
}

.feature-icon.tech {
  background: var(--el-color-warning-light-9);
}

.feature-icon.tech .el-icon {
  color: var(--el-color-warning);
}

.feature-icon.api {
  background: var(--el-color-info-light-9);
}

.feature-icon.api .el-icon {
  color: var(--el-color-info);
}

.feature-icon.responsive {
  background: var(--el-color-danger-light-9);
}

.feature-icon.responsive .el-icon {
  color: var(--el-color-danger);
}

.feature-icon.deploy {
  background: #f0f7ff;
}

.feature-icon.deploy .el-icon {
  color: #1890ff;
}

.feature-title {
  font-size: 1.25rem;
  font-weight: 600;
  margin-bottom: 16px;
  color: var(--el-text-color-primary);
}

.feature-desc {
  color: var(--el-text-color-secondary);
  line-height: 1.6;
}

/* 技术栈区域 */
.tech-stack-card {
  border-radius: 12px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.tech-stack {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 32px;
}

.tech-item {
  display: flex;
  align-items: center;
  gap: 16px;
}

.tech-icon {
  width: 56px;
  height: 56px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.tech-icon .el-icon {
  font-size: 28px;
  color: white;
}

.tech-info h4 {
  margin: 0 0 4px 0;
  font-size: 1.125rem;
  font-weight: 600;
}

.tech-info p {
  margin: 0;
  opacity: 0.9;
  font-size: 0.875rem;
}

/* 页脚样式 */
.footer {
  background: var(--el-bg-color-page);
  border-top: 1px solid var(--el-border-color-light);
  padding: 32px 0;
  margin-top: 60px;
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
  text-align: center;
}

.footer-links {
  display: flex;
  justify-content: center;
  gap: 32px;
  margin-bottom: 16px;
}

.copyright {
  color: var(--el-text-color-secondary);
  font-size: 0.875rem;
  margin: 0;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .nav-container {
    padding: 0 16px;
  }
  
  .nav-left {
    gap: 20px;
  }
  
  .welcome-banner {
    flex-direction: column;
    text-align: center;
    padding: 32px 24px;
  }
  
  .banner-title {
    font-size: 2rem;
  }
  
  .banner-illustration {
    margin-top: 24px;
  }
  
  .user-card-header {
    flex-direction: column;
    text-align: center;
  }
  
  .features-grid .el-col {
    margin-bottom: 16px;
  }
  
  .footer-links {
    flex-wrap: wrap;
    gap: 16px;
  }
}
</style>