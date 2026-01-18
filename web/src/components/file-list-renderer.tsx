import { VList } from "virtua";
import { FileItemRenderer } from "./file-item-renderer";
import type { FileMeta } from "../const";

type TFileListRenderer = {
  list: Array<FileMeta>;
};

export function FileListRenderer({ list }: TFileListRenderer) {
  return (
    <div className="flex flex-col h-full bg-surface-container">
      <VList className="h-full">
        {list.map((e) => (
          <FileItemRenderer key={e.filename} file={e} />
        ))}
      </VList>
    </div>
  );
}
