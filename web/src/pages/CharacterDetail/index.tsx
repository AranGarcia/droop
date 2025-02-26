import React, { useEffect, useState } from "react";
import { useParams } from "react-router";

import { characterType } from "../../types/character.types";
import { toTitleCase } from "../../shared/strings";

function getSignedNumber(value: number): string {
    var sign
    // negatives are already printed with sign 
    switch (true) {
        case value > 0:
            sign = "+";
            break;
        default:
            sign = "";
    }
    return `${sign}${value}`
}

function modifier(attribute: number): string {
    return getSignedNumber(Math.floor((attribute / 2) - 5))
}

function calculateProficiency(level: number): number {
    return 2 + Math.floor((level - 1)/4)
}

const Attribute = ({ name, attribute }) => {
    return <div>
        <b>{name}: </b> {attribute} ({modifier(attribute)})
    </div>
}

const CharacterDetail = () => {
    let { id } = useParams();
    const [character, setCharacter] = useState<characterType>({
        id: "",
        name: "",
        class: "",
        level: 0,

        max_health: 0,
        current_health: 0,
        temp_health: 0,
        strength: 0,
        dexterity: 0,
        constitution: 0,
        intelligence: 0,
        wisdom: 0,
        charisma: 0,
    });
    useEffect(() => {
        fetch(`http://localhost:8080/character/${id}`)
            .then((res) => {
                return res.json()
            })
            .then((data) => {
                setCharacter(data);
            });
    }, []);
    return (
        <div>
            <h1>{character.name}</h1>
            <h2> {toTitleCase(character.class)} level {character.level}</h2>

            <div>
                <b>Health</b>: {character.current_health} / {character.max_health}
            </div>

            <br/>
            <div>
                <b>Proficiency Bonus</b>: +{calculateProficiency(character.level)}
            </div>
            <br />

            <Attribute name="Strength" attribute={character.strength} />
            <Attribute name="Dexterity" attribute={character.dexterity} />
            <Attribute name="Constitution" attribute={character.constitution} />
            <Attribute name="Intelligence" attribute={character.intelligence} />
            <Attribute name="Wisdom" attribute={character.wisdom} />
            <Attribute name="Charisma" attribute={character.charisma} />
        </div>
    )
}

export default CharacterDetail;