import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import LoginBox from './components/LoginBox.vue'
import MainContent from './components/MainContent.vue'
import PostContainer from './components/PostContainer.vue'
import Sidebar from './components/Sidebar.vue'
import UserDataContainer from './components/UserDataContainer.vue'
import UserProfileActions from './components/UserProfileActions.vue'
import UserProfileHeader from './components/UserProfileHeader.vue'
import UserProfilePhotos from './components/UserProfilePhotos.vue'
import ViewComments from './components/ViewComments.vue'


import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("LoginBox", LoginBox)
app.component("MainContent",MainContent)
app.component("PostContainer",PostContainer)
app.component("Sidebar",Sidebar)
app.component("UserDataContainer",UserDataContainer)
app.component("UserProfileActions",UserProfileActions)
app.component("UserProfileHeader",UserProfileHeader)
app.component("UserProfilePhotos",UserProfilePhotos)
app.component("ViewComments",ViewComments)


app.use(router)
app.mount('#app')
