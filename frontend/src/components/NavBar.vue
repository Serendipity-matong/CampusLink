<template>
  <div class="navbar-wrapper">
    <el-menu
      :default-active="activeIndex"
      mode="horizontal"
      :ellipsis="false"
      background-color="#ffffff"
      text-color="#2C3E50"
      active-text-color="#1E88E5"
      @select="handleSelect"
      class="campus-navbar"
    >
      <div class="navbar-content">
        <div class="logo-section">
          <el-icon :size="28" color="#1E88E5"><School /></el-icon>
          <h1 class="logo-text">CampusLink</h1>
          <span class="logo-subtitle">校园互联</span>
        </div>
        
        <div class="menu-section">
          <el-menu-item index="/">
            <el-icon><HomeFilled /></el-icon>
            <span>首页</span>
          </el-menu-item>
          <el-menu-item index="/products">
            <el-icon><ShoppingBag /></el-icon>
            <span>二手市场</span>
          </el-menu-item>
          <el-menu-item index="/tasks">
            <el-icon><Bicycle /></el-icon>
            <span>跑腿代购</span>
          </el-menu-item>
          <el-menu-item index="/orders">
            <el-icon><Document /></el-icon>
            <span>我的订单</span>
          </el-menu-item>
        </div>
        
        <div class="user-section">
          <el-button type="primary" round size="default" @click="goToProfile">
            <el-icon><User /></el-icon>
            <span>个人中心</span>
          </el-button>
        </div>
      </div>
    </el-menu>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { 
  School, 
  HomeFilled, 
  ShoppingBag, 
  Bicycle, 
  Document, 
  User 
} from '@element-plus/icons-vue'

const router = useRouter()
const route = useRoute()
const activeIndex = ref('/')

watch(() => route.path, (newPath) => {
  activeIndex.value = newPath
}, { immediate: true })

const handleSelect = (key) => {
  router.push(key)
}

const goToProfile = () => {
  router.push('/profile')
}
</script>

<style scoped>
.navbar-wrapper {
  position: sticky;
  top: 0;
  z-index: 1000;
  box-shadow: 0 2px 12px rgba(30, 136, 229, 0.08);
}

.campus-navbar {
  border-bottom: 1px solid #E0E6ED;
}

.navbar-content {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
}

.logo-section {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-right: 40px;
  cursor: pointer;
}

.logo-text {
  font-size: 24px;
  font-weight: 700;
  background: linear-gradient(135deg, #1E88E5 0%, #66BB6A 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  margin: 0;
}

.logo-subtitle {
  font-size: 12px;
  color: #7F8C8D;
  font-weight: 500;
}

.menu-section {
  display: flex;
  flex: 1;
  gap: 8px;
}

.user-section {
  margin-left: 20px;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .navbar-content {
    padding: 0 12px;
  }
  
  .logo-subtitle {
    display: none;
  }
  
  .menu-section span {
    display: none;
  }
}
</style>

