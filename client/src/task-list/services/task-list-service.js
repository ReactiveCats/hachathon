import { wrappedFetch } from '../../shared/wrapped-fetch';

class TaskListService {
  loadItems() {
    return wrappedFetch('/api/tasks');
  }
}

export const taskListService = new TaskListService();
