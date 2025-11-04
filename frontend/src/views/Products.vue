<template>
  <div class="products-page campus-container">
    <h1 class="page-title">
      <el-icon :size="32"><ShoppingBag /></el-icon>
      二手市场
    </h1>

    <!-- 筛选区域 -->
    <div class="filter-section campus-card">
      <el-row :gutter="20">
        <el-col :xs="24" :sm="8" :md="6">
          <el-select v-model="filters.category" placeholder="商品分类" @change="applyFilters">
            <el-option label="全部分类" value="" />
            <el-option label="教材" value="textbook" />
            <el-option label="电子产品" value="electronics" />
            <el-option label="生活用品" value="daily" />
            <el-option label="交通工具" value="vehicle" />
          </el-select>
        </el-col>
        <el-col :xs="24" :sm="8" :md="6">
          <el-select v-model="filters.condition" placeholder="成色" @change="applyFilters">
            <el-option label="全新" value="new" />
            <el-option label="几乎全新" value="like_new" />
            <el-option label="良好" value="good" />
            <el-option label="一般" value="fair" />
          </el-select>
        </el-col>
        <el-col :xs="24" :sm="8" :md="12">
          <el-input v-model="filters.keyword" placeholder="搜索商品..." @keyup.enter="applyFilters">
            <template #prefix>
              <el-icon><Search /></el-icon>
            </template>
          </el-input>
        </el-col>
      </el-row>
    </div>

    <!-- 商品列表 -->
    <div class="products-grid">
      <div class="product-card campus-card" v-for="product in products" :key="product.id">
        <img :src="product.image" :alt="product.name" class="product-image" />
        <div class="product-info">
          <div class="product-header">
            <span class="campus-tag">{{ product.category }}</span>
            <span class="condition-tag" :class="getConditionClass(product.condition)">
              {{ getConditionText(product.condition) }}
            </span>
          </div>
          <h3>{{ product.name }}</h3>
          <p class="product-desc">{{ product.description }}</p>
          <div class="product-footer">
            <div class="price-section">
              <span class="price">¥{{ product.price }}</span>
              <span class="original-price" v-if="product.originalPrice">¥{{ product.originalPrice }}</span>
            </div>
            <div class="seller-info">
              <img :src="product.sellerAvatar" class="campus-avatar" />
              <span>{{ product.sellerName }}</span>
            </div>
          </div>
          <el-button type="primary" class="contact-btn" @click="contactSeller(product)">
            <el-icon><ChatDotRound /></el-icon>
            联系卖家
          </el-button>
        </div>
      </div>
    </div>

    <!-- 发布商品按钮 -->
    <el-button type="primary" size="large" circle class="fab-button" @click="publishProduct">
      <el-icon :size="24"><Plus /></el-icon>
    </el-button>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { ShoppingBag, Search, ChatDotRound, Plus } from '@element-plus/icons-vue'

const filters = ref({
  category: '',
  condition: '',
  keyword: ''
})

const products = ref([
  {
    id: 1,
    name: '高等数学教材（第七版）',
    description: '同济大学版，9成新，有少量笔记',
    price: 25,
    originalPrice: 45,
    category: '教材',
    condition: 'good',
    image: 'https://images.unsplash.com/photo-1543002588-bfa74002ed7e?w=400&h=300&fit=crop',
    sellerName: '张同学',
    sellerAvatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Zhang'
  },
  {
    id: 2,
    name: '自行车（捷安特）',
    description: '买了一年，车况良好，即将毕业低价转让',
    price: 200,
    originalPrice: 500,
    category: '交通工具',
    condition: 'good',
    image: 'https://images.unsplash.com/photo-1485965120184-e220f721d03e?w=400&h=300&fit=crop',
    sellerName: '李同学',
    sellerAvatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Li'
  },
  {
    id: 3,
    name: 'ThinkPad X1 Carbon',
    description: 'i7处理器，16G内存，512G固态，办公学习利器',
    price: 2500,
    originalPrice: 5000,
    category: '电子产品',
    condition: 'like_new',
    image: 'https://images.unsplash.com/photo-1496181133206-80ce9b88a853?w=400&h=300&fit=crop',
    sellerName: '王同学',
    sellerAvatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Wang'
  },
  {
    id: 4,
    name: '护眼台灯',
    description: '小米台灯Pro，可调节色温和亮度',
    price: 50,
    originalPrice: 100,
    category: '生活用品',
    condition: 'good',
    image: 'https://images.unsplash.com/photo-1507473885765-e6ed057f782c?w=400&h=300&fit=crop',
    sellerName: '陈同学',
    sellerAvatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Chen'
  },
  {
    id: 5,
    name: 'iPad 2021 (第9代)',
    description: '64G WiFi版，配原装笔，有保护壳和膜',
    price: 1800,
    originalPrice: 2600,
    category: '电子产品',
    condition: 'like_new',
    image: 'https://images.unsplash.com/photo-1544244015-0df4b3ffc6b0?w=400&h=300&fit=crop',
    sellerName: '刘同学',
    sellerAvatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Liu'
  },
  {
    id: 6,
    name: '线性代数教材',
    description: '清华大学出版社，全新未使用',
    price: 20,
    originalPrice: 35,
    category: '教材',
    condition: 'new',
    image: 'https://images.unsplash.com/photo-1532012197267-da84d127e765?w=400&h=300&fit=crop',
    sellerName: '赵同学',
    sellerAvatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Zhao'
  }
])

const getConditionText = (condition) => {
  const map = {
    new: '全新',
    like_new: '几乎全新',
    good: '良好',
    fair: '一般'
  }
  return map[condition] || condition
}

const getConditionClass = (condition) => {
  return `condition-${condition}`
}

const applyFilters = () => {
  ElMessage.success('筛选条件已应用')
}

const contactSeller = (product) => {
  ElMessage.success(`正在联系 ${product.sellerName}...`)
}

const publishProduct = () => {
  ElMessage.info('发布商品功能开发中...')
}
</script>

<style scoped>
.products-page {
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

.filter-section {
  margin-bottom: 30px;
  padding: 20px;
}

.products-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 24px;
}

.product-card {
  cursor: pointer;
  overflow: hidden;
  padding: 0;
  display: flex;
  flex-direction: column;
}

.product-image {
  width: 100%;
  height: 240px;
  object-fit: cover;
}

.product-info {
  padding: 20px;
  flex: 1;
  display: flex;
  flex-direction: column;
}

.product-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.condition-tag {
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
}

.condition-new {
  background: #E8F5E9;
  color: #66BB6A;
}

.condition-like_new {
  background: #E3F2FD;
  color: #42A5F5;
}

.condition-good {
  background: #FFF9C4;
  color: #FFA726;
}

.condition-fair {
  background: #F5F5F5;
  color: #757575;
}

.product-info h3 {
  font-size: 18px;
  margin-bottom: 8px;
  color: var(--campus-text-primary);
  font-weight: 600;
}

.product-desc {
  color: var(--campus-text-secondary);
  font-size: 14px;
  margin-bottom: 16px;
  line-height: 1.6;
  flex: 1;
}

.product-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  padding-top: 16px;
  border-top: 1px solid var(--campus-border);
}

.price-section {
  display: flex;
  align-items: baseline;
  gap: 8px;
}

.price {
  font-size: 24px;
  font-weight: 700;
  color: #E74C3C;
}

.original-price {
  font-size: 14px;
  color: var(--campus-text-light);
  text-decoration: line-through;
}

.seller-info {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  color: var(--campus-text-secondary);
}

.contact-btn {
  width: 100%;
}

.fab-button {
  position: fixed;
  bottom: 80px;
  right: 40px;
  width: 64px;
  height: 64px;
  box-shadow: var(--campus-shadow-lg);
}

@media (max-width: 768px) {
  .products-grid {
    grid-template-columns: 1fr;
  }
  
  .fab-button {
    bottom: 60px;
    right: 20px;
  }
}
</style>

