<template>
    <div v-if="userProfile.photos && userProfile.photos.length > 0" class="user-photos">
      <h2>Photos</h2>
      <div v-for="photo in userProfile.photos || []" :key="photo.imageID" class="photo-container">
        <img :src="getPhotoUrl(photo.photo_data)" alt="User Photo" class="user-photo" />
        <p>Likes: {{ photo.likeCount }}</p>
        <p>Comments: {{ photo.commentCount }}</p>
        <button class="photo-button" @click="deleteImage(photo.imageID)">Delete Image</button>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      userProfile: Object,
    },
    methods: {
      getPhotoUrl(photoData) {
        // Replace with your actual logic for converting photoData to a URL
        return `data:image/jpeg;base64,${photoData}`;
      },
  
      async deleteImage(imageID) {
        const userID = this.$route.params.userID;
        try {
          await fetch(`http://localhost:8000/users/${userID}/image/${imageID}`, {
            method: 'DELETE',
          });
          alert("Image deleted successfully!");
          // Optionally, you can fetch the updated user profile after deleting the image
          this.fetchUserProfile();
        } catch (error) {
          console.error("Error deleting image", error);
        }
      },
    },
  };
  </script>
  
  <style scoped>
  /* Styles for user photos */
  .user-photos {
    margin-top: 20px;
  }
  
  /* Styles for photo containers, images, and buttons */
  .photo-container {
    margin-bottom: 20px;
  }
  
  .user-photo {
    max-width: 100%;
    height: auto;
  }
  
  .photo-button {
    background-color: #dc3545; /* Bootstrap danger color for delete button */
    color: #fff;
    padding: 5px;
    border: none;
    cursor: pointer;
  }
  </style>
  