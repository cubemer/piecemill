import type { QueueEntry } from "@/types";

export function queueCountByMachine(
  entries: QueueEntry[]
): Record<number, number> {
  const counts: Record<number, number> = {};
  for (const e of entries) {
    counts[e.machine_id] = (counts[e.machine_id] ?? 0) + e.quantity;
  }
  return counts;
}
