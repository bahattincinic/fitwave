import { createRouter, createWebHistory } from 'vue-router';
import Cookies from 'js-cookie';
import LoginPage from '@/pages/LoginPage';
import HomePage from '@/pages/HomePage';
import SettingsPage from '@/pages/SettingsPage';
import AthletesPage from '@/pages/AthletesPage';
import ActivitiesPage from '@/pages/ActivitiesPage';
import GearsPage from '@/pages/GearsPage';
import { useUserStore } from '@/store/user';
import { getUserMe } from "@/services/user";

const routes = [
  { path: '/', name: 'Home', component: HomePage },
  { path: '/login', name: 'Login', component: LoginPage },
  { path: '/settings', name: 'Settings', component: SettingsPage },
  { path: '/activities', name: 'Activities', component: ActivitiesPage },
  { path: '/athletes', name: 'Athletes', component: AthletesPage },
  { path: '/gears', name: 'Gears', component: GearsPage }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

router.beforeEach(async (to, from, next) => {
  const userStore = useUserStore();
  const accessToken = Cookies.get('accessToken');

  if (userStore.accessToken === '' && accessToken) {
    try {
      const resp = await getUserMe(accessToken);
      userStore.setAccessToken(accessToken);
      userStore.setUser(resp.athlete);
    } catch {
      userStore.setAccessToken('');
      Cookies.remove('accessToken');
    }
  }
  next();
});

export default router;
