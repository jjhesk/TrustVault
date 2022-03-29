// SPDX-License-Identifier: agpl-3.0
pragma solidity ^0.6.12;

abstract contract IHashOracle {

    function y() external virtual view returns (bytes32);

    function w() external virtual view returns (bytes32);

    function x() external virtual view returns (bytes32);

    function natural(uint256 roundId, uint256 hash) external virtual returns (bytes32);

}
