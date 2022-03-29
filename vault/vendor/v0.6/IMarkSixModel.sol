// SPDX-License-Identifier: agpl-3.0
pragma solidity >=0.5.15;

interface IMarkSixModel {
    function priceLevel(uint8 found, bool hitExtra) external pure returns (uint8);

    function checkRepeat(uint8[6] memory sample) external pure returns (bool);

    function usdFullInteger(uint256 amount) external pure returns (uint256);

    function calcPrizeInteger(uint8 division) external pure returns (uint);
    //3.6(a),(b)
    function miniDistribution(uint256 af)
    external pure returns (uint256 ff1, uint256 ff2, uint256 ff3);
    //3.5(b)(i),(ii),(iii)
    function fundbearing(uint256 af, uint level1c, uint level2c, uint level3c)
    external pure returns
    (uint256 f1, uint256 f2, uint256 f3, uint256 n1, uint256 n2, uint256 n3);

    //claim fund checker
    function claimFundChecker(uint8 division, uint256 snowball, uint level1c, uint level2c, uint level3c, uint256 af)
    external pure returns
    (uint256 claim_amount);
    // current pool, b
    // fund reduction incurred
    // assumptions: 1. there is at least a top prize discovered. 2. the reduction of regular prizes are final

    function calcAfSb(bool haveGrandPrize, uint256 income, uint256 reduce, uint256 pool_size_actual, uint256 previous_snowball) external pure returns (uint256 _pf, uint256 _sb);
}
