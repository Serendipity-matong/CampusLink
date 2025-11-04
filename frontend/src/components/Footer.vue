<template>
  <footer class="campus-footer">
    <div class="footer-content">
      <div class="footer-section">
        <h3>CampusLink 校园互联</h3>
        <p>让校园生活更便捷</p>
      </div>
      
      <div class="footer-section">
        <h4>功能服务</h4>
        <ul>
          <li><a href="/products">二手交易</a></li>
          <li><a href="/tasks">跑腿代购</a></li>
          <li><a href="/orders">订单管理</a></li>
        </ul>
      </div>
      
      <div class="footer-section">
        <h4>关于我们</h4>
        <ul>
          <li><a href="#about">平台介绍</a></li>
          <li><a href="#contact">联系我们</a></li>
          <li><a href="#help">帮助中心</a></li>
        </ul>
      </div>
      
      <div class="footer-section">
        <h4>关注我们</h4>
        <div class="social-icons">
          <el-icon :size="24"><ChatDotRound /></el-icon>
          <el-icon :size="24"><Message /></el-icon>
          <el-icon :size="24"><Share /></el-icon>
        </div>
      </div>
    </div>
    
    <div class="footer-bottom">
      <p>© 2024 CampusLink 校园互联平台 | Go微服务架构</p>
    </div>
  </footer>
</template>

<script setup>
import { ChatDotRound, Message, Share } from '@element-plus/icons-vue'
</script>

<style scoped>
.campus-footer {
  background: #2C3E50;
  color: #ECF0F1;
  padding: 40px 0 20px;
  margin-top: auto;
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 30px;
  padding: 0 20px 30px;
}

.footer-section h3 {
  background: linear-gradient(135deg, #1E88E5 0%, #66BB6A 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  margin-bottom: 10px;
}

.footer-section h4 {
  color: #ECF0F1;
  font-size: 16px;
  margin-bottom: 15px;
}

.footer-section ul {
  list-style: none;
  padding: 0;
}

.footer-section li {
  margin-bottom: 8px;
}

.footer-section a {
  color: #BDC3C7;
  text-decoration: none;
  transition: color 0.3s;
}

.footer-section a:hover {
  color: #42A5F5;
}

.social-icons {
  display: flex;
  gap: 15px;
  color: #42A5F5;
}

.footer-bottom {
  text-align: center;
  padding-top: 20px;
  border-top: 1px solid #34495E;
  color: #95A5A6;
  font-size: 14px;
}

@media (max-width: 768px) {
  .footer-content {
    grid-template-columns: 1fr;
    gap: 20px;
  }
}
</style>

