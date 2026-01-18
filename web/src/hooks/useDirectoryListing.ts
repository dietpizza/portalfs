import { useEffect, useState } from "react";
import ky from "ky";
import type { FileMeta } from "../const";

type TListingResponse = Array<FileMeta>;

export function useDirectoryListing(volume: string, path: string = "/") {
  const [dirContent, setDirContent] = useState<TListingResponse | null>(null);
  const [isError, setError] = useState(false);
  const [isLoading, setLoading] = useState(false);

  async function fetchData() {
    setError(false);
    setLoading(true);
    try {
      const response = await ky.get<TListingResponse>(`/listing/${volume}`, {
        searchParams: { path },
      });
      const data = await response.json();

      setDirContent(data);
    } catch {
      setError(true);
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    fetchData();
  }, [path, volume]);

  return { dirContent, isLoading, isError, refetch: fetchData };
}
