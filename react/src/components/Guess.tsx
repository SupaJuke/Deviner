import { useContext } from "react";
import TokenContext from "../context";

const Guess: React.FC = () => {
  const token = useContext(TokenContext).token;
  return (
    <div
      style={{
        display: "flex",
        width: "100%",
        height: "100vh",
        background: "#1c2e4a",
        color: "white",
        fontSize: "3em",
        alignItems: "center",
        justifyContent: "center",
      }}
    >
      <p style={{ width: "50%", wordBreak: "break-all", textAlign: "center" }}>
        {token}
      </p>
    </div>
  );
};

export default Guess;
