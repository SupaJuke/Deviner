import { useState, useContext, useEffect } from "react";
import { Button, Modal, Form, Input } from "antd";
import TokenContext from "../context";

interface Credential {
  username: string;
  password: string;
}

interface Response {
  success: boolean;
  message: string;
  token?: string;
}

const LoginModal: React.FC = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const setToken = useContext(TokenContext).setToken;

  const showModal = (bool: boolean) => {
    setIsModalOpen(bool);
  };

  useEffect(() => {
    showModal(true);
  }, []);

  const post = async (url: string, cred: Credential): Promise<Response> => {
    const res = await fetch(url, {
      method: "POST",
      body: JSON.stringify(cred),
      headers: {
        "Content-Type": "application/json",
      },
    });
    return res.json();
  };

  const onFinish = async (cred: Credential) => {
    const url = "http://localhost:8080/login";
    const res = await post(url, cred);
    console.log(res);
    if (res.token) {
      setToken(res.token);
      showModal(false);
    }
  };

  const loginForm = () => (
    <Form
      name="basic"
      labelCol={{ span: 5 }}
      initialValues={{ remember: true }}
      onFinish={onFinish}
      autoComplete="off"
      labelAlign="left"
    >
      <Form.Item
        label="Username"
        name="username"
        rules={[{ required: true, message: "Please input your username!" }]}
      >
        <Input />
      </Form.Item>

      <Form.Item
        label="Password"
        name="password"
        rules={[{ required: true, message: "Please input your password!" }]}
      >
        <Input.Password />
      </Form.Item>

      <Form.Item style={{ width: "100%" }}>
        <Button type="primary" htmlType="submit" style={{ width: "100%" }}>
          Submit
        </Button>
      </Form.Item>
    </Form>
  );

  return (
    <Modal
      maskClosable={false}
      open={isModalOpen}
      closable={false}
      centered={true}
      destroyOnClose={true}
      bodyStyle={{ width: "100%" }}
      footer={null}
    >
      <center>
        <h1 style={{ marginTop: "0em" }}>Login</h1>
      </center>
      {loginForm()}
    </Modal>
  );
};

export default LoginModal;
