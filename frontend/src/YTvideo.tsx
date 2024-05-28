import { Media } from "./ctx/MediaCtx";

const YTVideo = ({ media }: { media: Media }) => {
  console.log(media)
  return (
    <div className="yt">
      <div className="yt__fg">
        <iframe
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
