<template>
  <div class="profile-container">
    <el-card class="profile-card">
      <template #header>
        <div class="card-header">
          <h2>个人信息</h2>
        </div>
      </template>
      
      <div v-if="loading" class="loading">
        <el-skeleton :rows="5" animated />
      </div>
      
      <div v-else-if="user" class="profile-content">
        <div class="user-info">
          <el-descriptions :column="1" border>
            <el-descriptions-item label="用户ID">{{ user.id }}</el-descriptions-item>
            <el-descriptions-item label="用户名">{{ user.username }}</el-descriptions-item>
            <el-descriptions-item label="邮箱">{{ user.email }}</el-descriptions-item>
            <el-descriptions-item label="注册时间">{{ formatDate(user.createdAt) }}</el-descriptions-item>
          </el-descriptions>
        </div>
        
        <div class="profile-actions">
          <el-button type="primary" @click="router.back()">返回</el-button>
        </div>
      </div>
      
      <div v-else class="error-message">
        <el-result icon="error" title="获取用户信息失败">
          <template #extra>
            <el-button type="primary" @click="fetchUserProfile">重试</el-button>
          </template>
        </el-result>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { useAuthStore } from '@/store/auth'
import authApi from '@/api/auth'

const router = useRouter()
const authStore = useAuthStore()

const loading = ref(true)
const user = ref(null)

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleString()
}

const fetchUserProfile = async () => {
  loading.value = true
  try {
    const response = await authApi.getProfile()
    user.value = response.data.user
  } catch (error) {
    ElMessage.error('获取用户信息失败')
    console.error('Error fetching profile:', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchUserProfile()
})
</script>

<style scoped>
.profile-container {
  padding: 20px;
}

.profile-card {
  max-width: 800px;
  margin: 0 auto;
}

.card-header {
  text-align: center;
}

.card-header h2 {
  margin: 0;
  color: #333;
}

.loading {
  padding: 20px;
}

.user-info {
  margin-bottom: 30px;
}

.profile-actions {
  text-align: center;
  margin-top: 30px;
}

.error-message {
  text-align: center;
}
</style>