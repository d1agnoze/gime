import { useEffect, useState } from "react";
import { GetMedia } from "../wailsjs/go/main/App";
import "react-photo-view/dist/react-photo-view.css";
import Image from "./Image";
import ControlBar from "./ControlBar";
import Menu from "./Menu";
import { Media, MediaCtx } from "./ctx/MediaCtx";

function App() {
  const [media, setMedia] = useState<Media>({
    Mime: "",
    Link: "",
  });
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
  }, [media]);

  return (
    <MediaCtx.Provider value={media}>
      <div id="App">
        <ControlBar />
        <div id="portal">
          <div>{type === "image" && <Image media={media}></Image>}</div>
          <Menu />
        </div>
      </div>
    </MediaCtx.Provider>
  );
}

export default App;
