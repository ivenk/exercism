export const gigasecond = (date) => {
    return new Date(date.getTime() + GIGASECOND*1000);
};

const GIGASECOND = 1000000000
