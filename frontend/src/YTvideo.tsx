import { useEffect, useRef } from "react";
import { Media } from "./ctx/MediaCtx";

const YTVideo = ({ media }: { media: Media }) => {
  const video = useRef<HTMLIFrameElement>(null);
  useEffect(() => {
    const checkIframeFocus = () => {
      if (document.activeElement === video.current) {
        video.current?.blur();
        window.focus();
      }
    };
    const intervalId = setInterval(checkIframeFocus, 100);
    return () => clearInterval(intervalId);
  }, []);

  return (
    <div className="yt">
      <div className="yt__fg">
        <iframe
          sandbox="allow-scripts"
          ref={video}
          src={media.Link}
          title="YouTube video player"
          allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
          referrerPolicy="strict-origin-when-cross-origin"
          allowFullScreen
        />
      </div>
    </div>
  );
};
export default YTVideo;
