import React from 'react';

const Logo: React.FC = () => {
    return (
        <div className="flex items-center">
            <img src="/zello-logo.svg" alt="Zello Logo" className="h-24 w-24 mr-2" />
            <label className="block mb-2 text-sm font-medium text-gray-700 dark:text-gray-300" htmlFor="searchTerm">
                Admin
            </label>
        </div>
    );
}

export default Logo;
