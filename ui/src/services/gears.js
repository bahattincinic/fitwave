import { API_BASE_URL } from './api';

export async function fetchGears(page) {
  const response = await fetch(`${API_BASE_URL}/gears?page=${page}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
  });

  if (!response.ok) {
    throw new Error('Could not fetch gears');
  }

  return await response.json();
}
