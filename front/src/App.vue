<template>
  <div id="app" class="wrapper">
    fdfdfd
    <Sidebar></Sidebar>
    <MainContents></MainContents>
    bbbb
  </div>
</template>

<script>
import Sidebar from './components/Layout/Sidebar.vue'
import MainContents from './components/Layout/MainContents.vue'

export default {
  name: 'App',
  components: {
    MainContents,
    Sidebar
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  max-width: 1000px;
  width:80%;
  margin-left:auto;
  margin-right:auto;
}

.wrapper{
  display: flex;
  flex: 1 1 auto;
}

</style>
