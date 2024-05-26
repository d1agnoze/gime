import { useEffect, useState } from "react";
import { GetMedia } from "../wailsjs/go/main/App";
import "react-photo-view/dist/react-photo-view.css";
import Image from "./Image";
import ControlBar from "./ControlBar";
import Menu from "./Menu";
import { Media, MediaCtx } from "./ctx/MediaCtx";

function App() {
  const [media, setMedia] = useState<Media>({ Mime: "", Link: "" });
  const [type, setType] = useState<"image" | "video" | null>(null);

  useEffect(() => {
    GetMedia().then((result) => {
      setMedia(result);
    });
  }, []);

  useEffect(() => {
    if (media.Mime.startsWith("image")) {
      setType("image");
    }
    if (media.Mime.startsWith("video") && media.Mime.includes("youtube")) {
      setType("video");
    }
  }, [media]);

  return (
    <MediaCtx.Provider value={media}>
      <div id="App">
        <ControlBar />
        <div id="portal">
          <div>
            {type === "image" && <Image media={media}></Image>}
            {type === "video" && (
              <iframe
                width="560"
                height="315"
                src="https://www.youtube.com/embed/XDrjnBUByW0?si=OQXUpJvySjYhf2a_"
                title="YouTube video player"
                allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
                referrerPolicy="strict-origin-when-cross-origin"
                allowFullScreen
              ></iframe>
            )}
          </div>
          <Menu />
        </div>
      </div>
    </MediaCtx.Provider>
  );
}

export default App;
