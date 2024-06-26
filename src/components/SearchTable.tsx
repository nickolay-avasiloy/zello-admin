import React, { useState } from 'react';
import SearchInput from './SearchInput';
import UserTable from './UserTable';

interface User {
    username: string;
    company: string;
    phoneNumber: string;
}

const SearchTable: React.FC = () => {
    const [searchTerm, setSearchTerm] = useState('');
    const [users, setUsers] = useState<User[]>([
        { username: 'Joshua', company: 'Zello', phoneNumber: '512-937-5555' },
        { username: 'Joshy', company: 'Craigs', phoneNumber: '512-879-5555' },
        { username: 'Simon', company: 'Joshi LLC', phoneNumber: '279-103-9876' },
    ]);

    const handleDelete = (username: string) => {
        setUsers(users.filter(user => user.username !== username));
    };

    const filteredUsers = users.filter(user =>
        user.username.toLowerCase().includes(searchTerm.toLowerCase())
    );

    const handleSearch = () => {
        // Implement search logic if needed
    };

    return (
        <div className="p-6 bg-white dark:bg-gray-800 dark:text-white">
            <SearchInput
                searchTerm={searchTerm}
                onSearchTermChange={setSearchTerm}
                onSearch={handleSearch}
            />
            <UserTable users={filteredUsers} onDelete={handleDelete} />
        </div>
    );
}

export default SearchTable;
