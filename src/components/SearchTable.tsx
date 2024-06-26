import React, {useEffect, useState} from 'react';
import SearchInput from './SearchInput';
import UserTable, {User} from './UserTable';

const API_URL = process.env.REACT_APP_API_URL;

const SearchTable: React.FC = () => {
    const [searchTerm, setSearchTerm] = useState('');
    const [users, setUsers] = useState<User[]>([]);

    useEffect(() => {
        const fetchInitialUsers = async () => {
            try {
                const response = await fetch(`${API_URL}/api/users`);
                const data = await response.json();
                setUsers(data);
            } catch (error) {
                console.error('Error fetching initial users:', error);
            }
        };

        fetchInitialUsers()
    }, []);

    const handleDelete = async (uuid: string) => {
        try {
            const response = await fetch(`${API_URL}/api/users/${uuid}`, {
                method: 'DELETE'
            });
            if (!response.ok) {
                throw new Error(`Error: ${response.statusText}`);
            }
            setUsers(users.filter(user => user.uuid !== uuid));
        } catch (error) {
            console.error('Error deleting user:', error);
        }
    };

    const handleSearch = async () => {
        try {
            const response = await fetch(`${API_URL}/api/search?q=${searchTerm}`);
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
