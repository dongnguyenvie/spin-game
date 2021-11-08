// SPDX-License-Identifier: MIT
pragma solidity ^0.8.1;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract SpinProccessor {
    address private _owner;
    ERC20 private _usdt;
    mapping(address => uint256) _deposits;

    event Deposit(address sender, uint256 amount);

    constructor(address usdt) {
        _owner = msg.sender;
        _usdt = ERC20(usdt);
    }

    modifier onlyOwner() {
        require(msg.sender == _owner, "only owner!!");
        _;
    }

    function totalDepositOf(address from) public view returns (uint256) {
        return _deposits[from];
    }

    function deposit() public payable {
        require(msg.value > 0, "The minimum amount is invalid");
        _deposits[msg.sender] += msg.value;
        emit Deposit(msg.sender, msg.value);
    }

    function depositUsd(uint256 amount) public {
        require(amount > 0, "The minimum amount is invalid");
        require(
            _usdt.allowance(msg.sender, address(this)) >= amount,
            "Please allow your money first"
        );

        emit Deposit(msg.sender, amount);
    }

    function withdraw() public onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "The minimum amount is invalid");
        payable(_owner).transfer(balance);
    }

    function withdrawUSdt() public onlyOwner {
        uint256 balance = _usdt.balanceOf(address(this));
        require(balance >= 0, "The minimum amount is invalid");
        _usdt.transfer(_owner, balance);
    }
}
