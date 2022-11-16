pragma solidity ^0.8.9;

import "./utils/Ownable.sol";
import "./utils/ReentrancyGuard.sol";

contract EscrowHandler is Ownable, ReentrancyGuard {


    struct Escrow{
        address buyer;
        address seller;
        uint price;
        bool complete;
    }


    mapping(uint => Escrow) public escrows;

    function createNewEscrow(
        address _buyer, 
        address _seller, 
        uint _price, 
        uint escrow_id
        ) public payable isRequiredEth(_price) isAuthorizedAddress(_buyer) nonReentrant 
    {
        Escrow memory newEscrow = Escrow(_buyer, _seller, _price, false); // price = price - commission
        escrows[escrow_id] = newEscrow;
    }

    function releaseEscrow(uint _escrow_id, address _address) public onlyOwner  { // ReentrancyGuard
        require(_address == escrows[_escrow_id].buyer || _address == escrows[_escrow_id].seller, "Can't release funds! Address not authorized");
        escrows[_escrow_id].complete = true;
        transfer(_address, escrows[_escrow_id].price);
    }

    function transfer(address _to, uint _amount) private onlyOwner nonReentrant {
        (bool sent, bytes memory data) = _to.call{value: _amount}("");
    }


    modifier isRequiredEth(uint _amountEth) {
        require(msg.value == _amountEth, "Transaction Failed - Insufficient Ether Amount");
        _;
    }

    modifier isAuthorizedAddress(address _address) {
        require(msg.sender == _address, "Transaction Failed - Address not authorized");
        _;
    }
}