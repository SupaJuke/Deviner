const Hint: React.FC<{ green: string; yellow: string }> = (props) => {
  return (
    <>
      <h3 style={{ fontSize: "1.5em", margin: 0 }}>Guess incorrect</h3>
      <p style={{ fontSize: "1.5em", margin: "5px" }}>
        <span style={{ background: "lightgreen" }}>
          Green Digit:
          <span style={{ fontWeight: "bold" }}> {props.green} </span>
        </span>

        <br />

        <span style={{ background: "gold" }}>
          Yellow Digit:
          <span style={{ fontWeight: "bold" }}> {props.yellow} </span>
        </span>
      </p>
    </>
  );
};

export default Hint;
