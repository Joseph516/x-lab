<!--
 * @Author: yaoyaoyu
 * @Date: 2021-05-28 10:43:55
-->
<template>
    <div class="aside-container">
        <user-info-card style="margin-left: 'auto',margin-right: 'auto'"></user-info-card>
    </div>

</template>

<script>
import UserInfoCard from '@/components/content/UserInfoCard.vue'

export default {
  name:'HomeAside',
  components:{
    UserInfoCard
  }
}
</script>

<style lang="less" scoped>
  .aside-container{
    width: 100%;
    padding-top: 20px;
    text-align: center;
    text-align: -webkit-center;
    background-color: rgb(232, 242, 252);
    border-radius: 30px;
    height: 1000px;

  }

</style>