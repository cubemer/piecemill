import { useAppStore } from "@/store";
import { MachineCard } from "@/components/MachineCard";
import { queueCountByMachine } from "@/lib/queueUtils";

const TOP_ROW = [11, 9, 7, 5, 3, 1];
const BOTTOM_ROW = [12, 10, 8, 6, 4, 2];
const GRID_ORDER = [...TOP_ROW, ...BOTTOM_ROW];

export function MachineGrid() {
  const machines = useAppStore((s) => s.machines);
  const queueEntries = useAppStore((s) => s.queueEntries);
  const recommendedIds = useAppStore((s) => s.recommendedIds);

  const counts = queueCountByMachine(queueEntries);
  const machineMap = new Map(machines.map((m) => [m.machine_id, m]));

  return (
    <div className="grid grid-cols-6 gap-3">
      {GRID_ORDER.map((mid) => {
        const machine = machineMap.get(mid);
        return (
          <MachineCard
            key={mid}
            machineId={mid}
            speed={machine?.speed ?? 25}
            queueCount={counts[mid] ?? 0}
            recommended={recommendedIds.includes(mid)}
          />
        );
      })}
    </div>
  );
}
