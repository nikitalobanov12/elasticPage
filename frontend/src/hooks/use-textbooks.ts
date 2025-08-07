import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { apiClient, type Textbook, type CreateTextbookRequest, type UploadTextbookRequest } from '~/lib/api-client';

// Query keys
export const textbookKeys = {
  all: ['textbooks'] as const,
  lists: () => [...textbookKeys.all, 'list'] as const,
  list: (userId?: string) => [...textbookKeys.lists(), { userId }] as const,
  details: () => [...textbookKeys.all, 'detail'] as const,
  detail: (id: string) => [...textbookKeys.details(), id] as const,
};

// Hooks
export function useTextbooks(userId?: string) {
  return useQuery({
    queryKey: textbookKeys.list(userId),
    queryFn: () => apiClient.getTextbooks(userId),
  });
}

export function useTextbook(id: string) {
  return useQuery({
    queryKey: textbookKeys.detail(id),
    queryFn: () => apiClient.getTextbook(id),
    enabled: !!id,
  });
}

export function useCreateTextbook() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: CreateTextbookRequest) => apiClient.createTextbook(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: textbookKeys.lists() });
    },
  });
}

export function useUpdateTextbook() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: Partial<CreateTextbookRequest> }) =>
      apiClient.updateTextbook(id, data),
    onSuccess: (_, { id }) => {
      queryClient.invalidateQueries({ queryKey: textbookKeys.detail(id) });
      queryClient.invalidateQueries({ queryKey: textbookKeys.lists() });
    },
  });
}

export function useDeleteTextbook() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => apiClient.deleteTextbook(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: textbookKeys.lists() });
    },
  });
}

export function useUploadTextbook() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: UploadTextbookRequest) => apiClient.uploadTextbook(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: textbookKeys.lists() });
    },
  });
}