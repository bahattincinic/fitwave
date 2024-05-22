import { API_BASE_URL } from './api';

export async function getUserConfig() {
  try {
    const response = await fetch(`${API_BASE_URL}/user/config`, {
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
    console.error('Error fetching user config:', error);
    throw error;
  }
}

export async function saveUserConfig(config) {
  try {
    const response = await fetch(`${API_BASE_URL}/user/config`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(config)
    });

    if (!response.ok) {
      throw new Error('Network response was not ok');
    }

    return await response.json();
  } catch (error) {
    console.error('Error saving user config:', error);
    throw error;
  }
}
