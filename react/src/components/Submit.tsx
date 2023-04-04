import { Button } from "antd";
import { UNIT, multUnit, calcNumButtons } from "../utils/size";

type Props = {
  guess: string;
  width: number;
  handleSubmit: () => void;
};

const Submit: React.FC<Props> = (props) => {
  const center = {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
  };

  return (
    <div style={{ margin: 10 }}>
      <Button
        ghost={true}
        disabled={props.guess.length !== 5}
        onClick={props.handleSubmit}
        style={{
          ...center,
          width: `${multUnit(calcNumButtons(props.width))}px`,
          minWidth: `${multUnit(2)}px`,
          maxWidth: `${multUnit(6)}px`,
          height: `${UNIT}px`,
          fontSize: "3em",
        }}
      >
        Submit
      </Button>
    </div>
  );
};

export default Submit;
