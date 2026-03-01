export interface Machine {
  id: string;
  machine_id: number;
  speed: number;
}

export interface QueueEntry {
  id: string;
  template_id: string;
  machine_id: number;
  quantity: number;
}

export interface PBListResponse<T> {
  page: number;
  perPage: number;
  totalPages: number;
  totalItems: number;
  items: T[];
}
