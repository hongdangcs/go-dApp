const Web3 = require('web3');
const fs = require('fs');
const solc = require('solc');

const web3 = new Web3('https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID');

const source = fs.readFileSync('contracts/MyContract.sol', 'utf8');
const compiled = solc.compile(source, 1).contracts[':MyContract'];
const abi = JSON.parse(compiled.interface);
const bytecode = compiled.bytecode;

const deploy = async () => {
    const accounts = await web3.eth.getAccounts();
    const result = await new web3.eth.Contract(abi)
        .deploy({ data: bytecode })
        .send({ from: accounts[0], gas: '1000000' });

    console.log('Contract deployed to', result.options.address);
};

deploy();