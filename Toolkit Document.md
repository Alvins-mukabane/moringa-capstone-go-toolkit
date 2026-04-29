# Getting Started with Solidity & Web3 – A Beginner’s Guide

## 1. Title & Objective
**Title:** "Prompt-Powered Kickstart: Building a Beginner’s Toolkit for Solidity"
**Technology Chosen:** Solidity (Ethereum Smart Contracts)
**Why Chosen:** Solidity is the primary language for building Decentralized Applications (DApps) on Ethereum and other EVM-compatible blockchains. It's the foundation of the Web3 revolution, enabling smart contracts that handle billions of dollars in value autonomously.
**End Goal:** Create a "Zombie Factory" smart contract that generates unique zombies with random DNA, based on the famous CryptoZombies curriculum.

## 2. Quick Summary of the Technology
Solidity is an object-oriented, high-level language for implementing smart contracts. It is statically typed, supports inheritance, libraries, and complex user-defined types.
**Where is it used?** Primarily on the Ethereum Virtual Machine (EVM). It's used for DeFi protocols (Uniswap, Aave), NFT marketplaces (OpenSea), and Decentralized Autonomous Organizations (DAOs).
**Real-world example:** Creating an NFT collection where each token has unique properties stored on-chain.

## 3. System Requirements
- **OS:** Linux, Mac, or Windows
- **Tools/Editors required:** VS Code (with the "Solidity" extension by Juan Blanco)
- **Packages:** Node.js, NPM, and Hardhat (Development Environment)

## 4. Installation & Setup Instructions

### Step 1: Install Node.js
Ensure you have Node.js installed. You can check with `node -v`.

### Step 2: Initialize Hardhat
```bash
mkdir capstone
cd capstone
npm init -y
npm install --save-dev hardhat @nomicfoundation/hardhat-toolbox
npx hardhat init
# Choose "Create an empty hardhat.config.js"
```

### Step 3: Create the Contract
Create a folder named `contracts` and add `ZombieFactory.sol`.

## 5. Minimal Working Example
This example implements the "Zombie Factory" logic. It allows users to create one unique zombie linked to their Ethereum address.

**Code (`contracts/ZombieFactory.sol`):**
```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract ZombieFactory {

    // Event to notify the frontend when a new zombie is created
    event NewZombie(uint zombieId, string name, uint dna);

    uint dnaDigits = 16;
    uint dnaModulus = 10 ** dnaDigits;

    struct Zombie {
        string name;
        uint dna;
    }

    // Array of all zombies
    Zombie[] public zombies;

    // Mappings to track ownership
    mapping (uint => address) public zombieToOwner;
    mapping (address => uint) ownerZombieCount;

    // Internal function to create a zombie and store it
    function _createZombie(string memory _name, uint _dna) internal {
        zombies.push(Zombie(_name, _dna));
        uint id = zombies.length - 1;
        zombieToOwner[id] = msg.sender;
        ownerZombieCount[msg.sender]++;
        emit NewZombie(id, _name, _dna);
    }

    // Generate a pseudo-random DNA from a string
    function _generateRandomDna(string memory _str) private view returns (uint) {
        uint rand = uint(keccak256(abi.encodePacked(_str)));
        return rand % dnaModulus;
    }

    // Public function for users to get their first zombie
    function createRandomZombie(string memory _name) public {
        require(ownerZombieCount[msg.sender] == 0, "Only one zombie per address!");
        uint randDna = _generateRandomDna(_name);
        _createZombie(_name, randDna);
    }
}
```

## 6. AI Prompt Journal

- **Prompt used:** "Explain the difference between `memory`, `storage`, and `calldata` in Solidity for a Java developer."
- **AI’s response summary:** The AI explained that `storage` is like a computer's hard drive (persistent on the blockchain), while `memory` is like RAM (temporary during execution).
- **Evaluation:** Crucial for understanding gas optimization and how data is passed between functions.

- **Prompt used:** "Give me a Hardhat test script using Ethers.js to verify that a contract emits an event when a function is called."
- **AI’s response summary:** The AI provided a snippet using the `expect(tx).to.emit(...)` syntax from the Hardhat Toolbox.
- **Evaluation:** Very helpful. Testing events is the best way to verify state changes in smart contracts.

## 7. Common Issues & Fixes
- **Error:** `ParserError: Expected ';' but got 'uint'`
  - **Fix:** Solidity is very strict about semicolons. Check every line!
- **Error:** `Gas limit exceeded`
  - **Fix:** Your function might be doing too much work or have an infinite loop. In our case, ensure your DNA generation isn't overly complex.
- **Issue:** The DNA is always the same for the same name.
  - **Fix:** `keccak256` is deterministic. For true randomness, you would need an Oracle like Chainlink VRF, but for this "Hello World" toolkit, deterministic DNA is fine for demonstration.

## 8. References
- [Solidity Documentation](https://docs.soliditylang.org/)
- [CryptoZombies Interactive Course](https://cryptozombies.io/)
- [Hardhat Documentation](https://hardhat.org/docs)
- [OpenZeppelin (Security Best Practices)](https://openzeppelin.com/contracts/)
