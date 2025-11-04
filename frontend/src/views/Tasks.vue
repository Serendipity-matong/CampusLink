<template>
  <div class="tasks-page campus-container">
    <h1 class="page-title">
      <el-icon :size="32"><Bicycle /></el-icon>
      跑腿代购
    </h1>

    <!-- 任务类型切换 -->
    <el-tabs v-model="activeTab" @tab-click="handleTabClick" class="task-tabs">
      <el-tab-pane label="全部任务" name="all"></el-tab-pane>
      <el-tab-pane label="代取快递" name="delivery"></el-tab-pane>
      <el-tab-pane label="代购" name="shopping"></el-tab-pane>
      <el-tab-pane label="代课" name="class"></el-tab-pane>
      <el-tab-pane label="其他" name="other"></el-tab-pane>
    </el-tabs>

    <!-- 任务列表 -->
    <div class="tasks-list">
      <div class="task-card campus-card" v-for="task in tasks" :key="task.id">
        <div class="task-header">
          <div class="task-title-section">
            <span class="campus-tag">{{ task.type }}</span>
            <h3>{{ task.title }}</h3>
          </div>
          <div class="task-reward">¥{{ task.reward }}</div>
        </div>

        <p class="task-description">{{ task.description }}</p>

        <div class="task-details">
          <div class="task-detail-item">
            <el-icon><Clock /></el-icon>
            <span>{{ task.time }}</span>
          </div>
          <div class="task-detail-item">
            <el-icon><Location /></el-icon>
            <span>{{ task.location }}</span>
          </div>
          <div class="task-detail-item">
            <el-icon><User /></el-icon>
            <span>{{ task.publisher }}</span>
          </div>
        </div>

        <div class="task-footer">
          <div class="publisher-info">
            <img :src="task.avatar" class="campus-avatar" />
            <div class="publisher-stats">
              <span class="rating">
                <el-icon><Star /></el-icon>
                {{ task.rating }}
              </span>
              <span class="orders">已完成{{ task.completedOrders }}单</span>
            </div>
          </div>
          <el-button type="primary" @click="acceptTask(task)">
            接受任务
          </el-button>
        </div>
      </div>
    </div>

    <!-- 发布任务按钮 -->
    <el-button type="primary" size="large" circle class="fab-button" @click="publishTask">
      <el-icon :size="24"><Plus /></el-icon>
    </el-button>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { 
  Bicycle, 
  Clock, 
  Location, 
  User, 
  Star, 
  Plus 
} from '@element-plus/icons-vue'

const activeTab = ref('all')

const tasks = ref([
  {
    id: 1,
    type: '代取快递',
    title: '菜鸟驿站取快递到宿舍',
    description: '需要帮忙从菜鸟驿站代取快递到东区12栋506，重量不超过5kg，取件码：4526',
    reward: 5,
    time: '2小时内',
    location: '东区菜鸟驿站',
    publisher: '张同学',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=task1',
    rating: 4.8,
    completedOrders: 23
  },
  {
    id: 2,
    type: '代购',
    title: '超市代买零食清单',
    description: '需要买一些零食，清单：薯片2包、可乐2瓶、饼干1盒，超市在西区，送到西区8栋302',
    reward: 10,
    time: '今天下午6点前',
    location: '西区超市',
    publisher: '李同学',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=task2',
    rating: 4.9,
    completedOrders: 45
  },
  {
    id: 3,
    type: '代课',
    title: '周三上午高数代点名',
    description: '有急事无法上课，需要帮忙代点名，教学楼A205，早上8点开始',
    reward: 20,
    time: '周三 8:00 AM',
    location: '教学楼A205',
    publisher: '王同学',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=task3',
    rating: 4.7,
    completedOrders: 12
  },
  {
    id: 4,
    type: '代取快递',
    title: '圆通快递取件',
    description: '圆通快递代取到北区5栋，取件码：9876，快递较大，需要骑车',
    reward: 8,
    time: '今天下午',
    location: '北区圆通自提点',
    publisher: '陈同学',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=task4',
    rating: 4.6,
    completedOrders: 18
  },
  {
    id: 5,
    type: '其他',
    title: '图书馆占座',
    description: '明天早上7点帮忙去图书馆3楼占个座位，期末复习需要',
    reward: 15,
    time: '明天早上7:00',
    location: '图书馆3楼',
    publisher: '赵同学',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=task5',
    rating: 4.9,
    completedOrders: 31
  }
])

const handleTabClick = () => {
  ElMessage.success('切换任务类型')
}

const acceptTask = (task) => {
  ElMessage.success(`已接受任务：${task.title}`)
}

const publishTask = () => {
  ElMessage.info('发布任务功能开发中...')
}
</script>

<style scoped>
.tasks-page {
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

.task-tabs {
  margin-bottom: 30px;
}

.tasks-list {
  display: grid;
  gap: 20px;
}

.task-card {
  padding: 24px;
}

.task-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 16px;
}

.task-title-section {
  flex: 1;
}

.task-title-section h3 {
  font-size: 20px;
  font-weight: 600;
  margin-top: 8px;
  color: var(--campus-text-primary);
}

.task-reward {
  font-size: 28px;
  font-weight: 700;
  color: #E74C3C;
}

.task-description {
  color: var(--campus-text-secondary);
  line-height: 1.8;
  margin-bottom: 20px;
}

.task-details {
  display: flex;
  flex-wrap: wrap;
  gap: 24px;
  margin-bottom: 20px;
  padding: 16px;
  background: var(--campus-bg-secondary);
  border-radius: var(--campus-radius-sm);
}

.task-detail-item {
  display: flex;
  align-items: center;
  gap: 6px;
  color: var(--campus-text-secondary);
  font-size: 14px;
}

.task-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 20px;
  border-top: 1px solid var(--campus-border);
}

.publisher-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.publisher-stats {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.rating {
  display: flex;
  align-items: center;
  gap: 4px;
  color: #FFA726;
  font-weight: 600;
}

.orders {
  font-size: 12px;
  color: var(--campus-text-secondary);
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
  .task-header {
    flex-direction: column;
    gap: 12px;
  }
  
  .task-details {
    flex-direction: column;
    gap: 12px;
  }
  
  .task-footer {
    flex-direction: column;
    gap: 16px;
    align-items: stretch;
  }
  
  .fab-button {
    bottom: 60px;
    right: 20px;
  }
}
</style>

