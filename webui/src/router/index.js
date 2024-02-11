import {createRouter, createWebHashHistory} from 'vue-router'

import LoginPage from '../views/LoginView.vue'
import UserStream from '../views/MainContainerView.vue'
import UserProfile from '../views/UserProfileView.vue'
//Vue.use(Router)
const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
    routes: [
        {
          path: '/',
          name: 'session',
          component: LoginPage
        },
        {
          path: '/users/:userID/stream',
          name: 'UserStream',
          component: UserStream
        },
        {
          path: '/users/:userID/profile',
          name: 'UserProfile',
          component: UserProfile
        },
        {
          path: '/logout',
          name: 'Logout',
          beforeEnter: (to, from, next) => {
            // Perform any logout logic here
            // Redirect to the LoginPage after logging out
            next({ name: 'session' });  // Assuming 'session' is the name of your LoginPage route
          }
        }

      ]

})
export default router;
