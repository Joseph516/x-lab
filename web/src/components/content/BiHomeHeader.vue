<!--
 * @Author: yaoyaoyu
 * @Date: 2021-05-28 13:22:03
-->
<template>
<header class="container">

  <div class="view">
    <img src="https://assets.codepen.io/2002878/bilibili-winter-view-1.jpg" class="morning" alt="">
    <img src="https://assets.codepen.io/2002878/bilibili-winter-view-2.jpg" class="afternoon" alt="">
    <video autoplay loop muted class="evening">
      <source src="https://assets.codepen.io/2002878/bilibili-winter-view-3.webm" type="video/webm" />
    </video>
    <img src="https://assets.codepen.io/2002878/bilibili-winter-view-3-snow.png" class="window-cover" alt="">
  </div>
  
  <div class="tree">
    <img src="https://assets.codepen.io/2002878/bilibili-winter-tree-1.png" class="morning" alt="">
    <img src="https://assets.codepen.io/2002878/bilibili-winter-tree-2.png" class="afternoon" alt="">
    <img src="https://assets.codepen.io/2002878/bilibili-winter-tree-3.png" class="evening" alt="">
  </div>

</header>

</template>

<script>


export default {
  name:'Header',
  mounted(){
    this.init(); 
  },
  created(){

  },
  methods:{
    init(){
        let startingPoint
        const header = document.getElementsByClassName('container')[0]

        header.addEventListener('mouseenter', (e) => {
          startingPoint = e.clientX
          header.classList.add('moving')
        })

        header.addEventListener('mouseout', (e) => {
          header.classList.remove('moving')
          header.style.setProperty('--percentage', 0.5)
        })

        header.addEventListener('mousemove', (e) => {
          let percentage = (e.clientX - startingPoint) / window.outerWidth + 0.5
          
          header.style.setProperty('--percentage', percentage)
        })
    }
  }
}
</script>

<style lang="less" scoped>

.container {
  height: 160px;
  position: relative;
  overflow: hidden;
  
  --percentage: 0.5;
}

.container .view, .container .tree {
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  display: flex;
  justify-content: center;
  align-items: center;
}

.container img, .container video {
  position: absolute;
  display: block;
  width: 120%;
  height: 100%;
  object-fit: cover;
}

.container .morning {
  z-index: 20;
  opacity: calc(1 - (var(--percentage) - 0.25) / 0.25);
}

.container .afternoon {
  z-index: 10;
  opacity: calc(1 - (var(--percentage) - 0.5) / 0.5);
}

.container .view {
  transform: translatex(calc(var(--percentage) * 100px));
}

.container .tree {
  transform: translatex(calc(var(--percentage) * 50px));
  filter: blur(3px);
}

.container .view,
.container .tree,
.container .morning,
.container .afternoon {
  transition: .2s all ease-in;
}

.container.moving .view,
.container.moving .tree,
.container.moving .morning,
.container.moving .afternoon {
  transition: none;
}

.container .window-cover {
  opacity: calc((var(--percentage) - 0.9) / 0.1);
}



</style>