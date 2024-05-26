import { useContext, useEffect, useState } from "react";
import { Quit } from "../wailsjs/go/main/App";
import { MediaCtx } from "./ctx/MediaCtx";

const Menu = () => {
  const [open, setOpen] = useState(false);
  const media = useContext(MediaCtx);

  useEffect(() => {
    document.addEventListener("keyup", handleKeys);
    return () => {
      document.removeEventListener("keyup", handleKeys);
    };
  }, []);

  const keyBinds: keyBind[] = [
    {
      key: ["Escape", "q"],
      callback: () => {
        Quit();
      },
      desc: "Quit GIME",
    },
    {
      key: ["?"],
      callback: () => {
        console.log(open);
        setOpen((p) => !p);
      },
      desc: "Open help menu",
    },
    {
      key: ["y"],
      callback: ( ) => {
        if (media?.Link == null) return;
        navigator.clipboard.writeText(media.Link);
      },
      desc: "Copy media link",
    }
  ];

  const handleKeys = (e: KeyboardEvent) => {
    keyBinds.forEach((keyBind) => {
      if (keyBind.key.includes(e.key)) {
        keyBind.callback();
      }
    });
  };

  return (
    <div className={`modal ${open ? "modal__active" : ""}`}>
      <h1>Help</h1>
      <ul>
        {keyBinds.map((keyBind, index) => (
          <li key={index}>
            {keyBind.key.map((key, idx) => (
              <kbd key={idx}>{key}</kbd>
            ))}
            - <p>{keyBind.desc}</p>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Menu;

interface keyBind {
  key: string[];
  callback: () => void;
  desc: string;
}
