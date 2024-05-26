import { createContext } from "react";

export interface Media {
  Mime: string;
  Link: string;
}

const MediaCtx = createContext<Media | null>(null);

export { MediaCtx };
