import { createRouter, createWebHistory } from 'vue-router';
import HomePage from '@/pages/HomePage';
import SettingsPage from '@/pages/SettingsPage';
import ActivitiesPage from '@/pages/ActivitiesPage';
import GearsPage from '@/pages/GearsPage';

const routes = [
  { path: '/', name: 'Home', component: HomePage },
  { path: '/settings', name: 'Settings', component: SettingsPage },
  { path: '/activities', name: 'Activities', component: ActivitiesPage },
  { path: '/gears', name: 'Gears', component: GearsPage }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
