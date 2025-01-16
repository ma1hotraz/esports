import React from 'react';
import html2canvas from 'html2canvas';

const ScreenshotButton = () => {
    const handleScreenshot = () => {
        html2canvas(document.body).then((canvas) => {
            const link = document.createElement('a');
            link.href = canvas.toDataURL('image/png');
            link.download = `screenshot-${new Date().toISOString()}.png`;
            link.click();
        });
    };

    return (
        <button onClick={handleScreenshot}>Take Screenshot</button>
    );
};

export default ScreenshotButton;
