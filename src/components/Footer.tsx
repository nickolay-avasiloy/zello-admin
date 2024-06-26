import React from 'react';
import { CpuChipIcon, DocumentMagnifyingGlassIcon } from '@heroicons/react/24/solid';

const Footer: React.FC = () => {
    return (
        <footer className="flex flex-col items-center space-y-2 p-4 bg-gray-100 dark:bg-gray-800">
            <div className="flex space-x-4">
                <a
                    href="https://swagger.io/"
                    target="_blank"
                    rel="noopener noreferrer"
                    className="flex items-center text-blue-500 hover:text-blue-700 dark:text-blue-400 dark:hover:text-blue-600"
                >
                    <DocumentMagnifyingGlassIcon className="w-5 h-5 mr-1" />
                    Swagger
                </a>
                <a
                    href="https://nick.engineering/"
                    target="_blank"
                    rel="noopener noreferrer"
                    className="flex items-center text-blue-500 hover:text-blue-700 dark:text-blue-400 dark:hover:text-blue-600"
                >
                    <CpuChipIcon className="w-5 h-5 mr-1" />
                    Nick Engineering
                </a>
            </div>
        </footer>
    );
}

export default Footer;
