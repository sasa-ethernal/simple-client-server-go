// MessageForm.js
import React, { useState, useEffect } from 'react';
import getWeb3 from './web3';
import contractABI from './contracts/msg_server.json';
const contractAddress = '0x069EF88755F86365d68C283743f050FDc7fa1b53';

function MessageForm() {
    const [web3, setWeb3] = useState(null);
    const [contract, setContract] = useState(null);
    const [message, setMessage] = useState('');

    useEffect(() => {
        const initWeb3 = async () => {
        try {
            const web3Instance = await getWeb3();
            const contractInstance = new web3Instance.eth.Contract(
            contractABI,
            contractAddress
            );
            setWeb3(web3Instance);
            setContract(contractInstance);
        } catch (error) {
            console.error(error);
        }
        };
        initWeb3();
    }, []);

    const sendMessage = async () => {
        try {

        // Convert the message string to bytes
        const messageBytes = web3.utils.fromAscii(message);
        console.log(messageBytes)
        
        //await contract.methods.RequestPolicy('0x9793A09d7b0048f2910c2c49aD96194A4CBa949A', messageBytes).send({ from: web3.eth.defaultAccount });
        await contract.methods.RequestPolicy('0x9793A09d7b0048f2910c2c49aD96194A4CBa949A', messageBytes).send({ from: '0x322e233424BFeCB75664cE63973EacE0D01Bb796' });

        
        // Handle success
        } catch (error) {
        console.error(error);
        // Handle error
        }
    };

    return (
        <div>
        <h2>Interact with the Contract</h2>
        <input
            type="text"
            placeholder="Enter your message"
            value={message}
            onChange={(e) => setMessage(e.target.value)}
        />
        <button onClick={sendMessage}>Send Message</button>
        </div>
    );
}

export default MessageForm;
