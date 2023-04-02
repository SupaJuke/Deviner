import { createContext } from "react";

const TokenContext = createContext({
  token: "",
  setToken: (_: string) => {},
});

export default TokenContext;
