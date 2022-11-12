pragma solidity ^0.8.16;

contract Storage {

    uint256 x;

    function setX ( uint256 _newVal ) public {
        x = _newVal;
    }

    function getX () public view returns ( uint256 ) {
        return x;
    }
}