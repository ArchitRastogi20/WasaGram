<template>
    <div class="post-container">
      <div class="post-header">
        <img :src="getImageUrl(post.photo_data)" alt="Post Image" class="profile-picture" />
        <div class="post-info">
          <p class="timestamp">{{ formatTimestamp(post.timestamp) }}</p>
          <p class="nickname">Username</p>  
        </div>
      </div>
  
      <img :src="getImageUrl(post.photo_data)" alt="Post Image" class="post-image" />
  
      <div class="post-actions">
        <div class="like-dislike">
          <button @click.prevent="toggleLike(post, 'like')" :class="{ active: post.liked }">
            Like
          </button>
          <p class="like-count">{{ post.likeCount }}</p>
          <span class="like-message" v-if="post.likeMessage">{{ post.likeMessage }}</span>
          <button @click.prevent="toggleLike(post, 'dislike')" :class="{ active: !post.liked }">
            Dislike
          </button>
          <span class="dislike-message" v-if="post.dislikeMessage">{{ post.dislikeMessage }}</span>
        </div>
  
        <div class="comment-section">
          <input v-model="post.newComment" type="text" placeholder="Add a comment" />
          <button @click="addComment(post)">Comment</button>
          <p class="comment-count">{{ post.comments.length }} comments</p>
        </div>
  
        <ViewComments v-if="post.showComments" :userId="$route.params.userID" :imageId="post.imageID" />
  
        <button @click="toggleShowComments(post)" :disabled="isLastClickedPost(post)">
          Show All Comments
        </button>
  
        <div v-if="post.showComments" class="comments-container">
          <div v-if="post.comments && post.comments.length">
            <div v-for="comment in post.comments" :key="comment.commentID" class="comment">
              <p>{{ comment.commentText }}</p>
              <span v-if="comment.userID === $route.params.userID" @click="deleteComment(post, comment)" class="delete-comment-btn">
                &times;
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import ViewComments from '@/components/ViewComments.vue';
  
  export default {
    components: {
      ViewComments,
    },
    props: {
      post: Object,
    },
    methods: {
    getImageUrl(photo_data) {
        if (photo_data) {
        return `data:image/jpeg;base64,${photo_data}`;
      }
      return 'placeholder-image-url';
    },
    isLastClickedPost(post) {
      // Check if the current post is the last clicked post
      return this.lastClickedPost === post;
    },

    formatTimestamp(timestamp) {
        const date = new Date(timestamp);
      const day = date.getDate().toString().padStart(2, '0');
      const month = (date.getMonth() + 1).toString().padStart(2, '0');
      const year = date.getFullYear().toString().slice(2);
      return `${day}/${month}/${year}`;
    },

    async toggleLike(post, action) {
        const userID = this.$route.params.userID;
      const urlCheck = `http://localhost:3000/users/${userID}/image/${post.imageID}/like/stream`;
      const urlLike = `http://localhost:3000/like/users/${userID}/image/${post.imageID}`;

      try {
        const response = await fetch(urlCheck);
        if (!response.ok) {
          throw new Error(`Failed to fetch like status for post ${post.imageID}`);
        }

        const likeExists = await response.json();

        if (action === 'like') {
          if (likeExists) {
            post.likeMessage = 'Image already liked';
          } else {
            await fetch(urlLike, { method: 'PUT' });
            post.likeCount++;
            post.liked = true;
            post.likeMessage = 'Image liked';
            post.dislikeMessage = ''; // Clear the dislike message
          }
        } else if (action === 'dislike') {
          if (likeExists) {
            await fetch(urlLike, { method: 'DELETE' });
            post.likeCount--;
            post.liked = false;
            post.dislikeMessage = 'Image disliked';
            post.likeMessage = ''; // Clear the like message
          } else {
            post.dislikeMessage = 'Image not liked';
          }
        }
      } catch (error) {
        console.error(error);
      }
    },

    async addComment(post) {
        const userID = this.$route.params.userID;
      const url = `http://localhost:3000/users/${userID}/image/${post.imageID}/comment`;
      const commentText = post.newComment;

      try {
        if (commentText) {
          await fetch(url, {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({ text: commentText }),
          });
          post.newComment = '';
          this.$router.go();
        }
      } catch (error) {
        console.error(error);
      }
    },

    deleteComment(post, comment) {
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

toggleShowComments(post) {
      // Toggle the 'showAllComments' property for the clicked post
  post.showAllComments = !post.showAllComments;

// Update lastClickedPost to the current post
this.lastClickedPost = post.showAllComments ? post : null;

// If 'showAllComments' is true, fetch comments for the post
if (post.showAllComments) {
  this.fetchComments(post);
}
    },
    fetchComments(post) {
        const userID = this.$route.params.userID // Replace with the actual user ID
        const imageID = post.imageID; // Replace with the actual image ID
        console.log(userID,imageID)
        const url = `http://localhost:3000/users/${userID}/image/${imageID}/comments/stream`;
  
        // Make the GET request to fetch comments
        fetch(url)
          .then((response) => response.json())
          .then((data) => {
            this.comments = data;
            console.log(this.comments);
          })
          
          .catch((error) => {
            console.error('Error fetching comments:', error);
          });
      },


    isLastClickedPost(post) {
       // Check if the current post is the last clicked post
       return this.lastClickedPost === post;
    },
    },
  };
  </script>
  
  <style scoped>
  /* Add your styles here */
  .post-container {
    border: 1px solid #ddd;
    margin-bottom: 20px;
    padding: 15px;
    box-sizing: border-box;
    background-color: #fff;
  }
  
  .post-header {
    display: flex;
    align-items: center;
  }
  
  .profile-picture {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    margin-right: 10px;
  }
  
  .post-info {
    flex-grow: 1;
  }
  
  .timestamp {
    font-size: 12px;
    color: #555;
  }
  
  .username {
    font-size: 16px;
    font-weight: bold;
  }
  
  .post-image {
    max-width: 100%;
    margin-top: 10px;
  }
  
  .post-actions {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-top: 10px;
  }
  
  .like-dislike button {
    padding: 8px;
    cursor: pointer;
    background-color: #3498db;
    color: #fff;
    border: none;
    border-radius: 5px;
  }
  
  .like-dislike button.active {
    background-color: #e74c3c;
  }
  
  .like-count {
    margin: 0 10px;
    font-weight: bold;
  }
  
  .like-message,
  .dislike-message {
    margin-left: 10px;
    color: #e74c3c;
  }
  
  .comment-section input {
    padding: 8px;
    flex-grow: 1;
  }
  
  .comment-section button {
    padding: 8px;
    margin-left: 10px;
    cursor: pointer;
    background-color: #2ecc71;
    color: #fff;
    border: none;
    border-radius: 5px;
  }
  
  .comment-count {
    margin-left: 10px;
    font-size: 12px;
    color: #555;
  }
  
  .comments-container {
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
  