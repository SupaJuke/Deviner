import { useContext, useState } from "react";
import TokenContext from "../context";
import { Button, Card, Space } from "antd";
import post, { GuessInput } from "../utils/post";

const Guess: React.FC = () => {
  const [guess, setGuess] = useState("");
  const token = useContext(TokenContext).token;

  const handleSubmit = async () => {
    const url = "http://localhost:8080/guess";
    const res = await post(
      url,
      { guess: guess },
      { Authentication: "token " + token }
    );
    console.log(res);
    // TODO: implement result checking
  };

  const handleNumber = (e: React.MouseEvent) => {
    if (guess.length >= 5) {
      return;
    }

    const value = e.currentTarget.getAttribute("value");
    setGuess(guess + value);
    console.log(guess);
  };

  const handleDel = () => {
    if (guess.length == 0) {
      return;
    }

    setGuess(guess.substring(0, guess.length - 1));
  };

  const nums = [...Array(10).keys()];
  nums.shift();
  nums.push(0);

  const center = {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
  };

  return (
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

        <br />

        <Button
          ghost={true}
          style={{
            ...center,
            width: "100%",
            minWidth: "250px",
            height: "80px",
            fontSize: "3em",
          }}
          disabled={guess.length !== 5}
          onClick={handleSubmit}
        >
          Submit
        </Button>

        <br />

        <Space
          size={[10, 10]}
          wrap={true}
          style={{
            width: "100%",
            minWidth: "250px",
            maxWidth: "550px",
            justifyContent: "center",
          }}
        >
          {nums.map((num) => {
            return (
              <Button
                value={num}
                key={num}
                onClick={handleNumber}
                ghost={true}
                style={{
                  width: "80px",
                  height: "80px",
                  fontSize: "3em",
                }}
              >
                {num}
              </Button>
            );
          })}
          <Button
            ghost={true}
            style={{
              width: "170px",
              height: "80px",
              fontSize: "3em",
            }}
            onClick={handleDel}
          >
            Delete
          </Button>
        </Space>
      </div>
    </div>
  );
};

export default Guess;
