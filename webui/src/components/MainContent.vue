<template>
    <div class="main-content">
      <h1>User Stream</h1>
      <PostContainer v-for="post in userPosts" :key="post.imageID" :post="post" />
    </div>
  </template>
  
  <script>
  import PostContainer from '@/components/PostContainer.vue';
  
  export default {
    components: {
      PostContainer,
    },
    props: {
      userPosts: Array,
    },
  };
  </script>
  
  <style scoped>
  /* Add your styles here */
  .main-content {
    width: 75%;
    float: left;
    padding: 20px;
    box-sizing: border-box;
  }
  </style>
  