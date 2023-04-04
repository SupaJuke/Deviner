import { useState, useContext, useEffect } from "react";
import { Button, Modal, Form, Input, message } from "antd";
import TokenContext from "../context";
import post, { CredentialInput } from "../utils/post";

const LoginModal: React.FC = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const setToken = useContext(TokenContext).setToken;
  const [messageApi, contextHolder] = message.useMessage();

  const showModal = (bool: boolean) => {
    setIsModalOpen(bool);
  };

  useEffect(() => {
    showModal(true);
  }, []);

  const handleSuccess = (user: string) => {
    messageApi.open({
      type: "success",
      content: `Logged in successfully. Welcome ${user}!`,
    });
  };

  const handleFailure = () => {
    messageApi.open({
      type: "error",
      content: "Username or password incorrect",
    });
  };

  const onFinish = async (cred: CredentialInput) => {
    const url = "http://localhost:8080/login";
    const res = await post(url, cred);
    if (res.token) {
      setToken(res.token);
      handleSuccess(cred.username);
      setTimeout(() => showModal(false), 1000);
    } else {
      handleFailure();
    }
    // TODO: alert or smth once login fails
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
    <>
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
      {contextHolder}
    </>
  );
};

export default LoginModal;
