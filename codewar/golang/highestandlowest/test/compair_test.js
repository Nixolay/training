const compair = require("./compair");

const mocha = require("mocha");
const chai = require("chai");

chai.should();

describe("HighAndLow", function() {
  it(
    "When function get string with numbers, " +
      "the function return string with high and low numbers",
    function() {
      it("should return -1 when the value is not present", function() {
        chai.assert.equal(
          compair.HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"),
          "42 -9"
        );
      });
    }
  );
});
