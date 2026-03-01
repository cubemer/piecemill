import { create } from "zustand";
import type { Machine, QueueEntry } from "@/types";
import { fetchMachines, fetchQueueEntries, fetchRecommendation } from "@/lib/api";

interface AppState {
  machines: Machine[];
  queueEntries: QueueEntry[];
  recommendedIds: number[];
  templateId: string;
  quantity: number;
  loading: boolean;
  error: string | null;

  setTemplateId: (id: string) => void;
  setQuantity: (qty: number) => void;
  loadData: () => Promise<void>;
  recommend: () => Promise<void>;
}

export const useAppStore = create<AppState>((set, get) => ({
  machines: [],
  queueEntries: [],
  recommendedIds: [],
  templateId: "",
  quantity: 1,
  loading: false,
  error: null,

  setTemplateId: (id) => set({ templateId: id, recommendedIds: [], error: null }),
  setQuantity: (qty) => set({ quantity: qty, recommendedIds: [], error: null }),

  loadData: async () => {
    try {
      const [machines, queueEntries] = await Promise.all([
        fetchMachines(),
        fetchQueueEntries(),
      ]);
      set({ machines, queueEntries });
    } catch (err) {
      set({ error: err instanceof Error ? err.message : "Failed to load data" });
    }
  },

  recommend: async () => {
    const { templateId, quantity } = get();
    if (!templateId || quantity < 1) return;

    set({ loading: true, error: null });
    try {
      const ids = await fetchRecommendation(templateId, quantity);
      set({ recommendedIds: ids, loading: false });
    } catch (err) {
      set({
        error: err instanceof Error ? err.message : "Recommendation failed",
        loading: false,
        recommendedIds: [],
      });
    }
  },
}));
