export const value = (colorsIn) => {
    return parseInt("" + COLORS.indexOf(colorsIn[0]) + COLORS.indexOf(colorsIn[1]));
};

const COLORS =["black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"];
