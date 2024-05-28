import { API_BASE_URL } from './api';

export const componentTypes = [
  { name: 'Table', code: 'table' },
  { name: 'Text', code: 'text' },
  { name: 'Pie Chart', code: 'pie_chart' },
  { name: 'Bar Chart', code: 'bar_chart' },
  { name: 'Line Chart', code: 'line_chart' },
];

export async function fetchComponents(dashId) {
  const endpoint = `${API_BASE_URL}/dashboards/${dashId}/components`;

  const response = await fetch(endpoint, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error('Could not fetch dashboard components');
  }

  return await response.json();
}

export async function createComponent(dashId, data) {
  const endpoint = `${API_BASE_URL}/dashboards/${dashId}/components`;

  const response = await fetch(endpoint, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  });

  if (!response.ok) {
    throw new Error('Could not create a component');
  }

  return await response.json();
}

export async function updateComponent(dashId, compId, data) {
  const endpoint = `${API_BASE_URL}/dashboards/${dashId}/components/${compId}`;

  const response = await fetch(endpoint, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  });

  if (!response.ok) {
    throw new Error('Could not create a component');
  }

  return await response.json();
}

export async function deleteComponent(dashId, compId) {
  const endpoint = `${API_BASE_URL}/dashboards/${dashId}/components/${compId}`;

  const response = await fetch(endpoint, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error('Could not delete a component');
  }
}

export async function runComponent(dashId, compId) {
  const endpoint = `${API_BASE_URL}/dashboards/${dashId}/components/${compId}`;

  const response = await fetch(endpoint, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error('Could not run component');
  }

  return await response.json();
}
