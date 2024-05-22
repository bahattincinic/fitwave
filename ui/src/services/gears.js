import { API_BASE_URL } from './api';

export async function fetchGears(page) {
  try {
    const response = await fetch(`${API_BASE_URL}/gears?page=${page}`, {
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
    console.error('Error fetching gears:', error);
    throw error;
  }
}
