<template>
    <div class="user-data-container">
      <h2>User Data:</h2>
      <p>UserID: {{ userData.userID }}</p>
      <p>Username: {{ userData.nickname }}</p>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      userData: Object,
    },
  };
  </script>
  
  <style scoped>
  .user-data-container {
    margin-top: 20px;
    color: white;
  }
  </style>
  