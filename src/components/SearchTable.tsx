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
    const [users, setUsers] = useState<User[]>([]);

    const handleDelete = (username: string) => {
        setUsers(users.filter(user => user.username !== username));
    };

    const handleSearch = async () => {
        try {
            const response = await fetch(`${process.env.REACT_APP_API_URL}/api/search?q=${searchTerm}`);
            const data = await response.json();
            setUsers(data);
        } catch (error) {
            console.error('Error fetching data:', error);
        }
    };

    return (
        <div className="p-6 bg-white dark:bg-gray-800 dark:text-white">
            <SearchInput
                searchTerm={searchTerm}
                onSearchTermChange={setSearchTerm}
                onSearch={handleSearch}
            />
            <UserTable users={users} onDelete={handleDelete} />
        </div>
    );
}

export default SearchTable;
