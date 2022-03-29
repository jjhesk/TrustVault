pragma solidity ^0.6.0;

abstract contract IMark6 {

    function drawWin() external virtual returns (bytes32);


    function numbersDrawn(
        uint256 _lotteryId,
        bytes32 _requestId,
        uint256 _randomNumber
    ) external virtual;


}
