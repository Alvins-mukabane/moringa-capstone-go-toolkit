  ███████╗ ██████╗ ██╗     ██╗██████╗ ██╗████████╗██╗   ██╗
  ██╔════╝██╔═══██╗██║     ██║██╔══██╗██║╚══██╔══╝╚██╗ ██╔╝
  ███████╗██║   ██║██║     ██║██║  ██║██║   ██║    ╚████╔╝ 
  ╚════██║██║   ██║██║     ██║██║  ██║██║   ██║     ╚██╔╝  
  ███████║╚██████╔╝███████╗██║██████╔╝██║   ██║      ██║   
  ╚══════╝ ╚═════╝ ╚══════╝╚═╝╚═════╝ ╚═╝   ╚═╝      ╚═╝   

# 🧟‍♂️ Solidity CryptoZombies Clone — Beginner's Toolkit
**Moringa School AI Capstone Project**  
*Prompt-Powered Kickstart: Building a Beginner's Toolkit for Solidity (Web3 Development)*

## 🎯 Project Description
A beginner-friendly toolkit for learning Solidity through Smart Contract development. This project builds a fully functional, multi-contract "Zombie Factory" game inspired by CryptoZombies — featuring DNA generation, feeding cooldowns, leveling up, battling, and ERC721 ownership transfers. It is guided by AI-assisted learning techniques from the Moringa AI curriculum.

Why Solidity for Web3? Solidity is the native language of the Ethereum Virtual Machine (EVM). It enables decentralized, permissionless, and immutable applications (DApps). Bugs in smart contracts can lead to catastrophic financial losses, making strict architecture and testing essential.

## ✨ Key Features
- 🧬 **Cryptographically pseudo-random DNA generation** for unique zombies.
- 🐈 **Cross-contract interactions** (mocking CryptoKitties feeding).
- ⚔️ **Probability-based battle system** tracking wins, losses, and levels.
- ⏳ **Time-based mechanics** (cooldowns between actions using `block.timestamp`).
- 🏠 **ERC721 Token Standard Implementation** for true digital ownership.
- 🧪 **4-test Mocha/Chai suite** covering creation, constraints, leveling, and transfers.
- 📓 **Full AI Prompt Journal** documenting the learning process (see `Toolkit Document.md`).

## 🖥️ System Requirements
- **OS:** macOS / Linux / Windows (WSL recommended)
- **Editor:** VS Code (with Juan Blanco's Solidity extension)
- **Node.js:** v16 or later
- **Git:** For cloning the repository

## ⚙️ Installation & Setup

1. **Install Node.js & npm**
   Download from [nodejs.org](https://nodejs.org/) or use NVM.
2. **Verify installation**
   ```bash
   node --version
   npm --version
   ```
3. **Clone this repository**
   ```bash
   git clone https://github.com/Alvins-mukabane/moringa-capstone-go-toolkit.git
   cd moringa-capstone-go-toolkit
   ```
4. **Install dependencies**
   ```bash
   npm install
   ```
   *This installs Hardhat, Ethers.js, and the necessary testing plugins.*

## 🚀 Basic Usage

**Compile the Smart Contracts**
```bash
npx hardhat compile
```

**Run the test suite**
```bash
npx hardhat test
```

## 💻 Expected Output

```text
Compiled 8 Solidity files successfully (evm target: paris).

  CryptoZombies Full Game
    ✔ Should create a new zombie and emit an event (74ms)
    ✔ Should fail if the same address tries to create a second zombie directly (105ms)
    ✔ Should allow leveling up with fee (90ms)
    ✔ Should allow transfer of ownership via ERC721 (94ms)

  4 passing (3s)
```

## 🧪 Test Results

| Test | What it verifies |
| :--- | :--- |
| `create a new zombie and emit an event` | Correct initial state, array length increase, and `NewZombie` event emission. |
| `fail if the same address tries to create a second zombie` | The `require` constraint enforcing 1 initial free zombie per user works. |
| `allow leveling up with fee` | Payable functions correctly accept Ether (`msg.value`) and modify the zombie's level. |
| `allow transfer of ownership via ERC721` | Standard NFT transfer mechanics (`ownerOf`, `balanceOf`, `transferFrom`) function securely. |

## 📁 Code Structure

```text
capstone/
├── contracts/
│   ├── Ownable.sol          # Access control modifier
│   ├── ERC721.sol           # NFT Standard Interface
│   ├── KittyInterface.sol   # External contract mocking
│   ├── ZombieFactory.sol    # Base zombie generation
│   ├── ZombieFeeding.sol    # Cooldown & feeding logic
│   ├── ZombieHelper.sol     # Name changes, level up fee
│   ├── ZombieAttack.sol     # Battle mechanics & RNG
│   └── ZombieOwnership.sol  # Final contract tying it all together
├── test/
│   └── ZombieFactory.test.js # Hardhat test suite
├── hardhat.config.js        # Hardhat configuration
├── Toolkit Document.md      # AI Prompt Journal & Guide
└── README.md                # This file
```

## 🔑 Key Solidity Concepts

| Concept | What it means | Python/JS equivalent |
| :--- | :--- | :--- |
| `msg.sender` | The address calling the function | Context / User session |
| `require` | Condition checking; reverts changes if false | `if (!cond) throw Error()` |
| `modifier` | Reusable code executed before/after a function | Decorators |
| `mapping` | Key-value store (e.g., `address => uint`) | Hash Map / Dictionary |
| `storage` vs `memory` | Permanent blockchain state vs Temporary execution space | Hard drive vs RAM |
| `payable` | Allows a function to receive Ether | N/A (Web3 specific) |
| `abstract contract` | Base contract missing some implementations | Abstract Class |

## 🔧 Configuration

The project targets the local Hardhat Network by default. To deploy to a testnet (like Sepolia), update `hardhat.config.js`:

```javascript
export default {
  solidity: "0.8.19",
  networks: {
    sepolia: {
      url: "YOUR_ALCHEMY_OR_INFURA_URL",
      accounts: ["YOUR_PRIVATE_KEY"]
    }
  }
};
```
⚠️ **Never commit real private keys to GitHub.** Use `.env` files.

## 🐛 Troubleshooting

| Error | Cause | Fix |
| :--- | :--- | :--- |
| `hardhat is not recognized` | Dependencies not installed | Run `npm install` |
| `Error HH801: Plugin requires dependencies` | Missing peer dependencies | Run `npm install` with `--legacy-peer-deps` or use the exact npm command Hardhat suggests. |
| `Gas limit exceeded` | Infinite loop or massive array operation | Optimize logic; avoid unbounded loops in Solidity. |
| `Reverted with reason string` | A `require` statement failed | Check the exact reason string in the test output. |

## 🧠 Learning Reflections

Coming from traditional languages like Python and JavaScript, these were the biggest mental shifts:

- **Immutability and Cost:** Everything you write to the blockchain costs money (Gas). In Python, you use arrays freely. In Solidity, you learn to pack variables and use mappings to save space.
- **Modifiers are powerful:** Instead of writing the same `if (user == admin)` check everywhere, `onlyOwner` makes the code infinitely cleaner.
- **Inheritance chaining:** Structuring contracts from `Factory` -> `Feeding` -> `Helper` -> `Ownership` keeps files small, readable, and focused on single responsibilities.
- **AI accelerated learning:** Structured prompts for explaining `storage` vs `memory`, generating boilerplate tests, and understanding ERC721 standards took me from zero Solidity to a full multi-contract system rapidly.

## 📚 References
- [Solidity Documentation](https://docs.soliditylang.org/)
- [CryptoZombies](https://cryptozombies.io/)
- [Hardhat Documentation](https://hardhat.org/)
- [OpenZeppelin Contracts](https://openzeppelin.com/contracts/)

## 🤝 Contributing
This is a learning project. If you find an error in the documentation or a bug in the code:
1. Fork the repository
2. Create a branch: `git checkout -b fix/your-fix-name`
3. Commit your changes: `git commit -m 'fix: describe your fix'`
4. Push and open a Pull Request

## 📄 License
MIT License — free to use, modify, and distribute with attribution.

Built with 💻 Solidity + 🛠️ Hardhat + 🤖 AI-assisted learning at Moringa School.
