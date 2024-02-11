<template>
    <div>
      <UserProfileHeader :userProfile="userProfile" />
  
      <!-- User Profile Actions Component -->
      <UserProfileActions />
  
      <!-- User Photos Component -->
      <UserProfilePhotos :userProfile="userProfile" />
    </div>
  </template>
  
  <script>
  import UserProfileHeader from "@/components/UserProfileHeader.vue";
  import UserProfileActions from "@/components/UserProfileActions.vue";
  import UserProfilePhotos from "@/components/UserProfilePhotos.vue";
  
  export default {
    components: {
      UserProfileHeader,
      UserProfileActions,
      UserProfilePhotos,
    },
    data() {
      return {
        userProfile: {},
      };
    },
    mounted() {
      this.fetchUserProfile();
    },
    methods: {
      async fetchUserProfile() {
        const userID = this.$route.params.userID;
        try {
          const response = await fetch(`http://localhost:3000/users/${userID}/profile`);
          this.userProfile = await response.json();
        } catch (error) {
          console.error("Error fetching user profile", error);
        }
      },
    },
  };
  </script>
  
  <style scoped>
  /* Add any view-specific styles here */
  </style>
  