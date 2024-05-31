import { getApiBaseURL } from './api';

export async function fetchAthletes(page) {
  const endpoint = `${getApiBaseURL()}/athletes?page=${page}`;

  const response = await fetch(endpoint, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error('Could not fetch athletes');
  }

  return await response.json();
}

export async function getAthlete(id) {
  const endpoint = `${getApiBaseURL()}/athletes/${id}`;

  const response = await fetch(endpoint, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error('Could not get athlete');
  }

  return await response.json();
}
