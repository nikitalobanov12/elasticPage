import { QueryClient } from '@tanstack/react-query';

// Types for our API
export interface User {
  id: string;
  email: string;
  name: string;
  created_at: string;
  updated_at: string;
}

export interface Textbook {
  id: string;
  title: string;
  description: string;
  file_path: string;
  user_id: string;
  user?: User;
  created_at: string;
  updated_at: string;
}

export interface CreateTextbookRequest {
  title: string;
  description?: string;
  user_id: string;
}

export interface UploadTextbookRequest {
  title: string;
  description?: string;
  user_id: string;
  file: File;
}

// API Client class
class ApiClient {
  private baseUrl: string;

  constructor(baseUrl: string = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080') {
    this.baseUrl = baseUrl;
  }

  private async request<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<T> {
    const url = `${this.baseUrl}${endpoint}`;
    
    const config: RequestInit = {
      headers: {
        'Content-Type': 'application/json',
        ...options.headers,
      },
      ...options,
    };

    const response = await fetch(url, config);

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      throw new Error(errorData.error || `HTTP error! status: ${response.status}`);
    }

    return response.json();
  }

  // Textbook API methods
  async getTextbooks(userId?: string): Promise<Textbook[]> {
    const params = userId ? `?user_id=${userId}` : '';
    return this.request<Textbook[]>(`/api/v1/textbooks${params}`);
  }

  async getTextbook(id: string): Promise<Textbook> {
    return this.request<Textbook>(`/api/v1/textbooks/${id}`);
  }

  async createTextbook(data: CreateTextbookRequest): Promise<Textbook> {
    return this.request<Textbook>('/api/v1/textbooks', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async updateTextbook(id: string, data: Partial<CreateTextbookRequest>): Promise<Textbook> {
    return this.request<Textbook>(`/api/v1/textbooks/${id}`, {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  async deleteTextbook(id: string): Promise<{ message: string }> {
    return this.request<{ message: string }>(`/api/v1/textbooks/${id}`, {
      method: 'DELETE',
    });
  }

  async uploadTextbook(data: UploadTextbookRequest): Promise<{ message: string; textbook: Textbook }> {
    const formData = new FormData();
    formData.append('title', data.title);
    formData.append('user_id', data.user_id);
    if (data.description) {
      formData.append('description', data.description);
    }
    formData.append('file', data.file);

    const response = await fetch(`${this.baseUrl}/api/v1/upload`, {
      method: 'POST',
      body: formData,
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      throw new Error(errorData.error || `HTTP error! status: ${response.status}`);
    }

    return response.json();
  }
}

// Export singleton instance
export const apiClient = new ApiClient();

// Query client for React Query
export const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 1000 * 60 * 5, // 5 minutes
      retry: 1,
    },
  },
});