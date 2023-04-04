interface Response {
  success: boolean;
  message: string;
  token?: string;
  green?: string;
  yellow?: string;
}

export interface CredentialInput {
  username: string;
  password: string;
}

export interface GuessInput {
  guess: string;
}

const post = async (
  url: string,
  input: CredentialInput | GuessInput,
  headers?: HeadersInit | undefined
): Promise<Response> => {
  const res = await fetch(url, {
    method: "POST",
    body: JSON.stringify(input),
    headers: {
      ...headers,
      "Content-Type": "application/json",
    },
  });
  return res.json();
};

export default post;
