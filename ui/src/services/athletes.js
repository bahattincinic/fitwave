import { API_BASE_URL } from './api';

export async function fetchAthletes(page) {
  const response = await fetch(`${API_BASE_URL}/athletes?page=${page}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
  });

  if (!response.ok) {
    throw new Error('Could not fetch athletes');
  }

  return await response.json();
}
