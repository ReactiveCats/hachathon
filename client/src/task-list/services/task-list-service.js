import { wrappedFetch } from '../../shared/wrapped-fetch';

class TaskListService {
  loadItems() {
    return wrappedFetch('/api/task');
  }
}

export const taskListService = new TaskListService();
