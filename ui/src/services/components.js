import { makeRequest } from './api';

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

export const dynamicQueryOptions = [
  {
    option: 'Today',
    value: '{{.Today}}',
    example: "start_date = '{{.Today}}'",
  },
  {
    option: 'Yesterday',
    value: '{{.Yesterday}}',
    example: "start_date = '{{.Yesterday}}'",
  },
  {
    option: 'This Week Start',
    value: '{{.ThisWeekStart}}',
    example: "start_date >= '{{.ThisWeekStart}}'",
  },
  {
    option: 'This Week End',
    value: '{{.ThisWeekEnd}}',
    example: "start_date <= '{{.ThisWeekEnd}}'",
  },
  {
    option: 'Last Week Start',
    value: '{{.LastWeekStart}}',
    example: "start_date >= '{{.LastWeekStart}}'",
  },
  {
    option: 'Last Week End',
    value: '{{.LastWeekEnd}}',
    example: "start_date <= '{{.LastWeekEnd}}'",
  },
  {
    option: 'This Month Start',
    value: '{{.ThisMonthStart}}',
    example: "start_date >= '{{.ThisMonthStart}}'",
  },
  {
    option: 'This Month End',
    value: '{{.ThisMonthEnd}}',
    example: "start_date <= '{{.ThisMonthEnd}}'",
  },
  {
    option: 'Last Month Start',
    value: '{{.LastMonthStart}}',
    example: "start_date >= '{{.LastMonthStart}}'",
  },
  {
    option: 'Last Month End',
    value: '{{.LastMonthEnd}}',
    example: "start_date <= '{{.LastMonthEnd}}'",
  },
];

export async function fetchComponents(accessToken, dashId) {
  return await makeRequest({
    endpoint: `/dashboards/${dashId}/components`,
    json: true,
    error: 'Could not fetch dashboard components',
    accessToken,
  });
}

export async function createComponent(accessToken, dashId, data) {
  return await makeRequest({
    endpoint: `/dashboards/${dashId}/components`,
    method: 'POST',
    body: data,
    json: true,
    error: 'Could not create a component',
    accessToken,
  });
}

export async function updateComponent(accessToken, dashId, compId, data) {
  return await makeRequest({
    endpoint: `/dashboards/${dashId}/components/${compId}`,
    method: 'PUT',
    body: data,
    json: true,
    error: 'Could not create a component',
    accessToken,
  });
}

export async function deleteComponent(accessToken, dashId, compId) {
  return await makeRequest({
    endpoint: `/dashboards/${dashId}/components/${compId}`,
    method: 'DELETE',
    error: 'Could not delete a component',
    accessToken,
  });
}

export async function runComponent(accessToken, dashId, compId) {
  return await makeRequest({
    endpoint: `/dashboards/${dashId}/components/${compId}/run`,
    method: 'POST',
    json: true,
    error: 'Could not run component',
    accessToken,
  });
}
