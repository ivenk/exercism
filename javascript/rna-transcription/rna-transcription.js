export const toRna = (dna) => {
    var result = "";
    for(var d of dna) {
	result += trans.get(d);
    }
    return result;
};

const trans = new Map([
    ['G','C'],
    ['C','G'],
    ['T','A'],
    ['A','U'],
]);
