<template>
  <div class="login-container">
    <el-card class="login-card">
      <template #header>
        <div class="card-header">
          <h2>修改密码</h2>
        </div>
      </template>
      
      <el-form 
        ref="settingFormRef" 
        :model="userForm" 
        :rules="userRules" 
        @submit.prevent="handleSetting"
      >
      <el-form-item prop="oldpassword">
          <el-input
            v-model="userForm.oldpassword"
            type="password"
            placeholder="旧密码"
            size="large"
            :prefix-icon="Lock"
            show-password
          />
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="userForm.password"
            type="password"
            placeholder="新密码"
            size="large"
            :prefix-icon="Lock"
            show-password
          />
        </el-form-item>
        
        <el-form-item prop="confirmPassword">
          <el-input
            v-model="userForm.confirmPassword"
            type="password"
            placeholder="确认新密码"
            size="large"
            :prefix-icon="Lock"
            show-password
          />
        </el-form-item>
        
        <el-form-item>
          <el-button 
            type="primary" 
            size="large" 
            :loading="loading"
            @click="handleSetting"
            style="width: 100%"
          >
            修改密码
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
import { Lock } from '@element-plus/icons-vue'
import { useAuthStore } from '@/store/auth'
import authApi from '@/api/auth'

const router = useRouter()
const authStore = useAuthStore()
const settingFormRef = ref()

const userForm = reactive({
  oldpassword: '',
  password: '',
  confirmPassword: ''
})

// 使用 computed 创建响应式的验证规则
const userRules = {
  password: [
    { required: true, message: '请输入新密码', trigger: 'blur' },
    { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '请确认新密码', trigger: 'blur' },
    { 
      validator: (rule, value, callback) => {
        if (value !== userForm.password) {
          callback(new Error('两次输入的密码不一致'))
        } else {
          callback()
        }
      }, 
      trigger: 'blur'
    }
  ]
}

const loading = ref(false)

const handleSetting = async () => {
  if (!settingFormRef.value) return
  
  // 验证表单
  await settingFormRef.value.validate(async (valid) => {
    if (!valid) {
      ElMessage.warning('请正确填写表单')
      return
    }
    
    loading.value = true
    
    try {
      const response = await authApi.settingPassword(userForm)
      
      ElMessage.success('密码修改成功')
      
      // 返回登录
      
        router.push('/Login')
      
      
    } catch (error) {
      ElMessage.error(error.response?.data?.error || '密码修改失败')
    } finally {
      loading.value = false
    }
  })
}
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
}

.login-card {
  width: 100%;
  max-width: 400px;
}

.card-header {
  text-align: center;
}

.card-header h2 {
  margin: 0;
  color: #333;
}
</style>