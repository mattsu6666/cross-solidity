const IBCHost = artifacts.require("IBCHost");
const MockClient = artifacts.require("MockClient");
const IBCClient = artifacts.require("IBCClient");
const IBCConnection = artifacts.require("IBCConnection");
const IBCChannel = artifacts.require("IBCChannel");
const IBCHandler = artifacts.require("IBCHandler");
const IBCMsgs = artifacts.require("IBCMsgs");
const IBCCommitment = artifacts.require("IBCCommitment");
const CrossSimpleModule = artifacts.require("CrossSimpleModule");
const MockCrossContract = artifacts.require("MockCrossContract");

const deployCore = async (deployer) => {
  await deployer.deploy(IBCCommitment);
  await deployer.link(IBCCommitment, [IBCClient, IBCConnection, IBCChannel, IBCHost, IBCHandler]);

  await deployer.deploy(IBCMsgs);
  await deployer.link(IBCMsgs, [
    IBCClient,
    IBCConnection,
    IBCChannel,
    IBCHandler
  ]);

  await deployer.deploy(IBCClient);
  await deployer.link(IBCClient, [IBCHandler, IBCConnection, IBCChannel, IBCCommitment]);

  await deployer.deploy(IBCConnection);
  await deployer.link(IBCConnection, [IBCHandler, IBCChannel, IBCCommitment]);

  await deployer.deploy(IBCChannel);
  await deployer.link(IBCChannel, [IBCHandler, IBCCommitment]);

  await deployer.deploy(MockClient, IBCHandler.address);

  await deployer.deploy(IBCHost);
  await deployer.deploy(IBCHandler, IBCHost.address);
};

const deployApp = async (deployer) => {
  console.log("deploying app contracts");

  await deployer.deploy(MockCrossContract);
  await deployer.deploy(CrossSimpleModule, IBCHost.address, IBCHandler.address, MockCrossContract.address, true);
};

module.exports = async function (deployer) {
  await deployCore(deployer);
  await deployApp(deployer);
};
