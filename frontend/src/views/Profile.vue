<template>
  <div class="profile-page campus-container">
    <h1 class="page-title">
      <el-icon :size="32"><User /></el-icon>
      个人中心
    </h1>

    <!-- 用户信息卡片 -->
    <div class="user-card campus-card">
      <div class="user-header campus-gradient-bg">
        <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=UserProfile" class="user-avatar-large" />
        <div class="user-info">
          <h2>{{ userInfo.name }}</h2>
          <p>{{ userInfo.studentId }} | {{ userInfo.major }}</p>
          <div class="user-badges">
            <el-tag type="primary" size="small">
              <el-icon><StarFilled /></el-icon>
              信用优秀
            </el-tag>
            <el-tag type="success" size="small">已认证</el-tag>
          </div>
        </div>
      </div>
      
      <div class="user-stats">
        <div class="stat-item">
          <div class="stat-value">{{ userInfo.completedOrders }}</div>
          <div class="stat-label">完成订单</div>
        </div>
        <div class="stat-divider"></div>
        <div class="stat-item">
          <div class="stat-value">{{ userInfo.rating }}</div>
          <div class="stat-label">评分</div>
        </div>
        <div class="stat-divider"></div>
        <div class="stat-item">
          <div class="stat-value">¥{{ userInfo.balance }}</div>
          <div class="stat-label">余额</div>
        </div>
      </div>
    </div>

    <!-- 功能菜单 -->
    <div class="menu-section">
      <h3 class="menu-title">我的服务</h3>
      <div class="menu-grid">
        <div class="menu-item campus-card" @click="goToOrders">
          <el-icon :size="32" color="#1E88E5"><Document /></el-icon>
          <span>我的订单</span>
        </div>
        <div class="menu-item campus-card" @click="goToProducts">
          <el-icon :size="32" color="#66BB6A"><ShoppingBag /></el-icon>
          <span>我发布的</span>
        </div>
        <div class="menu-item campus-card" @click="goToFavorites">
          <el-icon :size="32" color="#FFA726"><Star /></el-icon>
          <span>我的收藏</span>
        </div>
        <div class="menu-item campus-card" @click="goToWallet">
          <el-icon :size="32" color="#E74C3C"><Wallet /></el-icon>
          <span>我的钱包</span>
        </div>
      </div>
    </div>

    <!-- 设置菜单 -->
    <div class="menu-section">
      <h3 class="menu-title">设置</h3>
      <div class="settings-list campus-card">
        <div class="setting-item" @click="editProfile">
          <div class="setting-left">
            <el-icon><Edit /></el-icon>
            <span>编辑资料</span>
          </div>
          <el-icon><ArrowRight /></el-icon>
        </div>
        <div class="setting-item" @click="changePassword">
          <div class="setting-left">
            <el-icon><Lock /></el-icon>
            <span>修改密码</span>
          </div>
          <el-icon><ArrowRight /></el-icon>
        </div>
        <div class="setting-item" @click="addressManagement">
          <div class="setting-left">
            <el-icon><Location /></el-icon>
            <span>地址管理</span>
          </div>
          <el-icon><ArrowRight /></el-icon>
        </div>
        <div class="setting-item" @click="helpCenter">
          <div class="setting-left">
            <el-icon><QuestionFilled /></el-icon>
            <span>帮助中心</span>
          </div>
          <el-icon><ArrowRight /></el-icon>
        </div>
      </div>
    </div>

    <!-- 退出登录按钮 -->
    <el-button type="danger" plain class="logout-btn" @click="logout">
      <el-icon><SwitchButton /></el-icon>
      退出登录
    </el-button>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { 
  User, 
  StarFilled, 
  Document, 
  ShoppingBag, 
  Star, 
  Wallet,
  Edit,
  Lock,
  Location,
  QuestionFilled,
  ArrowRight,
  SwitchButton
} from '@element-plus/icons-vue'

const router = useRouter()

const userInfo = ref({
  name: '张同学',
  studentId: '202112345',
  major: '计算机科学与技术',
  completedOrders: 23,
  rating: 4.8,
  balance: 128.50
})

const goToOrders = () => {
  router.push('/orders')
}

const goToProducts = () => {
  ElMessage.info('我发布的商品/任务')
}

const goToFavorites = () => {
  ElMessage.info('我的收藏')
}

const goToWallet = () => {
  ElMessage.info('我的钱包')
}

const editProfile = () => {
  ElMessage.info('编辑资料功能开发中...')
}

const changePassword = () => {
  ElMessage.info('修改密码功能开发中...')
}

const addressManagement = () => {
  ElMessage.info('地址管理功能开发中...')
}

const helpCenter = () => {
  ElMessage.info('帮助中心')
}

const logout = () => {
  ElMessageBox.confirm('确定要退出登录吗？', '退出登录', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    ElMessage.success('已退出登录')
    router.push('/login')
  })
}
</script>

<style scoped>
.profile-page {
  padding: 40px 20px;
  min-height: calc(100vh - 200px);
}

.page-title {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 32px;
  font-weight: 700;
  margin-bottom: 30px;
  color: var(--campus-text-primary);
}

.user-card {
  padding: 0;
  overflow: hidden;
  margin-bottom: 30px;
}

.user-header {
  padding: 40px 30px;
  display: flex;
  gap: 24px;
  align-items: center;
  color: white;
}

.user-avatar-large {
  width: 100px;
  height: 100px;
  border-radius: 50%;
  border: 4px solid rgba(255, 255, 255, 0.3);
}

.user-info h2 {
  font-size: 28px;
  margin-bottom: 8px;
}

.user-info p {
  margin-bottom: 12px;
  opacity: 0.9;
}

.user-badges {
  display: flex;
  gap: 8px;
}

.user-stats {
  display: flex;
  padding: 30px;
}

.stat-item {
  flex: 1;
  text-align: center;
}

.stat-value {
  font-size: 28px;
  font-weight: 700;
  color: var(--campus-primary);
  margin-bottom: 8px;
}

.stat-label {
  font-size: 14px;
  color: var(--campus-text-secondary);
}

.stat-divider {
  width: 1px;
  background: var(--campus-border);
  margin: 0 20px;
}

.menu-section {
  margin-bottom: 30px;
}

.menu-title {
  font-size: 18px;
  font-weight: 600;
  margin-bottom: 16px;
  color: var(--campus-text-primary);
}

.menu-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
}

.menu-item {
  padding: 24px;
  text-align: center;
  cursor: pointer;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12px;
}

.menu-item span {
  font-size: 14px;
  color: var(--campus-text-primary);
}

.settings-list {
  padding: 0;
}

.setting-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  cursor: pointer;
  transition: background 0.3s;
  border-bottom: 1px solid var(--campus-border);
}

.setting-item:last-child {
  border-bottom: none;
}

.setting-item:hover {
  background: var(--campus-bg-hover);
}

.setting-left {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 16px;
  color: var(--campus-text-primary);
}

.logout-btn {
  width: 100%;
  margin-top: 30px;
  padding: 16px;
  font-size: 16px;
}

@media (max-width: 768px) {
  .user-header {
    flex-direction: column;
    text-align: center;
  }
  
  .menu-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .user-stats {
    flex-direction: column;
    gap: 16px;
  }
  
  .stat-divider {
    display: none;
  }
}
</style>

