import {
  DocumentIcon,
  EllipsisVerticalIcon,
} from "@heroicons/react/24/outline";
import { filesize } from "filesize";
import { formatDate } from "../utils";
import type { FileMeta } from "../const";
import { Button } from "./button";
import { Ripple } from "./ripple";

type TFileItemRenderer = {
  file: FileMeta;
};

export function FileItemRenderer({ file }: TFileItemRenderer) {
  const formattedDate = formatDate(file.mtime);
  const formattedSize = filesize(file.size);

  const metaString = [formattedDate, formattedSize].join(" • ");

  return (
    <Ripple>
      <div className="flex items-center w-full py-4 select-none">
        <div className="flex items-center justify-center w-16">
          <DocumentIcon className="h-6 w-6" />
        </div>
        <div className="flex flex-1 min-w-0 flex-col">
          <span className="text-md text-on-surface truncate">
            {file.filename}
          </span>
          <span className="text-sm text-on-surface opacity-70 truncate">
            {metaString}
          </span>
        </div>
        <div className="flex items-center justify-center min-w-16">
          <Button
            corner="full"
            icon={<EllipsisVerticalIcon className="h-6 w-6" />}
          />
        </div>
      </div>
    </Ripple>
  );
}
