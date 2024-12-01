import React from 'react';
import '../css/Modal.css';

const Modal = ({ showModal, setShowModal, handleSubmit, children }) => {
    if (!showModal) return null;

    const closeModal = () => setShowModal(false);

    return (
        <div className="modal-overlay">
            <div className="modal-content">
                <button className="close-btn" onClick={closeModal}>X</button>
                <form onSubmit={handleSubmit}>
                    {children} {/* Render the form content dynamically */}
                    <button type="submit">Submit</button>
                </form>
            </div>
        </div>
    );
};

export default Modal;
