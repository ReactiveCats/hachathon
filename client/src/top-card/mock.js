export const mockTask = (overrides) => ({
  title: 'Title',
  description: 'Description',
  priority: 'very_low',
  complexity: 'very_high',
  hardDeadline: new Date(),
  softDeadline: new Date(),
  status: '',
  ...overrides,
});