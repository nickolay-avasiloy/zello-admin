import React from 'react';
import { TrashIcon } from '@heroicons/react/24/solid';

interface User {
    username: string;
    company: string;
    phoneNumber: string;
}

interface UserTableProps {
    users: User[];
    onDelete: (username: string) => void;
}

const UserTable: React.FC<UserTableProps> = ({ users, onDelete }) => {
    return (
        <table className="min-w-full bg-white border border-gray-300 dark:bg-gray-700 dark:border-gray-600">
            <thead>
            <tr>
                <th className="px-4 py-2 text-left bg-gray-200 border-b border-gray-300 dark:bg-gray-600 dark:border-gray-500">Username</th>
                <th className="px-4 py-2 text-left bg-gray-200 border-b border-gray-300 dark:bg-gray-600 dark:border-gray-500">Company</th>
                <th className="px-4 py-2 text-left bg-gray-200 border-b border-gray-300 dark:bg-gray-600 dark:border-gray-500">Phone Num.</th>
                <th className="px-4 py-2 text-left bg-gray-200 border-b border-gray-300 dark:bg-gray-600 dark:border-gray-500"></th>
            </tr>
            </thead>
            <tbody>
            {users.map((user, index) => (
                <tr key={index} className="border-b dark:border-gray-500">
                    <td className="px-4 py-2">{user.username}</td>
                    <td className="px-4 py-2">{user.company}</td>
                    <td className="px-4 py-2">{user.phoneNumber}</td>
                    <td className="px-4 py-2">
                        <button
                            onClick={() => onDelete(user.username)}
                            className="flex items-center px-2 py-1 text-white bg-red-500 rounded hover:bg-red-700"
                        >
                            <TrashIcon className="w-5 h-5" />
                            <span className="ml-1">Delete</span>
                        </button>
                    </td>
                </tr>
            ))}
            </tbody>
        </table>
    );
}

export default UserTable;
