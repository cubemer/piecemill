import { Badge } from "@/components/ui/badge";

function speedColor(speed: number): string {
  if (speed >= 23) return "bg-green-600 text-white hover:bg-green-600";
  if (speed >= 20) return "bg-yellow-500 text-black hover:bg-yellow-500";
  return "bg-red-600 text-white hover:bg-red-600";
}

export function SpeedBadge({ speed }: { speed: number }) {
  return <Badge className={speedColor(speed)}>{speed}</Badge>;
}
