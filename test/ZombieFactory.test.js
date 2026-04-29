import { expect } from "chai";
import pkg from "hardhat";
const { ethers } = pkg;

describe("CryptoZombies Full Game", function () {
  let zombieOwnership;
  let owner, addr1, addr2;

  beforeEach(async function () {
    [owner, addr1, addr2] = await ethers.getSigners();
    const ZombieOwnership = await ethers.getContractFactory("ZombieOwnership");
    zombieOwnership = await ZombieOwnership.deploy();
    await zombieOwnership.waitForDeployment();
  });

  it("Should create a new zombie and emit an event", async function () {
    const zombieName = "Survivor";
    await expect(zombieOwnership.createRandomZombie(zombieName))
      .to.emit(zombieOwnership, "NewZombie");

    const zombie = await zombieOwnership.zombies(0);
    expect(zombie.name).to.equal(zombieName);
    expect(zombie.level).to.equal(1);
    expect(zombie.winCount).to.equal(0);
  });

  it("Should fail if the same address tries to create a second zombie directly", async function () {
    await zombieOwnership.createRandomZombie("Zombie 1");
    await expect(zombieOwnership.createRandomZombie("Zombie 2"))
      .to.be.revertedWith("Only one zombie per address in this demo!");
  });

  it("Should allow leveling up with fee", async function () {
    await zombieOwnership.createRandomZombie("Zombie 1");
    
    const levelUpFee = ethers.parseEther("0.001");
    await zombieOwnership.levelUp(0, { value: levelUpFee });

    const zombie = await zombieOwnership.zombies(0);
    expect(zombie.level).to.equal(2);
  });

  it("Should allow transfer of ownership via ERC721", async function () {
    await zombieOwnership.createRandomZombie("Zombie 1");
    
    expect(await zombieOwnership.ownerOf(0)).to.equal(owner.address);
    expect(await zombieOwnership.balanceOf(owner.address)).to.equal(1n);

    // Transfer from owner to addr1
    await zombieOwnership.transferFrom(owner.address, addr1.address, 0);

    expect(await zombieOwnership.ownerOf(0)).to.equal(addr1.address);
    expect(await zombieOwnership.balanceOf(owner.address)).to.equal(0n);
    expect(await zombieOwnership.balanceOf(addr1.address)).to.equal(1n);
  });
});
