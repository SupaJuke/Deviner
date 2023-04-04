import { useContext, useState, useRef, useEffect } from "react";
import TokenContext from "../context";
import { Space, message } from "antd";
import post from "../utils/post";
import Hint from "./Hint";
import Submit from "./Submit";
import Inputs from "./Inputs";

// Constants

const Guess: React.FC = () => {
  const token = useContext(TokenContext).token;
  const [guess, setGuess] = useState("");
  const [width, setWidth] = useState(0);
  const [messageApi, contextHolder] = message.useMessage();
  const ref = useRef<HTMLDivElement>(null);

  // Dynamically resizing submit button
  const handleResize = () => {
    if (ref.current && ref.current.offsetWidth !== width) {
      setWidth(ref.current.offsetWidth);
    }
  };
  const resizeObserver = new ResizeObserver(handleResize);

  useEffect(() => {
    if (ref.current) {
      resizeObserver.observe(ref.current);
    }

    return () => {
      resizeObserver.disconnect();
    };
  }, []);

  // POST calls handlers
  const handleSuccess = (guess: string) => {
    messageApi.open({
      type: "success",
      content: `Guess ${guess} is correct! Now generating new guess...`,
    });
  };

  const handleFailure = (green: string, yellow: string) => {
    messageApi.open({
      type: "error",
      content: <Hint green={green} yellow={yellow} />,
      duration: 10,
    });
  };

  const handleSubmit = async () => {
    const url = "http://localhost:8080/guess";
    const res = await post(
      url,
      { guess: guess },
      { Authentication: "token " + token }
    );
    if (res.success) {
      handleSuccess(guess);
      setGuess("");
    } else {
      if (res.green && res.yellow) {
        handleFailure(res.green, res.yellow);
      }
    }
  };

  // Mouse events handlers
  const handleNumber = (e: React.MouseEvent) => {
    if (guess.length >= 5) {
      return;
    }

    const value = e.currentTarget.getAttribute("value");
    setGuess(guess + value);
  };

  const handleDel = () => {
    if (guess.length == 0) {
      return;
    }

    setGuess(guess.substring(0, guess.length - 1));
  };

  // Styling
  const center = {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
  };

  return (
    <>
      <div
        style={{
          ...center,
          width: "100%",
          height: "100vh",
          background: "#1c2e4a",
          color: "white",
        }}
      >
        <div
          ref={ref}
          style={{
            ...center,
            flexDirection: "column",
            width: "25%",
          }}
        >
          <div
            style={{
              ...center,
              width: "100%",
              minWidth: "250px",
              height: "125px",
              borderStyle: "solid",
              borderRadius: "20px",
              fontSize: "5em",
            }}
          >
            {guess}
          </div>
          <Submit guess={guess} width={width} handleSubmit={handleSubmit} />
          <Inputs handleNumber={handleNumber} handleDel={handleDel} />
        </div>
      </div>
      {contextHolder}
    </>
  );
};

export default Guess;
