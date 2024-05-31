import { createRouter, createWebHistory } from 'vue-router';
import Cookies from 'js-cookie';
import LoginPage from '@/pages/LoginPage';
import HomePage from '@/pages/HomePage';
import SettingsPage from '@/pages/SettingsPage';
import DashboardDetail from '@/pages/DashboardDetail';
import AthletesPage from '@/pages/AthletesPage';
import ActivitiesPage from '@/pages/ActivitiesPage';
import GearsPage from '@/pages/GearsPage';
import { useUserStore } from '@/store/user';
import { getUserMe } from '@/services/user';

const routes = [
  { path: '/', name: 'Home', component: HomePage },
  {
    path: '/app/dashboard/:id',
    name: 'DashboardDetail',
    component: DashboardDetail,
  },
  { path: '/app/login', name: 'Login', component: LoginPage },
  { path: '/app/settings', name: 'Settings', component: SettingsPage },
  { path: '/app/activities', name: 'Activities', component: ActivitiesPage },
  { path: '/app/athletes', name: 'Athletes', component: AthletesPage },
  { path: '/app/gears', name: 'Gears', component: GearsPage },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach(async (to, from, next) => {
  const userStore = useUserStore();
  const accessToken = Cookies.get('accessToken');

  if (userStore.accessToken === '' && accessToken) {
    try {
      const resp = await getUserMe(accessToken);
      userStore.setUser(accessToken, resp.athlete);
    } catch {
      userStore.logout();
    }
  }
  next();
});

export default router;
