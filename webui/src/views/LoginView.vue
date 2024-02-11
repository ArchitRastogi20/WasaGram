<template>
    <div class="login-container">
      <LoginBox @makeAPICall="makeAPICall" />
      <UserDataContainer v-if="userData" :userData="userData" />
    </div>
  </template>
  
  <script>
  import LoginBox from "@/components/LoginBox.vue";
  import UserDataContainer from "@/components/UserDataContainer.vue";
  
  export default {
    name: "LoginView",
    data() {
      return {
        userData: null,
      };
    },
    methods: {
      async makeAPICall(payload) {
        try {
          const response = await fetch("http://localhost:3000/session", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(payload),
          });
  
          if (!response.ok) {
            const errorMessage = `HTTP error! Status: ${response.status}`;
            console.error(errorMessage);
            throw new Error(errorMessage);
          }
  
          try {
            const responseData = await response.json();
            this.userData = responseData;
            console.warn(responseData);
          } catch (jsonError) {
            console.error("Error parsing JSON response", jsonError);
          }
          // Store the userID as the token in local storage
        localStorage.setItem('token', userID);
  
          // Navigate to "/users" after getting the data
          this.$router.push(`/users/${payload.userID}/stream`);
        } catch (error) {
          console.error(error);
        }
      },
    },
  };
  </script>
  