export const UNIT = 80;
export const GAP = 10;

export const multUnit = (m: number) => {
  return m * UNIT + (m - 1) * GAP;
};

export const calcNumButtons = (width: number): number => {
  const numPreCalc = (width / UNIT) | 0;
  if (width - numPreCalc * UNIT - (numPreCalc - 1) * GAP <= 0) {
    return numPreCalc - 1;
  }
  return numPreCalc;
};
