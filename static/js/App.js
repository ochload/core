import React from 'react';
import { createRoot } from 'react-dom/client';
import Modal from "./Components/Modal";
import "./../css/index.css";

function App() {
  return (
    <div className="bg-white dark:bg-slate-900">
      <Modal />
    </div>
  )
}

const app = document.getElementById('app');
const root = createRoot(app);
root.render(
    <React.StrictMode>
        <App/>
    </React.StrictMode>
);