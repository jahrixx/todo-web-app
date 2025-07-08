import { writable, get, derived } from "svelte/store";

export interface Task {
  id: number;
  title: string;
  description?: string;
  completed: boolean;
}

export const tasks = writable<Task[]>([]);
export const filter = writable<'all' | 'active' | 'completed'>('all');

export const addTask = (task: Task) => {
  tasks.update((oldTasks) => [...oldTasks, task]);
};

export const removeTask = (taskToRemove: Task) => {
  tasks.update((oldTasks) =>
    oldTasks.filter((task) => task.id !== taskToRemove.id)
  );
};

export const updateTask = (taskId: number, updatedTask: Task) => {
  tasks.update((tasks) =>
    tasks.map((task) => (task.id === taskId ? updatedTask : task))
  );
};

export const filteredTasks = derived(
  [tasks, filter],
  ([$tasks, $filter]) => {
    switch($filter) {
      case 'active':
        return $tasks.filter((task) => !task.completed);
      case 'completed':
        return $tasks.filter((task) => task.completed);
      default:
        return $tasks;
    }
  }
);

export function remainingTasks() {
  return get(tasks).filter((task) => !task.completed);
}