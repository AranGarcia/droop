import React from 'react';
import { useState, useEffect } from 'react';
import { Link, generatePath } from "react-router";


import { characterType } from '../../types/character.ts'

function CharacterRow({ character }: { character: characterType }) {
    return (
        <tr key={character.id}>
            <td>
                <Link to={generatePath("/character/:id", {id: character.id})}>{character.name}</Link>
            </td>
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