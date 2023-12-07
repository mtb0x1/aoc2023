const fs = require("fs");

const CARDS = {
    A: 13,
    K: 12,
    Q: 11,
    J: 10,
    T: 9,
    9: 8,
    8: 7,
    7: 6,
    6: 5,
    5: 4,
    4: 3,
    3: 2,
    2: 1
};

function parseInput() {
    const data = fs.readFileSync("./input.txt", {
        encoding: "utf8",
        flag: "r"
    });
    return data.split("\n").map((line) => line.split(" "));
}

function comparePart1(a, b) {
    CARDS.J = 10;
    return compare(a, b);
}

function comparePart2(a, b) {
    CARDS.J = 0;
    return compare(a, b, true);
}

function compare(a, b, part2 = false) {
    let aCards = a[0],
        bCards = b[0];

    let aType = getType(aCards, part2);
    let bType = getType(bCards, part2);

    if (aType === bType) {
        for (let i = 0; i < 5; i++) {
            let astr = CARDS[aCards[i]],
                bstr = CARDS[bCards[i]];

            if (astr === bstr) continue;
            else return astr > bstr ? 1 : -1;
        }
    } else {
        return aType > bType ? 1 : -1;
    }
}

function getType(cards, part2 = false) {
    let type;

    if (part2 && cards.includes("J")) {
        for (let card in CARDS) {
            let potentialType = getHandType(
                buildHand(cards.replaceAll("J", card))
            );  

            type = !type || potentialType > type ? potentialType : type;
        }
    } else {
        type = getHandType(buildHand(cards));
    }

    return type;
}

function buildHand(cards) {
    let hand = new Map();

    for (let c of cards) {
        hand.set(c, hand.get(c) ? hand.get(c) + 1 : 1);
    }

    return hand;
}

function getHandType(hand) {
    if (hand.size === 1) {
        return 6;
    } else if (hand.size === 2) {
        let firstCardCount = hand.values().next().value;
        return firstCardCount === 4 || firstCardCount === 1 ? 5 : 4;
    } else if (hand.size === 3) {
        if (Array.from(hand.values()).includes(3)) {
            return 3;
        } else {
            return 2;
        }
    } else if (hand.size === 4) {
        return 1;
    } else if (hand.size === 5) {
        return 0;
    }
}

function solve(which) {
    return parseInput()
    .sort(which)
    .reduce((sum, [_, value], i) => sum + parseInt(value) * (i + 1), 0);
}

function solve_part1() {
    return solve(comparePart1);
}

function solve_part2() {
    return solve(comparePart2);
}

function main() {
    var p1 = solve_part1();
    var p2 = solve_part2();
    console.log(`part 1 = ${p1}\npart 2 = ${p2}`);
}

main();