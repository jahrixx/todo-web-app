import { writable } from "svelte/store";
import { browser } from "$app/environment";

const storedToken = browser ? localStorage.getItem("token") : null;
const API_URL = import.meta.env.VITE_API_URL;

export const token = writable(storedToken || null);
export const userId = writable<string | null>(null);

token.subscribe((value) => {
  if (browser) {
    if (value) {
      localStorage.setItem("token", value);
    } else {
      localStorage.removeItem("token");
    }
  }
})

export async function registerUser(email: string, password: string) {
  const res = await fetch(`${API_URL}/auth/register`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password }),
  });

  if (!res.ok) throw new Error("Registration failed");

  console.log(`${API_URL}/auth/register`);  
  return await res.json();
}

export async function loginUser(email: string, password: string) {
  const res = await fetch(`${API_URL}/auth/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password }),
  });

  if (!res.ok) throw new Error("Invalid credentials");

  console.log(`${API_URL}/auth/login`);  
  return await res.json(); // { token, user_id }
}
