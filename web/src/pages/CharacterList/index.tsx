import React from 'react';
import { useState, useEffect } from 'react';

import { characterType } from '../../types/character.types'

function CharacterRow ({character}: {character: characterType}) {
    return (
        <tr key={character.id}>
            <td>{character.name}</td>
            <td>{character.class}</td>
            <td>{character.level}</td>
        </tr>)
}

const Rows = () => {
    const [characters, setCharacters] = useState([]);
    useEffect(() => {
        fetch('http://localhost:8080/characters')
            .then((res) => {
                return res.json()
            })
            .then((data) => {
                setCharacters(data.characters);
            });
    }, []);
    if (characters.length == 0) {
        return null;
    }

    const rows = characters.map((c: characterType) => {
        return <CharacterRow
            character={c}
            key={c.id}
            />;
    });
    return (
        <table>
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Class</th>
                    <th>Level</th>
                </tr>
            </thead>
            <tbody>
                {rows}
            </tbody>
        </table>)
}

const CharacterList = () => {
    return (<div>
        <h1>Characters</h1>
        <Rows />
    </div>)
}

export default CharacterList;