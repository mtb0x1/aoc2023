const fs = require("fs");
 
function parseInput() {
    let input = fs.readFileSync("./input.txt", {
        encoding: "utf8",
        flag: "r"
    });
    return input.split("\n");
}

//https://en.wikipedia.org/wiki/Least_common_multiple
function LeastCommonMultiple(nums){
    let factors = {};
    
    // Find prime factors for each number in the array
    for(let i = 0; i < nums.length; i++) {
        x = 2;
        while(x <= nums[i]) {
            while(nums[i]%x == 0) {
                nums[i] = nums[i]/x;
                factors[x] =  (factors[x] + 1) || 1;
            }
            x++;
        }
    }

    let LCM = 1;
    // Calculate LCM using prime factorization
    Object.keys(factors).forEach(factor => {
        LCM = LCM * Number.parseInt(factor)**factors[factor]; 
    });
    return LCM;
}

function solve_part1(data){
    let directions = data[0];
    let network = {};
    for(let i = 2; i < data.length; i++) {
        const node = data[i].replace(/[()]/g, "").split(" = ");
        network[node[0]] = {L: node[1].split(", ")[0], R: node[1].split(", ")[1]};
        
    }

    let count = 0;
    
    let next = "AAA";
    for(let i = 0; i <= directions.length; i++) {
        if(i == directions.length) {
            i=0;
        }
        count++;
        next = network[next][directions[i]];
        if(next === "ZZZ") {
            break;
        }
    }
    return count;
}

function solve_part2(data){
    let directions = data[0];
    let nexts = [];

    let steps = [];
    let distances = [];

    const network = {};
    for(let i = 2; i < data.length; i++) {
        const node = data[i].replace(/[()]/g, "").split(" = ");
        network[node[0]] = {L: node[1].split(", ")[0], R: node[1].split(", ")[1]};
        
    }

    Object.keys(network).forEach(key => {
        if(key[2] === "A") {
            nexts.push(key);
        }
    });


    for(let j = 0; j < nexts.length; j++) {
        steps.push(0);
        for(let i = 0; i <= directions.length; i++) {
            if(i == directions.length) {
                i=0;
            }
            steps[j]++;
            nexts[j] = network[nexts[j]][directions[i]];
            if(nexts[j][2] === "Z") {
                distances.push(steps[j]/directions.length);
                break;
            }
        }
    }

    return LeastCommonMultiple(distances)*directions.length;
}

function main() {
    const data = parseInput();
    let p1 = solve_part1(data);
    let p2 = solve_part2(data,[]);
    console.log(`part 1 = ${p1}\npart 2 = ${p2}`);
}

main();