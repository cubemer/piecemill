import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { SpeedBadge } from "@/components/SpeedBadge";
import { cn } from "@/lib/utils";

interface MachineCardProps {
  machineId: number;
  speed: number;
  queueCount: number;
  recommended: boolean;
}

export function MachineCard({
  machineId,
  speed,
  queueCount,
  recommended,
}: MachineCardProps) {
  return (
    <Card
      className={cn(
        "py-3 gap-2",
        recommended && "ring-2 ring-blue-400 bg-blue-950"
      )}
    >
      <CardHeader className="flex-row items-center justify-between px-4 gap-2">
        <CardTitle className="text-lg">M{machineId}</CardTitle>
        <SpeedBadge speed={speed} />
      </CardHeader>
      <CardContent className="px-4">
        <p className="text-sm text-muted-foreground">
          Queue: <span className="text-foreground font-medium">{queueCount}</span>
        </p>
      </CardContent>
    </Card>
  );
}
