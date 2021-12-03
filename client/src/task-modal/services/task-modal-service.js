import { wrappedFetch } from '../../shared/wrapped-fetch';

class TaskModalService {
  update(taskId, updatedTask) {
    return wrappedFetch(`/api/tasks/${taskId}`, {
      method: 'PUT',
      body: updatedTask,
    });
  }
}

export const taskModalService = new TaskModalService();
