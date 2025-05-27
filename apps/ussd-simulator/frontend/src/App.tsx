import {useEffect, useState} from 'react';

export default function App() {
    const [input, setInput] = useState("");
    const handleDelete = () => setInput((prev) => prev.slice(0, -1));
    return (
        <div style={{
              width: '360px',
              height: '720px',
              margin: '20px',
              background: 'white',
              borderRadius: '40px',
              padding: '20px 10px 10px',
              boxShadow: '0 0 20px rgba(0,0,0,0.2)',
              display: 'flex',
              flexDirection: 'column',
              position: 'relative',
            }}>
            <div style={{
                  color: '#666',
                  display: 'flex',
                  padding: '8px',
                  justifyContent: 'space-between',
                }}>
                <div>
                    9:41 &nbsp;
                </div>
                <div style={{border: 'none', background: "none", outline: 'none'}}>
                    Africell 📶
                </div>
            </div>
        </div>
    );
}