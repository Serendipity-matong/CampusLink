<template>
  <div class="home-page">
    <!-- 英雄区域 -->
    <section class="hero-section campus-gradient-bg">
      <div class="campus-container">
        <div class="hero-content">
          <h1 class="hero-title">
            <el-icon :size="48"><School /></el-icon>
            欢迎来到 CampusLink
          </h1>
          <p class="hero-subtitle">校园互联平台 · 让校园生活更便捷</p>
          <div class="hero-buttons">
            <el-button type="primary" size="large" round @click="goToProducts">
              <el-icon><ShoppingBag /></el-icon>
              逛逛二手市场
            </el-button>
            <el-button type="success" size="large" round @click="goToTasks">
              <el-icon><Bicycle /></el-icon>
              发布跑腿任务
            </el-button>
          </div>
        </div>
      </div>
    </section>

    <!-- 功能卡片 -->
    <section class="features-section campus-container">
      <h2 class="section-title">核心功能</h2>
      <div class="features-grid">
        <div class="feature-card campus-card" v-for="feature in features" :key="feature.title">
          <div class="feature-icon" :style="{ background: feature.color }">
            <el-icon :size="32"><component :is="feature.icon" /></el-icon>
          </div>
          <h3>{{ feature.title }}</h3>
          <p>{{ feature.description }}</p>
        </div>
      </div>
    </section>

    <!-- 热门商品 -->
    <section class="products-section campus-container">
      <div class="section-header">
        <h2 class="section-title">热门商品</h2>
        <el-button text type="primary" @click="goToProducts">查看更多 →</el-button>
      </div>
      <div class="products-grid">
        <div class="product-card campus-card" v-for="product in hotProducts" :key="product.id">
          <img :src="product.image" :alt="product.name" class="product-image" />
          <div class="product-info">
            <h4>{{ product.name }}</h4>
            <div class="product-meta">
              <span class="campus-tag">{{ product.category }}</span>
              <span class="price">¥{{ product.price }}</span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- 最新任务 -->
    <section class="tasks-section campus-container">
      <div class="section-header">
        <h2 class="section-title">最新任务</h2>
        <el-button text type="primary" @click="goToTasks">查看更多 →</el-button>
      </div>
      <div class="tasks-list">
        <div class="task-card campus-card" v-for="task in recentTasks" :key="task.id">
          <div class="task-header">
            <span class="campus-tag">{{ task.type }}</span>
            <span class="task-reward">¥{{ task.reward }}</span>
          </div>
          <h4>{{ task.title }}</h4>
          <p class="task-desc">{{ task.description }}</p>
          <div class="task-footer">
            <span class="task-time">
              <el-icon><Clock /></el-icon>
              {{ task.time }}
            </span>
            <span class="task-location">
              <el-icon><Location /></el-icon>
              {{ task.location }}
            </span>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { 
  School, 
  ShoppingBag, 
  Bicycle, 
  Document, 
  User,
  Clock,
  Location
} from '@element-plus/icons-vue'

const router = useRouter()

const features = [
  {
    icon: 'ShoppingBag',
    title: '二手交易',
    description: '买卖闲置物品，让资源循环利用',
    color: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)'
  },
  {
    icon: 'Bicycle',
    title: '跑腿代购',
    description: '帮你代取快递、代买物品',
    color: 'linear-gradient(135deg, #f093fb 0%, #f5576c 100%)'
  },
  {
    icon: 'Document',
    title: '订单管理',
    description: '实时追踪订单状态',
    color: 'linear-gradient(135deg, #4facfe 0%, #00f2fe 100%)'
  },
  {
    icon: 'User',
    title: '个人中心',
    description: '管理个人信息和交易记录',
    color: 'linear-gradient(135deg, #43e97b 0%, #38f9d7 100%)'
  }
]

const hotProducts = ref([
  { 
    id: 1, 
    name: '高等数学教材（第七版）', 
    price: 25, 
    category: '教材',
    image: 'https://images.unsplash.com/photo-1543002588-bfa74002ed7e?w=400&h=300&fit=crop'
  },
  { 
    id: 2, 
    name: '自行车（9成新）', 
    price: 200, 
    category: '交通工具',
    image: 'https://images.unsplash.com/photo-1485965120184-e220f721d03e?w=400&h=300&fit=crop'
  },
  { 
    id: 3, 
    name: 'ThinkPad 笔记本', 
    price: 2500, 
    category: '电子产品',
    image: 'https://images.unsplash.com/photo-1496181133206-80ce9b88a853?w=400&h=300&fit=crop'
  },
  { 
    id: 4, 
    name: '台灯（护眼款）', 
    price: 50, 
    category: '生活用品',
    image: 'https://images.unsplash.com/photo-1507473885765-e6ed057f782c?w=400&h=300&fit=crop'
  }
])

const recentTasks = ref([
  {
    id: 1,
    type: '代取快递',
    title: '帮忙取快递到宿舍',
    description: '菜鸟驿站代取快递到东区12栋，重量不超过5kg',
    reward: 5,
    time: '2小时内',
    location: '东区'
  },
  {
    id: 2,
    type: '代购',
    title: '超市代买零食',
    description: '需要买一些零食，清单已列好',
    reward: 10,
    time: '今天下午',
    location: '西区'
  },
  {
    id: 3,
    type: '代课',
    title: '周三上午代点名',
    description: '有事无法上课，需要帮忙代点名',
    reward: 20,
    time: '周三 8:00',
    location: '教学楼A205'
  }
])

const goToProducts = () => router.push('/products')
const goToTasks = () => router.push('/tasks')
</script>

<style scoped>
.home-page {
  flex: 1;
}

.hero-section {
  padding: 80px 20px;
  text-align: center;
  color: white;
}

.hero-content {
  max-width: 800px;
  margin: 0 auto;
}

.hero-title {
  font-size: 48px;
  font-weight: 700;
  margin-bottom: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 16px;
}

.hero-subtitle {
  font-size: 20px;
  margin-bottom: 40px;
  opacity: 0.95;
}

.hero-buttons {
  display: flex;
  gap: 20px;
  justify-content: center;
}

.features-section,
.products-section,
.tasks-section {
  padding: 60px 20px;
}

.section-title {
  font-size: 32px;
  font-weight: 700;
  margin-bottom: 40px;
  text-align: center;
  color: var(--campus-text-primary);
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30px;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 24px;
}

.feature-card {
  text-align: center;
  padding: 40px 20px;
}

.feature-icon {
  width: 80px;
  height: 80px;
  border-radius: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 20px;
  color: white;
}

.feature-card h3 {
  font-size: 20px;
  margin-bottom: 12px;
  color: var(--campus-text-primary);
}

.feature-card p {
  color: var(--campus-text-secondary);
  line-height: 1.6;
}

.products-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 24px;
}

.product-card {
  cursor: pointer;
  overflow: hidden;
  padding: 0;
}

.product-image {
  width: 100%;
  height: 200px;
  object-fit: cover;
}

.product-info {
  padding: 16px;
}

.product-info h4 {
  font-size: 16px;
  margin-bottom: 12px;
  color: var(--campus-text-primary);
}

.product-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.price {
  font-size: 20px;
  font-weight: 700;
  color: #E74C3C;
}

.tasks-list {
  display: grid;
  gap: 16px;
}

.task-card {
  cursor: pointer;
}

.task-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.task-reward {
  font-size: 20px;
  font-weight: 700;
  color: #E74C3C;
}

.task-card h4 {
  font-size: 18px;
  margin-bottom: 8px;
  color: var(--campus-text-primary);
}

.task-desc {
  color: var(--campus-text-secondary);
  margin-bottom: 12px;
}

.task-footer {
  display: flex;
  gap: 20px;
  font-size: 14px;
  color: var(--campus-text-secondary);
}

.task-time,
.task-location {
  display: flex;
  align-items: center;
  gap: 4px;
}

@media (max-width: 768px) {
  .hero-title {
    font-size: 32px;
  }
  
  .hero-buttons {
    flex-direction: column;
  }
  
  .section-header {
    flex-direction: column;
    gap: 16px;
  }
}
</style>

