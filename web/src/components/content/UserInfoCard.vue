<!--
 * @Author: yaoyaoyu
 * @Date: 2021-05-28 13:02:43
-->
<template>
    <div class="userInfoCard">
      <!-- 用户头像 -->
      <div class="userImg">
        <img src="@/assets/img/userimgdemo.jpg" alt="">
      </div>
      <!-- 用户信息下 -->
      <div class="userInfo">
        <div class="username">{{userInfo.username}} <i style="cursor:pointer" class="el-icon-s-promotion
" @click="logout"></i> </div>
        <ul>
          <li><i class="el-icon-phone-outline"> &nbsp;&nbsp;18835130952</i></li>
          <li><i class="el-icon-message">&nbsp;&nbsp;xblifebetter@gmail.com</i></li>
        </ul>
      </div>
      <!-- 档案 -->
      <div>
        <el-button ><i class="el-icon-document-checked"></i>档案</el-button>
      </div>
    </div>
</template>

<script>
export default {
  name:'UserInfoCard',
  data(){
    return{
      userInfo:{
        username:'Yao Xiangbo',
        email:'ylifebetter@163.com',
      }
    }
  },
  created(){
      
  },
  methods:{
    getUserInfo(){
      this.userInfo.username = window.sessionStorage.getItem('loginedUsername');
    },

    // 登出
    logout(){
      window.sessionStorage.removeItem('token');
      window.sessionStorage.removeItem('loginedUsername');
      this.$router.push('/login');
    }
  }
}
</script>

<style lang="less" scoped>
  ul,li{
    margin: 0;
    padding: 0;
    list-style: none;
  }

  .userInfoCard {
    padding-top: 20px;
    display: flex;
    flex-flow: row wrap;
    width: 240px;
    height: 260px;
    background-color: rgb(198, 226, 255);
    border-radius: 20px;
    justify-content: center;
    // align-items: center;
    
    .userImg{
      widows: 100px;
      height: 100px;
      img{
        width: 100px;
        height: 100px;
        border-radius: 50%;
        object-fit:cover;
      }
    }
    .userInfo{
      >.username{
        color: #606266;
        text-align: center;
        font-size: 20px;
      };
      ul{
        li{
          margin-top: 5px;
          color: #909399;
        }
      }
    }
    .el-button{
      border-radius: 8px;
      background-color: rgb(246, 250, 255);
      color: rgb(121, 187, 255);
      border: 1px solid rgb(121, 187, 255);
    }
     .el-button:hover{
       color: rgb(252, 246, 246);
       background-color: rgb(147, 175, 236);
     }

  }
</style>