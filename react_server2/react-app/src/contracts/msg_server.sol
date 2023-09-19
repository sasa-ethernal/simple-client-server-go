// SPDX-License-Identifier: GPL-3.0

pragma solidity >0.8.15;

contract Msg_server {

    event Request(address recipient, address sender, bytes content);
    event Deliver(address recipient, bytes content);

    // mapping(uint256 => address) public transaction_origins;
    mapping(address => string) public rsa_keys;

    function RequestPolicy(address recipient, bytes memory content) public {
        emit Request(recipient, msg.sender, content);
    }

    function DeliverPolicy(address recipient, bytes memory content) public {
        emit Deliver(recipient, content);
    }
    
    function UpdateRsaKey(string memory public_key) public {
        rsa_keys[msg.sender] = public_key;
    }
}