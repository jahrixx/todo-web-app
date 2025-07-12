import type { Task } from "$lib/types/task";

export function filterTasksBySearch(tasks: Task[], query: string): Task[] {
  if (!query.trim()) return tasks;

  const lowerQuery = query.toLowerCase();
  return tasks.filter(
    (t) =>
      t.title.toLowerCase().includes(lowerQuery) ||
      (t.description?.toLowerCase().includes(lowerQuery) ?? false)
  );
}