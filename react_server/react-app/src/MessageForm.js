// MessageForm.js
import React, { useState, useEffect } from 'react';

function MessageForm() {
    const [messageData, setMessageData] = useState("");
    const [receivedMessage, setReceivedMessage] = useState('');

    useEffect(() => {
        // Connecto to BC
        // Subscribe to events
    }, []);

    const handleSubmit = (e) => {
        e.preventDefault();

        // send to BC

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
