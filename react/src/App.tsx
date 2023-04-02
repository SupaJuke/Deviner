import React, { useState } from "react";
import LoginModal from "./components/Login";
import Guess from "./components/Guess";
import TokenContext from "./context";

const App: React.FC<{}> = () => {
  const [token, setToken] = useState("");
  const context = { token, setToken };

  return (
    <TokenContext.Provider value={context}>
      <LoginModal />
      <Guess />
    </TokenContext.Provider>
  );
};

export default App;
