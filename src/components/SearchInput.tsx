import React from 'react';

interface SearchInputProps {
    searchTerm: string;
    onSearchTermChange: (term: string) => void;
    onSearch: () => void;
}

const SearchInput: React.FC<SearchInputProps> = ({ searchTerm, onSearchTermChange, onSearch }) => {
    return (
        <div className="mb-4">
            <label className="block mb-2 text-sm font-medium text-gray-700 dark:text-gray-300" htmlFor="searchTerm">
                Search Term:
            </label>
            <div className="flex">
                <input
                    type="text"
                    id="searchTerm"
                    className="w-full px-4 py-2 border rounded-l-md focus:outline-none focus:ring focus:border-blue-300"
                    value={searchTerm}
                    onChange={(e) => onSearchTermChange(e.target.value)}
                />
                <button
                    onClick={onSearch}
                    className="px-4 py-2 text-white bg-blue-500 rounded-r-md hover:bg-blue-700"
                >
                    Search
                </button>
            </div>
        </div>
    );
}

export default SearchInput;
