import { writable, get, derived } from "svelte/store";

export interface Task {
  id: number;
  title: string;
  description?: string;
  completed: boolean;
}

export const tasks = writable<Task[]>([]);

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

export const filteredTasks = (filter: "all" | "active" | "completed") => {
  return derived(tasks, $tasks => {
    switch(filter) {
      case "active": return $tasks.filter(t => !t.completed);
      case "completed": return $tasks.filter(t => t.completed);
      default: return $tasks;
    }
  });
};

export function remainingTasks() {
  return get(tasks).filter((task) => !task.completed);
}