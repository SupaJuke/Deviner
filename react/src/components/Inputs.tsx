import { Space, Button } from "antd";
import { UNIT, multUnit } from "../utils/size";

type Handlers = {
  handleNumber: (e: React.MouseEvent) => void;
  handleDel: () => void;
};

const Inputs: React.FC<Handlers> = ({ handleNumber, handleDel }) => {
  const nums = [...Array(10).keys()];
  nums.shift();
  nums.push(0);

  return (
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
            ghost={true}
            onClick={(e) => handleNumber(e)}
            style={{
              width: `${UNIT}px`,
              height: `${UNIT}px`,
              fontSize: "3em",
            }}
          >
            {num}
          </Button>
        );
      })}
      <Button
        ghost={true}
        onClick={() => handleDel()}
        style={{
          width: `${multUnit(2)}px`,
          height: `${UNIT}px`,
          fontSize: "3em",
        }}
      >
        Delete
      </Button>
    </Space>
  );
};

export default Inputs;
