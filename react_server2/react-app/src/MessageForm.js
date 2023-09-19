import React, { useState, useEffect } from "react";
import Web3 from "web3";
import MessageContract from "./contracts/msg_server.json";
import { Container, Form, Button, ListGroup } from "react-bootstrap";

function MessagingApp() {
    const [web3, setWeb3] = useState(null);
    const [contract, setContract] = useState(null);
    const [messages, setMessages] = useState([]);
    const [messageContent, setMessageContent] = useState("");
    const [recipient, setRecipient] = useState("");

    useEffect(() => {
        const initWeb3 = async () => {
            if (window.ethereum) {
                const web3Instance = new Web3(window.ethereum);
                try {
                    await window.ethereum.enable();
                    setWeb3(web3Instance);
                    const networkId = await web3Instance.eth.net.getId();
                    const deployedNetwork = MessageContract.networks[networkId];
                    const contractInstance = new web3Instance.eth.Contract(
                        MessageContract.abi,
                        deployedNetwork && deployedNetwork.address
                    );
                    setContract(contractInstance);
                } catch (error) {
                    console.error("User denied account access or something went wrong.");
                }
            }
        };

        initWeb3();
    }, []);

    const sendMessage = async () => {
        try {
            await contract.methods.sendMessage(recipient, messageContent).send({ from: web3.eth.defaultAccount });
            setMessageContent("");
            loadMessages();
        } catch (error) {
            console.error(error);
        }
    };

    const loadMessages = async () => {
        const messageCount = await contract.methods.getMessageCount(recipient).call();
        const messagePromises = [];
        for (let i = 0; i < messageCount; i++) {
            messagePromises.push(contract.methods.messages(recipient, i).call());
        }
        const newMessages = await Promise.all(messagePromises);
        setMessages(newMessages);
    };

    return (
        <Container>
            <h1>Ethereum Messaging App</h1>
            <Form>
                <Form.Group controlId="recipient">
                    <Form.Label>Recipient Address:</Form.Label>
                    <Form.Control
                        type="text"
                        placeholder="Recipient Address"
                        value={recipient}
                        onChange={(e) => setRecipient(e.target.value)}
                    />
                </Form.Group>
                <Form.Group controlId="messageContent">
                    <Form.Label>Message:</Form.Label>
                    <Form.Control
                        as="textarea"
                        rows={3}
                        value={messageContent}
                        onChange={(e) => setMessageContent(e.target.value)}
                    />
                </Form.Group>
                <Button variant="primary" onClick={sendMessage}>
                    Send Message
                </Button>
            </Form>
            <h2>Messages:</h2>
            <ListGroup>
                {messages.map((message, index) => (
                    <ListGroup.Item key={index}>
                        <strong>Sender:</strong> {message.sender}<br />
                        <strong>Content:</strong> {message.content}
                    </ListGroup.Item>
                ))}
            </ListGroup>
        </Container>
    );
}

export default MessagingApp;
