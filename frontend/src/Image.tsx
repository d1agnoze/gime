import { useEffect, useRef } from "react";
import { PhotoProvider, PhotoView } from "react-photo-view";
import "react-photo-view/dist/react-photo-view.css";

const Image = ({ media }: { media: Image_Props }) => {
  const trig = useRef<HTMLButtonElement>(null);

  useEffect(() => {
    setTimeout(() => trig.current?.click(), 10);
  });

  return (
    <div>
      <PhotoProvider
        bannerVisible={false}
        photoClosable={false}
        maskClosable={false}
        overlayRender={(_) => OverLay(media.Link)}
        portalContainer={document.getElementById("portal")!}
        brokenElement={<div>GIME could not load the image</div>}
      >
        <div className="container">
          <PhotoView src={media.Link}>
            <button ref={trig}>Reopen photo</button>
          </PhotoView>
        </div>
      </PhotoProvider>
    </div>
  );
};

const OverLay = (url: string) => {
  return (
    <div className="rpv_overlay">
      <div className="rpv_overlay_inner">
        <a href={url} onClick={(e) => onLinkClick(e, url)}>
          {url}
        </a>
      </div>
      <div className="rpv_overlay_copy">
        <button onClick={() => copy(url)}>
          <i className="fa-regular fa-clone"></i>
        </button>
      </div>
    </div>
  );
};

const onLinkClick = (e: any, link: string | undefined | null) => {
  e.preventDefault();
  if (link == null) return;
  //@ts-ignore: Wails runtime action
  window.runtime.BrowserOpenURL(link);
};

const copy = (content: string) => {
  navigator.clipboard.writeText(content);
};

interface Image_Props {
  Mime: string;
  Link: string;
}

export default Image;
