import React from 'react';
import SearchTable from './components/SearchTable';
import DarkModeToggle from './components/DarkModeToggle';
import Logo from './components/Logo';

const App: React.FC = () => {
    return (
        <div className="container mx-auto mt-10">
            <div className="flex justify-between items-center mb-4">
                <Logo />
                <DarkModeToggle />
            </div>
            <SearchTable />
        </div>
    );
}



export default App;
