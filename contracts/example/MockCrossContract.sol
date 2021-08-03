// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "../core/IContractModule.sol";

contract MockCrossContract is IContractModule {
    function onContractCall(CommitMode commitMode, bytes calldata callInfo) external override returns (bytes memory) {
        return new bytes(1);
    }
}
