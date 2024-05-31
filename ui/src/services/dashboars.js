import { getApiBaseURL } from './api';

export async function fetchDashboards() {
  const endpoint = `${getApiBaseURL()}/dashboards`;

  const response = await fetch(endpoint, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error('Could not fetch dashboards');
  }

  return await response.json();
}

export async function getDashboard(id) {
  const endpoint = `${getApiBaseURL()}/dashboards/${id}`;

  const response = await fetch(endpoint, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error('Could not get dashboard');
  }

  return await response.json();
}

export async function createDashboard(data) {
  const endpoint = `${getApiBaseURL()}/dashboards`;

  const response = await fetch(endpoint, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  });

  if (!response.ok) {
    throw new Error('Could not create a dashboard');
  }

  return await response.json();
}

export async function updateDashboard(id, data) {
  const endpoint = `${getApiBaseURL()}/dashboards/${id}`;

  const response = await fetch(endpoint, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  });

  if (!response.ok) {
    throw new Error('Could not update a dashboard');
  }

  return await response.json();
}

export async function deleteDashboard(id) {
  const endpoint = `${getApiBaseURL()}/dashboards/${id}`;

  const response = await fetch(endpoint, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error('Could not delete a dashboard');
  }
}

export async function runDashboard(id) {
  const endpoint = `${getApiBaseURL()}/dashboards/${id}/run`;

  const response = await fetch(endpoint, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error('Could not run dashboard');
  }

  return await response.json();
}
