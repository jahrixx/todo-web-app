import { writable, get, derived } from "svelte/store";
import { filterTasksBySearch } from "$lib/utils/search";

export interface Task {
  id: number;
  title: string;
  description?: string;
  completed: boolean;
}

export const tasks = writable<Task[]>([]);
export const filter = writable<'all' | 'active' | 'completed'>('all');

export const searchQuery = writable('');

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
  [tasks, filter, searchQuery],
  ([$tasks, $filter, $searchQuery]) => {
    const searchedTasks = filterTasksBySearch($tasks, $searchQuery);

    switch($filter) {
      case 'active':
        return searchedTasks.filter((task) => !task.completed);
      case 'completed':
        return searchedTasks.filter((task) => task.completed);
      default:
        return searchedTasks;
    }
  }
);

export function remainingTasks() {
  return get(tasks).filter((task) => !task.completed);
}