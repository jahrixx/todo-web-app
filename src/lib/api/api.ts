import axios from "axios";

const API = import.meta.env.VITE_API_URL;

export const register = (email: string, password: string) => {
    axios.post(`${API}/register`, { email, password });
}

export const login = async(email: string, password: string) => {
    const response = await axios.post(`${API}/login`, { email, password });
    return response.data;
}

export const getTasks = async(token: string) => {
    try {
        const response = await axios.get(`${API}/tasks`, { headers: { Authorization: `Bearer ${token}` } });
        return Array.isArray(response.data) ? response.data : [];
    } catch (error) {
        console.error(error);
        return [];
    }
}

export const createTask = async(token: string, task: { title: string, description?: string }) => {
    try {
        await axios.post(`${API}/tasks`, task, { headers: { Authorization: `Bearer ${token}` } });
    } catch (error) {
        console.error(error);
    }
}

export const updateTask = async(token: string, task: { id: string, title: string, description?: string, completed?: boolean }) => {
    try {
        await axios.put(`${API}/tasks/${task.id}`, task, { headers: { Authorization: `Bearer ${token}` } });
    } catch (error) {
        console.error(error);
    }
}

export const deleteTask = async(token: string, id: string) => {
    try {
        await axios.delete(`${API}/tasks/${id}`, { headers: { Authorization: `Bearer ${token}` } });
    } catch (error) {
        console.error(error);
    }
}