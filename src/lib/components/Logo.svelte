<div class="typing-container">
    <h3 class="typing-text"> My Notes Planner</h3>
</div>

<style>
    .typing-container {
        display: flex;
        align-items: center;
        justify-content: center;
        height: 100%;
        margin-top: -7px;
        margin-right: -5px;
    }

    .typing-text {
        color: #1d4ed8;
        font-size: 2rem;
        font-weight: 700;
        overflow: hidden;
        white-space: nowrap;
        border-right: 3px solid #2563eb;
        animation: typing 3s steps(15) forwards, blink .5s step-end 6 forwards;
    }

    @keyframes typing {
        from { width: 0; }
        to { width: 100%; }
    }

    @keyframes blink {
        from, to { border-color: transparent; }
        50% { border-color: #2563eb; }
    }
</style>