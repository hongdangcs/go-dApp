const Web3 = require('web3');
const web3 = new Web3('https://mainnet.infura.io/v3/');

const contractAddress = '';
const abi = [];

const contract = new web3.eth.Contract(abi, contractAddress);

const getValue = async () => {
    const value = await contract.methods.getValue().call();
    console.log('Contract value:', value);
};

getValue();