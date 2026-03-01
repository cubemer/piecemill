import { useEffect } from "react";
import { useAppStore } from "@/store";
import { ControlBar } from "@/components/ControlBar";
import { MachineGrid } from "@/components/MachineGrid";

function App() {
  const loadData = useAppStore((s) => s.loadData);

  useEffect(() => {
    loadData();
  }, [loadData]);

  return (
    <div className="dark min-h-screen bg-background text-foreground p-6 flex flex-col gap-6">
      <ControlBar />
      <MachineGrid />
    </div>
  );
}

export default App;
