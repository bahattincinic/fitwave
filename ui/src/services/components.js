import { API_BASE_URL } from './api';

export const componentTypeEnum = {
  table: 'table',
  pieChart: 'pie_chart',
  barChart: 'bar_chart',
  lineChart: 'line_chart',
};

export const componentTypes = [
  { name: 'Table', code: componentTypeEnum.table },
  { name: 'Pie Chart', code: componentTypeEnum.pieChart },
  { name: 'Bar Chart', code: componentTypeEnum.barChart },
  { name: 'Line Chart', code: componentTypeEnum.lineChart },
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
  const endpoint = `${API_BASE_URL}/dashboards/${dashId}/components/${compId}/run`;

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
