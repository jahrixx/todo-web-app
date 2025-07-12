import { getTasks, createTask, updateTask, deleteTask } from "$lib/api/api";
import { token } from "$lib/stores/auth";
import { get, writable, derived } from "svelte/store";
import { filterTasksBySearch } from "$lib/utils/search"; 
import type { Task } from "$lib/types/task";

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

// export async function fetchAllTasks() {
//   if (!token || !token) return;

//   try {
//     const data = await getTasks(get(token)); // <- call your API
//     tasks.set(data); // <- update the store
//   } catch (err) {
//     console.error("Failed to fetch tasks:", err);
//     tasks.set([]);
//   }
// }

export async function fetchAllTasks() {
  const authToken = get(token);
    if(!authToken) return;
    
    try {
      const data = await getTasks(authToken);
      tasks.set(data);
      console.log("Fetched tasks:", data);
    } catch (error) {
      console.error("Failed to fetch tasks:", error);
      tasks.set([]);
    }
}

// export async function createNewTask(task: { title: string; description?: string }): Promise<void> {
//   const authToken = get(token);
//   if (!authToken) throw new Error("No token found");
//   await createTask(authToken, task);
// }

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
// export async function updateTaskRemote(task: any): Promise<void> {
//   const authToken = get(token);
//   if (!authToken) throw new Error("No token found");
//   await updateTask(authToken, task);
// }

// export async function deleteTaskRemote(id: string): Promise<void> {
//   const authToken = get(token);
//   if (!authToken) throw new Error("No token found");
//   await deleteTask(authToken, id);
// }
