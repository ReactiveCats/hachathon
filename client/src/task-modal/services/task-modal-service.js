import { wrappedFetch } from '../../shared/wrapped-fetch';

class TaskModalService {
  create(task) {
    return wrappedFetch('/api/tasks', {
      method: 'POST',
      body: task,
    });
  }

  update(taskId, updatedTask) {
    return wrappedFetch(`/api/tasks/${taskId}`, {
      method: 'PUT',
      body: updatedTask,
    });
  }
}

export const taskModalService = new TaskModalService();
