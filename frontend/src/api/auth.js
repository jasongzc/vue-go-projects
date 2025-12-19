import axios from 'axios'

const API_URL = import.meta.env.VITE_API_URL || '/api'

export default {
  // 注册
  register(data) {
    return axios.post(`${API_URL}/register`, data)
  },
  
  // 登录
  login(data) {
    return axios.post(`${API_URL}/login`, data)
  },
  
  // 获取用户信息
  getProfile() {
    return axios.get(`${API_URL}/profile`)
  } ,

  // 修改密码
  settingPassword(data) {
    return axios.post(`${API_URL}/changepassword`, data)
  }
}