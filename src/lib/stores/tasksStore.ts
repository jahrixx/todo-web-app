import { getTasks, createTask, updateTask, deleteTask } from "$lib/api/api";
import { token } from "$lib/stores/auth";
import { get, writable, derived } from "svelte/store";
import { filterTasksBySearch } from "$lib/utils/search"; 
import type { Task } from "$lib/types/task";
  
export const loading = writable(true);
export const tasks = writable<Task[]>([]);
export const filter = writable<'all' | 'active' | 'completed'>('all');
export const searchQuery = writable('');

export const filteredTasks = derived(
  [tasks, filter, searchQuery],
  ([$tasks, $filter, $searchQuery]) => {
    const searchedTasks = filterTasksBySearch($tasks, $searchQuery);

    switch ($filter) {
      case 'active':
        return searchedTasks.filter((task) => !task.completed);
      case 'completed':
        return searchedTasks.filter((task) => task.completed);
      default:
        return searchedTasks;
    }
  }
);

export async function fetchAllTasks() {
  loading.set(true);
  const authToken = get(token);
    if(!authToken) return;
    
    try {
      const data = await getTasks(authToken);
      tasks.set(data);
      console.log("Fetched tasks:", data);
    } catch (error) {
      console.error("Failed to fetch tasks:", error);
      tasks.set([]);
    } finally {
      loading.set(false);
    }
}

export async function createNewTaskRemote(task: {title: string, description?: string}): Promise<void> {
  const authToken = get(token);
    if(!authToken) throw new Error("No token found!");

    await createTask(authToken, task);
    await fetchAllTasks();
}

export async function updateTaskRemote(task: Task): Promise<void> {
  const authToken = get(token);
    if(!authToken) throw new Error("No token found!");

    await updateTask(authToken, task);
    await fetchAllTasks();
}

export async function deleteTaskRemote(id: string): Promise<void> {
  const authToken = get(token);
    if(!authToken) throw new Error("No token found!");

    await deleteTask(authToken, id);
    await fetchAllTasks();
}
