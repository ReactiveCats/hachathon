import ms from 'ms';

export const convertTime = (value) => {
  const seconds = parseInt(value, 10)

  if (seconds.toString() === value) {
    return seconds;
  }

  if (typeof value === 'string') {
    return ms(value) / 1000;
  }

  return ms(value * 1000);
};