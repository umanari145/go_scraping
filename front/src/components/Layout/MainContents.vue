<template>
  <div id="maincontents">
      aaa
      <ThreadList></ThreadList>
  </div>
</template>

<script>

import ThreadList from '../Thread/list.vue'

export default {
  name: 'MainContents',
  components:{
      ThreadList
  }
}
</script>

<style>
#maincontents{
    width:80%;
}

</style>
