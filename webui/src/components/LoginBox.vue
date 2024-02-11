<template>
    <div class="login-box">
      <img src="@/assets/wasagram-logo.png" alt="Wasagram Logo" class="logo" />
      <h1>Login Page</h1>
  
      <div class="input-group">
        <label for="userID">UserID:</label>
        <input type="text" id="userID" v-model="payload.userID" />
      </div>
  
      <div class="input-group">
        <label for="nickname">Username:</label>
        <input type="text" id="nickname" v-model="payload.nickname" />
      </div>
  
      <button @click="submitForm">Submit</button>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        payload: {
          userID: "",
          nickname: "",
        },
      };
    },
    methods: {
      submitForm() {
        this.$emit("makeAPICall", { ...this.payload });
      },
    },
  };
  </script>
  
  <style scoped>
  .login-box {
    background-color: white;
    padding: 20px;
    border-radius: 8px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    text-align: center;
  }
  
  .logo {
    width: 100px; /* Adjust as needed */
    height: auto;
    margin-bottom: 20px;
  }
  
  .input-group {
    margin-bottom: 15px;
  }
  
  button {
    background-color: #89232e;
    color: white;
    padding: 10px 20px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  </style>
  