<template>
    <div class="profile-header">
      <h1>{{ userProfile.username }}</h1>
      <p>{{ userProfile.firstname }} {{ userProfile.lastname }}</p>
      <p>Email: {{ userProfile.email }}</p>
      <p>Followers: {{ userProfile.followersCount }}</p>
      <p>Following: {{ userProfile.followingCount }}</p>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      userProfile: Object,
    },
  };
  </script>
  
  <style scoped>
  /* Styles for profile header */
  .profile-header {
    background-color: #89232e;
    color: #fff;
    padding: 20px;
    text-align: center;
  }
  </style>
  