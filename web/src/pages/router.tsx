import React from 'react';
import { Route, Routes } from 'react-router';

import CharacterList from './CharacterList';
import CharacterDetail from './CharacterDetail';

const Router = () => {
    return (
    <Routes>
        <Route path="/" element={<CharacterList />}></Route>
        <Route path="/character/:id" element={<CharacterDetail />}></Route>
    </Routes>
    )
}

export default Router;