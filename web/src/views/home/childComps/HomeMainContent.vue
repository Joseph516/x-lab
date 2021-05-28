<!--
 * @Author: yaoyaoyu
 * @Date: 2021-05-28 13:06:57
-->
<template>
  <div>
    <div class="home-main-container">
      <not-full-calendar/>
    </div>
  </div>
</template>

<script>
import NotFullCalendar from '@/components/common/NotFullCalendar.vue'
export default {
  name:'HomeMainContent',
  components:{
    NotFullCalendar
  }
}
</script>

<style lang="less" scoped>
  .home-main-container{
    min-width: 700px;
    width: 100%;
    height: 1200px;
    background-color: #e8f2fc;
    border-radius: 30px;
  }

</style>