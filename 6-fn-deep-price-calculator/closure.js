let result = 0;

function calculator() {
  function add(value) {
    result += value;
    return result;
  }

  function subtract(value) {
    result -= value;
    return result;
  }

  function reset() {
    result = 0;
    return result;
  }

  return { add, subtract, reset };
}

const resV2 = calculator();

console.log(resV2.add(10));
console.log(resV2.add(20));
console.log(resV2.add(30));
console.log(resV2.subtract(20));
console.log(resV2.reset());
