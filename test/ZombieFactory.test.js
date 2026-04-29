import { expect } from "chai";
import pkg from "hardhat";
const { ethers } = pkg;

describe("ZombieFactory", function () {
  it("Should create a new zombie and emit an event", async function () {
    const ZombieFactory = await ethers.getContractFactory("ZombieFactory");
    const zombieFactory = await ZombieFactory.deploy();
    await zombieFactory.waitForDeployment();

    const zombieName = "Survivor";
    await expect(zombieFactory.createRandomZombie(zombieName))
      .to.emit(zombieFactory, "NewZombie");

    const zombies = await zombieFactory.zombies(0);
    expect(zombies.name).to.equal(zombieName);
  });

  it("Should fail if the same address tries to create a second zombie", async function () {
    const ZombieFactory = await ethers.getContractFactory("ZombieFactory");
    const zombieFactory = await ZombieFactory.deploy();
    await zombieFactory.waitForDeployment();

    await zombieFactory.createRandomZombie("Zombie 1");
    await expect(zombieFactory.createRandomZombie("Zombie 2"))
      .to.be.revertedWith("Only one zombie per address in this demo!");
  });
});
