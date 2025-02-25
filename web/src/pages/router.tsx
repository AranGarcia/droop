import React from 'react';
import { Route, Routes } from 'react-router';

import CharacterList from './CharacterList';

const Router = () => {
    return (
    <Routes>
        <Route path="/" element={<CharacterList />}></Route>
    </Routes>
    )
}

export default Router;