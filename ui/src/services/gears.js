import { getApiBaseURL } from './api';

export async function fetchGears(page) {
  const endpoint = `${getApiBaseURL()}/gears?page=${page}`;

  const response = await fetch(endpoint, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error('Could not fetch gears');
  }

  return await response.json();
}

export async function getGear(id) {
  const endpoint = `${getApiBaseURL()}/gears/${id}`;

  const response = await fetch(endpoint, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error('Could not get gear');
  }

  return await response.json();
}
