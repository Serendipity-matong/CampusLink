<template>
  <div class="orders-page campus-container">
    <h1 class="page-title">
      <el-icon :size="32"><Document /></el-icon>
      我的订单
    </h1>

    <!-- 订单类型切换 -->
    <el-tabs v-model="activeTab" class="order-tabs">
      <el-tab-pane label="全部订单" name="all"></el-tab-pane>
      <el-tab-pane label="二手交易" name="product"></el-tab-pane>
      <el-tab-pane label="跑腿任务" name="task"></el-tab-pane>
    </el-tabs>

    <!-- 订单列表 -->
    <div class="orders-list" v-if="orders.length > 0">
      <div class="order-card campus-card" v-for="order in orders" :key="order.id">
        <div class="order-header">
          <div class="order-info">
            <span class="order-number">订单号: {{ order.orderNo }}</span>
            <span class="order-time">{{ order.createTime }}</span>
          </div>
          <el-tag :type="getStatusType(order.status)">{{ order.statusText }}</el-tag>
        </div>

        <div class="order-content">
          <img :src="order.image" class="order-image" />
          <div class="order-details">
            <h3>{{ order.title }}</h3>
            <p>{{ order.description }}</p>
            <div class="order-meta">
              <span class="order-type">
                <el-icon><ShoppingBag /></el-icon>
                {{ order.type }}
              </span>
              <span class="order-price">¥{{ order.price }}</span>
            </div>
          </div>
        </div>

        <div class="order-footer">
          <div class="order-total">
            <span>总计：</span>
            <span class="total-price">¥{{ order.totalPrice }}</span>
          </div>
          <div class="order-actions">
            <el-button v-if="order.status === 'pending'" type="primary" @click="payOrder(order)">
              去支付
            </el-button>
            <el-button v-if="order.status === 'paid'" type="success" @click="confirmOrder(order)">
              确认收货
            </el-button>
            <el-button v-if="order.status === 'completed'" @click="rateOrder(order)">
              评价
            </el-button>
            <el-button v-if="order.status === 'pending'" text type="danger" @click="cancelOrder(order)">
              取消订单
            </el-button>
            <el-button text @click="viewOrderDetail(order)">
              订单详情
            </el-button>
          </div>
        </div>
      </div>
    </div>

    <!-- 空状态 -->
    <el-empty v-else description="暂无订单" :image-size="200" />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Document, ShoppingBag } from '@element-plus/icons-vue'

const activeTab = ref('all')

const orders = ref([
  {
    id: 1,
    orderNo: '202411041234567890',
    type: '二手交易',
    title: '高等数学教材（第七版）',
    description: '同济大学版，9成新',
    price: 25,
    totalPrice: 25,
    status: 'pending',
    statusText: '待支付',
    image: 'https://images.unsplash.com/photo-1543002588-bfa74002ed7e?w=200&h=200&fit=crop',
    createTime: '2024-11-04 10:30:00'
  },
  {
    id: 2,
    orderNo: '202411041234567891',
    type: '跑腿任务',
    title: '菜鸟驿站取快递到宿舍',
    description: '取件码：4526',
    price: 5,
    totalPrice: 5,
    status: 'paid',
    statusText: '进行中',
    image: 'https://images.unsplash.com/photo-1566576912321-d58ddd7a6088?w=200&h=200&fit=crop',
    createTime: '2024-11-04 09:15:00'
  },
  {
    id: 3,
    orderNo: '202411031234567892',
    type: '二手交易',
    title: '护眼台灯',
    description: '小米台灯Pro',
    price: 50,
    totalPrice: 50,
    status: 'completed',
    statusText: '已完成',
    image: 'https://images.unsplash.com/photo-1507473885765-e6ed057f782c?w=200&h=200&fit=crop',
    createTime: '2024-11-03 14:20:00'
  },
  {
    id: 4,
    orderNo: '202411021234567893',
    type: '跑腿任务',
    title: '超市代买零食清单',
    description: '薯片、可乐、饼干',
    price: 10,
    totalPrice: 10,
    status: 'completed',
    statusText: '已完成',
    image: 'https://images.unsplash.com/photo-1534723452862-4c874018d66d?w=200&h=200&fit=crop',
    createTime: '2024-11-02 16:45:00'
  }
])

const getStatusType = (status) => {
  const typeMap = {
    pending: 'warning',
    paid: 'primary',
    completed: 'success',
    cancelled: 'info'
  }
  return typeMap[status] || 'info'
}

const payOrder = (order) => {
  ElMessage.success(`正在跳转支付...订单号：${order.orderNo}`)
}

const confirmOrder = (order) => {
  ElMessageBox.confirm('确认收货后，订单将完成', '确认收货', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    order.status = 'completed'
    order.statusText = '已完成'
    ElMessage.success('已确认收货')
  })
}

const rateOrder = (order) => {
  ElMessage.info('评价功能开发中...')
}

const cancelOrder = (order) => {
  ElMessageBox.confirm('确定要取消这个订单吗？', '取消订单', {
    confirmButtonText: '确定',
    cancelButtonText: '再想想',
    type: 'warning'
  }).then(() => {
    order.status = 'cancelled'
    order.statusText = '已取消'
    ElMessage.success('订单已取消')
  })
}

const viewOrderDetail = (order) => {
  ElMessage.info(`查看订单详情：${order.orderNo}`)
}
</script>

<style scoped>
.orders-page {
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

.order-tabs {
  margin-bottom: 30px;
}

.orders-list {
  display: grid;
  gap: 20px;
}

.order-card {
  padding: 24px;
}

.order-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  padding-bottom: 16px;
  border-bottom: 1px solid var(--campus-border);
}

.order-info {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.order-number {
  font-size: 14px;
  font-weight: 600;
  color: var(--campus-text-primary);
}

.order-time {
  font-size: 12px;
  color: var(--campus-text-secondary);
}

.order-content {
  display: flex;
  gap: 20px;
  margin-bottom: 20px;
}

.order-image {
  width: 120px;
  height: 120px;
  object-fit: cover;
  border-radius: var(--campus-radius-sm);
  flex-shrink: 0;
}

.order-details {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.order-details h3 {
  font-size: 18px;
  font-weight: 600;
  margin-bottom: 8px;
  color: var(--campus-text-primary);
}

.order-details p {
  color: var(--campus-text-secondary);
  font-size: 14px;
  margin-bottom: 12px;
}

.order-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.order-type {
  display: flex;
  align-items: center;
  gap: 4px;
  color: var(--campus-text-secondary);
  font-size: 14px;
}

.order-price {
  font-size: 20px;
  font-weight: 700;
  color: #E74C3C;
}

.order-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 16px;
  border-top: 1px solid var(--campus-border);
}

.order-total {
  display: flex;
  align-items: center;
  gap: 8px;
}

.total-price {
  font-size: 24px;
  font-weight: 700;
  color: #E74C3C;
}

.order-actions {
  display: flex;
  gap: 12px;
}

@media (max-width: 768px) {
  .order-content {
    flex-direction: column;
  }
  
  .order-image {
    width: 100%;
    height: 200px;
  }
  
  .order-footer {
    flex-direction: column;
    gap: 16px;
    align-items: stretch;
  }
  
  .order-actions {
    flex-direction: column;
  }
}
</style>

