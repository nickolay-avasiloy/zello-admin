import React from 'react';

const Logo: React.FC = () => {
    return (
        <div className="flex items-center">
            <img src="/zello-logo.svg" alt="Zello Logo" className="h-24 w-24 mr-2" />
        </div>
    );
}

export default Logo;
