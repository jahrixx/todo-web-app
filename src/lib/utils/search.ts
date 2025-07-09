import type { Task } from "$lib/stores/tasks";

export function filterTasksBySearch(tasks: Task[], query: string) {
    if(!query.trim()) return tasks;

    const searchTerm = query.toLowerCase();
    return tasks.filter((task) => task.title.toLowerCase().includes(searchTerm));
}