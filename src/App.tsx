import React from 'react';
import './App.css';
import SearchTable from './components/SearchTable';
import DarkModeToggle from './components/DarkModeToggle';

const App: React.FC = () => {
    return (
        <div className="container mx-auto mt-10">
            <div className="flex justify-end mb-4">
                <DarkModeToggle />
            </div>
            <SearchTable />
        </div>
    );
}


export default App;
