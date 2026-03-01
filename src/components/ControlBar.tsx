import { useAppStore } from "@/store";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

export function ControlBar() {
  const templateId = useAppStore((s) => s.templateId);
  const quantity = useAppStore((s) => s.quantity);
  const loading = useAppStore((s) => s.loading);
  const error = useAppStore((s) => s.error);
  const setTemplateId = useAppStore((s) => s.setTemplateId);
  const setQuantity = useAppStore((s) => s.setQuantity);
  const recommend = useAppStore((s) => s.recommend);

  const disabled = loading || !templateId || quantity < 1;

  return (
    <div className="flex items-center gap-3">
      <h1 className="text-xl font-bold whitespace-nowrap mr-4">
        Piecemill Dispatch
      </h1>
      <Input
        type="text"
        placeholder="Template ID"
        value={templateId}
        onChange={(e) => setTemplateId(e.target.value)}
        className="w-44"
      />
      <Input
        type="number"
        min={1}
        placeholder="Qty"
        value={quantity}
        onChange={(e) => setQuantity(Number(e.target.value))}
        className="w-24"
      />
      <Button disabled={disabled} onClick={recommend}>
        {loading ? "..." : "Recommend"}
      </Button>
      {error && (
        <span className="text-red-400 text-sm ml-2">{error}</span>
      )}
    </div>
  );
}
