<template>
    <div v-if="showComments" class="view-comments">
      <div v-if="comments && comments.length">
        <div v-for="comment in comments" :key="comment.commentID" class="comment">
          <p>{{ comment.commentText }}</p>
          <span v-if="comment.userID === userId" @click="deleteComment(comment)" class="delete-comment-btn">
            &times;
          </span>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      showComments: Boolean,
      comments: Array,
      userId: String,
    },
    methods: {
      deleteComment(comment) {
        const userID = this.$route.params.userID;
      const imageID = post.imageID;
      const commentID = comment.commentID;
      const url = `http://localhost:3000/users/${userID}/image/${imageID}/comment/${commentID}`;

      try {
        fetch(url, { method: 'DELETE' })
          .then(() => {
            // Remove the deleted comment from the post
            post.comments = post.comments.filter(c => c.commentID !== commentID);
          })
          .catch(error => console.error(error));
      } catch (error) {
        console.error(error);
      }
      },
    },
  };
  </script>
  
  <style scoped>
  /* Add your styles here */
  .view-comments {
    margin-top: 10px;
  }
  
  .comment {
    border-bottom: 1px solid #ddd;
    padding: 8px 0;
    display: flex;
    justify-content: space-between;
  }
  
  .delete-comment-btn {
    cursor: pointer;
    color: red;
  }
  </style>
  