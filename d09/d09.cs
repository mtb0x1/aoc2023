int[][] parse(){
    return File.ReadAllLines("input.txt")
        .Select(x => x.Split(' ').Select(x => int.Parse(x.Trim())).ToArray())
        .ToArray();
}

int GetNextValue(int[] data, bool next){
    if(data.Distinct().Count() == 1){
        return data.First();
    }

    List<int> differences = data.Skip(1).Select((value, index) => value - data[index]).ToList();

    if(next){
        return data.Last() + GetNextValue(differences.ToArray(), next);
    }else{
        return data.First() - GetNextValue(differences.ToArray(), next);
    }
}

int solve(int[][] data,bool next){
   return data.Sum(x => GetNextValue(x, next));
}

int solve_part1(int[][] data){
   return solve(data, true);
}

int solve_part2(int[][] data){
   return solve(data, false);
}

void main(){
    int[][] data = parse();
    int p1 = solve_part1(data);
    int p2 = solve_part2(data);
    Console.WriteLine("part 1 = {0}\npart 2 = {1}",p1,p2);
}

main();