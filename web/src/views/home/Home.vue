<!--
 * @Author: yaoyaoyu
 * @Date: 2021-05-28 09:59:28
-->
<template>
  <div>
    <header class="header">
        <bi-home-header/>
          <div class="logo">
            <img src="~@/assets/img/logo.png" alt="">
          </div>
    </header>
    <div class="container">
      <aside class="aside">
          <home-aside/>
      </aside>
      <section class="section1">

          <home-main-content/>

      </section>
      <section class="section2">
          <home-right-aside/>
      </section>
    </div>
  </div>
</template>

<script>

import HomeHeader from '@/components/content/HomeHeader.vue'
import HomeAside from '@/views/home/childComps/HomeAside.vue'
import HomeRightAside from '@/views/home/childComps/HomeRightAside.vue'
import HomeMainContent from '@/views/home/childComps/HomeMainContent.vue'
import BiHomeHeader from '@/components/content/BiHomeHeader.vue'


export default {
  name:'Home',
  components:{
    HomeHeader,
    BiHomeHeader,
    HomeAside,
    HomeRightAside,
    HomeMainContent
    
  },
  data(){
    return{
      
    }
  },
}
</script>

<style lang="less" scoped>
  .header{
    width: 100%;
  }
  .container{
    background-color: #f6f6f6;
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    justify-content: space-around;
  }

  .aside{ 
    padding: 20px;
    flex-grow: 1;
    flex-shrink: 0;
  }
  .section1{
    padding: 20px;
    flex-grow: 2
  }
  .section2{
    padding: 20px;
    flex-grow: 1;

  }


  .logo {
    width: 300px;
    height: 135px;
    z-index: 999;
    position: absolute;
    top: 20px;
    left: 30px;
    overflow: hidden;
  }



</style>