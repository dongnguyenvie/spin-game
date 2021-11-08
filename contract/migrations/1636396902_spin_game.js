const spinProccessor = artifacts.require("SpinProccessor");

module.exports = function (_deployer, network, accounts) {
  console.log("accounts", accounts);
  const owner = accounts[0];
  const numberConbfirmationRequired = 2;
  const usdtAddr = "0x3b00ef435fa4fcff5c209a37d1f3dcff37c705ad";
  _deployer.deploy(spinProccessor, usdtAddr);
};
