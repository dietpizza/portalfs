import React, { useState } from "react";

import { FileListRenderer, PageScaffold } from "../components";
import { useDirectoryListing } from "../hooks";

export const HomePage: React.FC = () => {
  const [path, setPath] = useState("/");
  const [volume, setVolume] = useState("SweetMix");

  const { dirContent, isLoading, isError } = useDirectoryListing(volume, path);

  return (
    <PageScaffold>
      {dirContent?.length && <FileListRenderer list={dirContent} />}
    </PageScaffold>
  );
};
