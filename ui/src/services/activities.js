import { API_BASE_URL } from './api';

export async function fetchActivities(page) {
  try {
    const response = await fetch(`${API_BASE_URL}/activities?page=${page}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    });

    if (!response.ok) {
      throw new Error('Network response was not ok');
    }

    return await response.json();
  } catch (error) {
    console.error('Error fetching activities:', error);
    throw error;
  }
}
