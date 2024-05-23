import { API_BASE_URL } from './api';

export async function fetchActivities(page) {
  const response = await fetch(`${API_BASE_URL}/activities?page=${page}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
  });

  if (!response.ok) {
    throw new Error('Could not fetch activities');
  }

  return await response.json();
}
