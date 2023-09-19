import Web3 from 'web3';

const getWeb3 = () => {
  return new Promise(async (resolve, reject) => {
    if (window.ethereum) {
      const web3 = new Web3(window.ethereum);
      try {
        await window.ethereum.enable();
        resolve(web3);
      } catch (error) {
        reject(error);
      }
    } else if (window.web3) {
      const web3 = new Web3(web3.currentProvider);
      resolve(web3);
    } else {
      reject('Please install MetaMask or use a dapp browser.');
    }
  });
};

export default getWeb3;
