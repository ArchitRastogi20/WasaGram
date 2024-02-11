<template>
	<div id="app">
	  <router-view />
	</div>
  </template>
  
  <script setup>
  import { onMounted } from 'vue';
  import { useRouter } from 'vue-router';
  import axios from 'axios';
  
  const router = useRouter();
  
  // Check if the user is logged in and has a valid token
  const checkAuthentication = () => {
	const token = localStorage.getItem('token');
	return token ? true : false;
  };
  
  // Set the Bearer token for Axios requests
  const setBearerToken = () => {
	const token = localStorage.getItem('token');
	if (token) {
	  axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
	}
  };
  
  onMounted(() => {
	// Check authentication on app mount
	if (!checkAuthentication()) {
	  router.replace('/login');
	}
  
	// Set Bearer token for subsequent requests
	setBearerToken();
  });
  
  </script>
  
  <script>
  import LoginPage from './views/LoginView.vue'; // Import the LoginView component
  
  export default {
	components: {
	  LoginPage, // Register the LoginPage component
	},
  };
  </script>
  
  <style>
  #app {
	font-family: 'Avenir', Helvetica, Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	text-align: center;
	color: #2c3e50;
	margin-top: 60px;
  }
  </style>
  