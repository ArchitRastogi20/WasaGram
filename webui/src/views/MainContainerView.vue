<template>
    <div class="main-container">
      <Sidebar :userID="$route.params.userID" />
      <MainContent :userPosts="userPosts" @updateComments="updateComments" />
    </div>
  </template>
  
  <script>
  import Sidebar from '@/components/Sidebar.vue';
  import MainContent from '@/components/MainContent.vue';
  
  export default {
    components: {
      Sidebar,
      MainContent,
    },
    data() {
      return {
        userPosts: [],
      };
    },
    created() {
      this.fetchUserStream();
    },
    methods: {
      fetchUserStream() {
        const userID = this.$route.params.userID;
        const url = `http://localhost:3000/users/${userID}/stream`;
  
        try {
          fetch(url)
            .then((response) => {
              if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
              }
              return response.json();
            })
            .then((data) => {
              this.userPosts = data.map((post) => ({
                ...post,
                liked: false,
                newComment: '',
                comments: [],
              }));
            })
            .catch((error) => {
              console.error(error);
            });
        } catch (error) {
          console.error(error);
        }
      },
  
      updateComments({ postId, comments }) {
        // Find the post by postId and update its comments
        const updatedPosts = this.userPosts.map((post) => {
          if (post.imageID === postId) {
            return { ...post, comments };
          }
          return post;
        });
        this.userPosts = updatedPosts;
      },
    },
  };
  </script>
  
  <style scoped>
  /* Add your styles here */
  </style>
  