<!--
 * @Author: yaoyaoyu
 * @Date: 2021-05-30 09:26:38
-->
<template>
  <div class="">
      <div class="section over-hide">
        <div class="container">
          <!-- 流行区域 -->
          <span class="fly"></span>
          <span class="fly1"></span>
          <span class="fly2"></span>
          <span class="fly3"></span>
          <span class="fly4"></span>
          <span class="fly5"></span>
          <span class="fly6"></span>

          <div class="row full-height justify-content-center">
            <div class="text-center align-self-center py-5">

              <div class="section text-center py-5 py-md-0">
                      <input class="pricing" type="checkbox" id="pricing" name="pricing"/>
                      <label for="pricing"><span class="block-diff">登录<span class="float-right">注册</span></span></label>
                <div class="card-3d-wrap mx-auto">
                  <div class="card-3d-wrapper">

                    <div class="card-front">
                      <div class="pricing-wrap">
                        <!-- logo显示区域 -->
                      </div>
                      <!-- 登录 -->
                      <div class="login-box">
                           <div class="user-box">
                              <input type="text" name="" required="">
                              <label>username</label>
                            </div>
                            <div class="user-box">
                              <input type="password" name="" required="">
                              <label>password</label>
                            </div>
                            <a href="#" id="submit">
                              <span></span>
                              <span></span>
                              <span></span>
                              <span></span>
                              Submit
                            </a>
                      </div>
                    </div>
                    <div class="card-back">
                      <div class="pricing-wrap">
                        
                      </div>
                      <!-- 注册区域 -->
                      <div class="login-box">
                           <div class="user-box">
                              <input type="text" name="" required="">
                              <label>Username</label>
                            </div>
                            <div class="user-box">
                              <input type="password" name="" required="">
                              <label>Password</label>
                            </div>
                            <div class="user-box">
                              <input type="password" name="" required="">
                              <label>ReRassword</label>
                            </div>
                            <a href="#" id="submit">
                              <span></span>
                              <span></span>
                              <span></span>
                              <span></span>
                              register
                            </a>
                      </div>
                    </div>

                  </div>
                </div>

              </div>
            </div>
          </div>
        </div>
      </div>
  </div>
</template>

<script>
export default {
  name:'Login',
  components:{

  },
  data(){
    return{
      // 登录表单信息
      loginFrom :{
        username:'',
        password:'',
      }
      
    }
  },
  methods:{

    // 登录
    submitLogin(){

    },

    // 注册
    submitRegister(){

    }

  }

}
</script>

<style lang="less" scoped>

:root{
  font-size: 20px;
}
.container{
  background: url('../../assets/img/bg3.png') no-repeat;
  background-position: center;
  background-size: cover;

	font-family: 'Poppins', sans-serif;
	font-weight: 500;
	font-size: 15px;
	line-height: 1.7;
	color: #102770;
	background-color: #95cff7;
	overflow-x: hidden;
}

.over-hide {
  overflow: hidden;
}
.size-22{
  font-size: 22px;
}
.section{
  position: relative;
  width: 100%;
  z-index: 1;
  display: block;
}
.full-height{
  min-height: 100vh;
}
[type="checkbox"]:checked,
[type="checkbox"]:not(:checked){
  position: absolute;
  left: -9999px;
}
.pricing:checked + label,
.pricing:not(:checked) + label{
  z-index: 100;
  position: relative;
  top: 100px;
  display: block;
  text-align: center;
  width: 220px;
  height: 44px;
  border-radius: 4px;
  padding: 0;
  // margin: 0 auto;
  margin-top: 120px;
  left: 30%;
  transform: translateX(-50%);
  cursor: pointer;
  font-family: 'Poppins', sans-serif;
  font-weight: 600;
  text-transform: uppercase;
  font-size: 14px;
  letter-spacing: 1px;
  line-height: 44px;
  padding: 0 25px;
  // padding-right: 27px;
  overflow: hidden;
  color: #fff;
  text-align: center;

}
.pricing:checked + label:before,
.pricing:not(:checked) + label:before{
  position: absolute;
  content: '';
  z-index: -2;
  background-color: #5170d3;
  width: 100%;
  height: 100%;
  display: block;
  top: 0;
  left: 0;
}
.pricing:checked + label:after,
.pricing:not(:checked) + label:after{
  position: absolute;
  content: '';
  z-index: -1;
  background-color: pink;
  width: 50%;
  height: 44px;
  display: block;
  top: 0px;
  left: 0px;
  border-radius: 2px;
  transition: left 300ms linear;
}
.pricing:checked + label:after {
  left: 50%;
}
.block-diff {
  display: block;

}

.float-right{
  margin-left: 50%;
}
.card-3d-wrap {
  position: relative;
  left: 30%;
  transform: translateX(-50%);
  width: 400px;
  max-width: calc(100% - 20px);
  height: 400px;
  -webkit-transform-style: preserve-3d;
  transform-style: preserve-3d;
  perspective: 1000px;
  margin-top: 90px;
  
  opacity:0.7;

}
.card-3d-wrapper {
  width: 100%;
  height: 100%;
  position:absolute;    
  top: 0;
  left: 0;  
  -webkit-transform-style: preserve-3d;
  transform-style: preserve-3d;
  transition: transform 700ms 400ms ease-out; 
  box-shadow: 3px 2px 45px -10px rgba(17,119,248,1);
  -webkit-box-shadow: 3px 2px 45px -10px rgba(17,119,248,1);
  -moz-box-shadow: 3px 2px 45px -10px rgba(17,119,248,1);
}
.card-front, .card-back {
  width: 100%;
  height: 100%;
  background-color: rgb(253, 251, 251);
  position: absolute;
  border-radius: 10px;
  left: 0;
  top: 0;
  -webkit-transform-style: preserve-3d;
  transform-style: preserve-3d;
  -webkit-backface-visibility: hidden;
  -moz-backface-visibility: hidden;
  -o-backface-visibility: hidden;
  backface-visibility: hidden;
  box-shadow: 0 12px 35px 0 rgba(16, 39, 112,.07);
}
.card-back {
  transform: rotateY(180deg);
}
.pricing:checked ~ .card-3d-wrap .card-3d-wrapper {
  transform: rotateY(180deg);
  transition: transform 700ms 400ms ease-out; 
}
.pricing-wrap{
  position: relative;
  height: 80px;
  width: 100%;
  display: block;
  z-index: 1;
  background: url('../../assets/img/logo.png') 0px -20px;
  -webkit-transform-style: preserve-3d;
  transform-style: preserve-3d;
  -webkit-backface-visibility: hidden;
  -moz-backface-visibility: hidden;
  -o-backface-visibility: hidden;
  backface-visibility: hidden;
}

@keyframes border-transform{
    0%,100% { border-radius: 63% 37% 54% 46% / 55% 48% 52% 45%; } 
  14% { border-radius: 40% 60% 54% 46% / 49% 60% 40% 51%; } 
  28% { border-radius: 54% 46% 38% 62% / 49% 70% 30% 51%; } 
  42% { border-radius: 61% 39% 55% 45% / 61% 38% 62% 39%; } 
  56% { border-radius: 61% 39% 67% 33% / 70% 50% 50% 30%; } 
  70% { border-radius: 50% 50% 34% 66% / 56% 68% 32% 44%; } 
  84% { border-radius: 46% 54% 50% 50% / 35% 61% 39% 65%; } 
}



.login-box{
  padding: 35px;
  
}

.login-box .user-box {
  position: relative;
}

.login-box .user-box input {
  width: 100%;
  padding: 10px 0;
  font-size: 16px;
  color: rgb(99, 13, 13);
  margin-bottom: 30px;
  border: none;
  border-bottom: 1px solid #000;
  outline: none;
  background: transparent;
}
.login-box .user-box label {
  position: absolute;
  top:0;
  left: 0;
  padding: 10px 0;
  font-size: 16px;
  color: #000;
  pointer-events: none;
  transition: .5s;
}

.login-box .user-box input:focus ~ label,
.login-box .user-box input:valid ~ label {
  top: -20px;
  left: 0;
  color: #6e8cee;
  font-size: 12px;
}

.login-box a {
  position: relative;
  display: block;
  color: #03e9f4;
  font-size: 16px;
  text-decoration: none;
  text-transform: uppercase;
  overflow: hidden;
  transition: .5s;
  letter-spacing: 4px
}

.login-box a:hover {
  background: #03e9f4;
  color: #fff;
  border-radius: 5px;
  box-shadow: 0 0 5px #03e9f4,
              0 0 25px #03e9f4,
              0 0 50px #03e9f4,
              0 0 100px #03e9f4;
} 
.login-box a span {
  position: absolute;
  display: block;
}

.login-box:hover a span:nth-child(1){
  top: 0;
  left: -100%;
  width: 100%;
  height: 2px;
  background: linear-gradient(90deg, transparent, #03e9f4);
  animation: btn-anim1 1s linear infinite;
}


@keyframes btn-anim1 {
  0% {
    left: -100%;
  }
  50%,100% {
    left: 100%;
  }
}

.login-box:hover a span:nth-child(2) {
  top: -100%;
  right: 0;
  width: 2px;
  height: 100%;
  background: linear-gradient(180deg, transparent, #03e9f4);
  animation: btn-anim2 1s linear infinite;
  animation-delay: .25s
}

@keyframes btn-anim2 {
  0% {
    top: -100%;
  }
  50%,100% {
    top: 100%;
  }
}

.login-box:hover a span:nth-child(3) {
  bottom: 0;
  right: -100%;
  width: 100%;
  height: 2px;
  background: linear-gradient(270deg, transparent, #03e9f4);
  animation: btn-anim3 1s linear infinite;
  animation-delay: .5s
}

@keyframes btn-anim3 {
  0% {
    right: -100%;
  }
  50%,100% {
    right: 100%;
  }
}

.login-box:hover a span:nth-child(4) {
  bottom: -100%;
  left: 0;
  width: 2px;
  height: 100%;
  background: linear-gradient(360deg, transparent, #03e9f4);
  animation: btn-anim4 1s linear infinite;
  animation-delay: .75s
}

@keyframes btn-anim4 {
  0% {
    bottom: -100%;
  }
  50%,100% {
    bottom: 100%;
  }
}

 .fly{
  position: absolute;
  display: block;

}

.container:hover .fly{
  bottom: 100%;
  left: 100%;
  width: 2px;
  height: 50%;
  transform: rotateZ(225deg);
  background: linear-gradient(360deg, transparent, #03e9f4);
  animation: fly 5s infinite;
  animation-delay: 4.25s
}

@keyframes fly {
  0% {
      transform: rotateZ(225deg);
  }
  50%,100% {
    bottom: -50%;
    left: -10%;
    transform: rotateZ(225deg);
  }
}




.container:hover .fly1{
  position: absolute;
  display: block;

  bottom: 130%;
  left: 130%;
  width: 2px;
  height: 50%;
  transform: rotateZ(225deg);
  background: linear-gradient(360deg, transparent, #03e9f4);
  animation: fly1 5s infinite;
  animation-delay: 3.25s
}

@keyframes fly1 {
  0% {
      transform: rotateZ(225deg);
  }
  50%,100% {
    bottom: -50%;
    left: 10%;
    transform: rotateZ(225deg);
  }
}

.container:hover .fly2{
  position: absolute;
  display: block;

  bottom: 130%;
  left: 90%;
  width: 2px;
  height: 50%;
  transform: rotateZ(225deg);
  background: linear-gradient(360deg, transparent, #03e9f4);
  animation: fly2 5s infinite;
  animation-delay: 2.95s
}

@keyframes fly2 {
  0% {
      transform: rotateZ(225deg);
  }
  50%,100% {
    bottom: -50%;
    left: -30%;
    transform: rotateZ(225deg);
  }
}
.container:hover  .fly3{
  position: absolute;
  display: block;

  bottom: 130%;
  left: 40%;
  width: 2px;
  height: 50%;
  transform: rotateZ(225deg);
  background: linear-gradient(360deg, transparent, #03e9f4);
  animation: fly3 5s infinite;
  animation-delay: 1.84s
}

@keyframes fly3 {
  0% {
      transform: rotateZ(225deg);
  }
  50%,100% {
    bottom: -70%;
    left: -60%;
    transform: rotateZ(225deg);
  }
}



.container:hover .fly4{
  position: absolute;
  display: block;

  bottom: 130%;
  left: 120%;
  width: 2px;
  height: 50%;
  transform: rotateZ(225deg);
  background: linear-gradient(360deg, transparent, #03e9f4);
  animation: fly4 5s infinite;
  animation-delay: 1.50s
}

@keyframes fly4 {
  0% {
      transform: rotateZ(225deg);
  }
  50%,100% {
    bottom: -70%;
    left: -50%;
    transform: rotateZ(225deg);
  }
}

.container:hover .fly5{
  position: absolute;
  display: block;

  bottom: 130%;
  left: 100%;
  width: 2px;
  height: 50%;
  transform: rotateZ(225deg);
  background: linear-gradient(360deg, transparent, #03e9f4);
  animation: fly5 5s infinite;
  animation-delay: .10s
}

@keyframes fly5 {
  0% {
      transform: rotateZ(225deg);
  }
  50%,100% {
    bottom: -70%;
    left: -30%;
    transform: rotateZ(225deg);
  }
}



.container:hover .fly6{
  position: absolute;
  display: block;

  bottom: 130%;
  left: 180%;
  width: 2px;
  height: 50%;
  transform: rotateZ(225deg);
  background: linear-gradient(360deg, transparent, #03e9f4);
  animation: fly6 5s infinite;
  animation-delay: 3.40s
}

@keyframes fly6 {
  0% {
      transform: rotateZ(225deg);
  }
  50%,100% {
    bottom: -50%;
    left: 60%;
    transform: rotateZ(225deg);
  }
}
</style> 