import { API_BASE_URL } from './api';

export async function fetchDashboards() {
  const response = await fetch(`${API_BASE_URL}/dashboards`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
  });

  if (!response.ok) {
    throw new Error('Network response was not ok');
  }

  return await response.json();
}
