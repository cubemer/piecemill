import type { Machine, QueueEntry, PBListResponse } from "@/types";

const BASE = import.meta.env.VITE_API_BASE ?? "";

export async function fetchMachines(): Promise<Machine[]> {
  const res = await fetch(
    `${BASE}/api/collections/machines/records?perPage=50`
  );
  if (!res.ok) throw new Error("Failed to fetch machines");
  const data: PBListResponse<Machine> = await res.json();
  return data.items;
}

export async function fetchQueueEntries(): Promise<QueueEntry[]> {
  const res = await fetch(
    `${BASE}/api/collections/queue_entries/records?perPage=500`
  );
  if (!res.ok) throw new Error("Failed to fetch queue entries");
  const data: PBListResponse<QueueEntry> = await res.json();
  return data.items;
}

export async function fetchRecommendation(
  templateId: string,
  quantity: number
): Promise<number[]> {
  const params = new URLSearchParams({
    template_id: templateId,
    quantity: String(quantity),
  });
  const res = await fetch(`${BASE}/recommend?${params}`);
  if (!res.ok) {
    const body = await res.json().catch(() => ({}));
    throw new Error(
      (body as { error?: string }).error ?? `Request failed (${res.status})`
    );
  }
  return res.json();
}
