// MessageForm.js
import React, { useState, useEffect } from "react";
import Web3 from "web3";
import MessageContract from "./contracts/msg_server.json";

function MessageForm() {
    const [web3, setWeb3] = useState(null);
    const [contract, setContract] = useState(null);
    const [messageData, setMessageData] = useState({
        address: '',
        transaction: '',
        policy: '',
        vm_address: '',
      });
    const [receivedMessage, setReceivedMessage] = useState('');

    // useEffect(() => {
    //     const initWeb3 = async () => {
    //         if (window.ethereum) {
    //             const web3Instance = new Web3(window.ethereum);
    //             try {
    //                 await window.ethereum.enable();
    //                 setWeb3(web3Instance);
    //                 const networkId = await web3Instance.eth.net.getId();
    //                 const deployedNetwork = MessageContract.networks[networkId];
    //                 const contractInstance = new web3Instance.eth.Contract(
    //                     MessageContract.abi,
    //                     deployedNetwork && deployedNetwork.address
    //                 );
    //                 setContract(contractInstance);
    //             } catch (error) {
    //                 console.error("User denied account access or something went wrong.");
    //             }
    //         }
    //     };

    //     initWeb3();
    // }, []);

    const handleSubmit = async (e) => {
        e.preventDefault();

        try {
            await contract.methods.RequestPolicy(messageData.address, messageData).send({ from: web3.eth.defaultAccount });
            setMessageData({
                address: '',
                transaction: '',
                policy: '',
                vm_address: '',
              });
        } catch (error) {
            console.error(error);
        }

        setMessageData("");
    };

    return (
        <div>
            <form onSubmit={handleSubmit}>
                <label>
                Address (Client IP):
                <input
                    type="text"
                    value={messageData.address}
                    onChange={(e) =>
                    setMessageData({ ...messageData, address: e.target.value })
                    }
                />
                </label>
                <label>
                Transaction:
                <input
                    type="text"
                    value={messageData.transaction}
                    onChange={(e) =>
                    setMessageData({ ...messageData, transaction: e.target.value })
                    }
                />
                </label>
                <label>
                Policy:
                <input
                    type="text"
                    value={messageData.policy}
                    onChange={(e) =>
                    setMessageData({ ...messageData, policy: e.target.value })
                    }
                />
                </label>
                <label>
                VM Address:
                <input
                    type="text"
                    value={messageData.vm_address}
                    onChange={(e) =>
                    setMessageData({ ...messageData, vm_address: e.target.value })
                    }
                />
                </label>
                <button type="submit">Send</button>
            </form>
            <div>
                <strong>Received Message:</strong> {receivedMessage}
            </div>
        </div>
    );
}

export default MessageForm;
