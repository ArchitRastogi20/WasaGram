<template>
    <div class="sidebar">
      <router-link :to="`/users/${userID}/profile`">My Profile</router-link>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      userID: String,
    },
  };
  </script>
  
  <style scoped>
  /* Add your styles here */
  .sidebar {
    width: 25%;
    float: left;
    padding: 20px;
    box-sizing: border-box;
    background-color: #89232e;
  }
  
  .sidebar a {
    color: black;
  }
  </style>
  