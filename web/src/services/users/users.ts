import apiClient from '../api/client';

interface User {
  id: string;
  name: string;
  email: string;
}

export const userService = {
  async getUsers(page: number = 1): Promise<User[]> {
    const response = await apiClient.get(`/users?page=${page}`);
    return response.data;
  },

  async createUser(userData: Omit<User, 'id'>): Promise<User> {
    const response = await apiClient.post('/users', userData);
    return response.data;
  }
};
