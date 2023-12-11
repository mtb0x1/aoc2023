import java.io.File;
import java.io.FileNotFoundException;
import java.util.*;
 
public class d10 {
 
    public record Coordinates(int x, int y) {
         public int getX() {
            return x;
        }
        public int getY() {
            return y;
        }
    }
 
    public LinkedList<Coordinates> queue = new LinkedList<Coordinates>();
    public char[][] input = new char[140][140];
    public Map<Coordinates, Boolean> visited = new HashMap<Coordinates, Boolean>();

    public void AddtoQueue(Coordinates next) {
        if (!visited.get(next)) {
            visited.put(next, true);
            queue.add(next);
        }
    }
    public boolean nextRight(int x, int y) {
        char nextCharacter = input[x][y + 1];
        return "-J7".contains(nextCharacter + "");
    }
    public boolean nextLeft(int x, int y) {
        char nextCharacter = input[x][y - 1];
        return "-LF".contains(nextCharacter + "");
    }
    public boolean nextDown(int x, int y) {
        char nextCharacter = input[x + 1][y];
        return "|JL".contains(nextCharacter + "");
    }
    public boolean nextUp(int x, int y) {
        char nextCharacter = input[x - 1][y];
        return "|7F".contains(nextCharacter + "");
    }
    public boolean canGoUp(char character) {
        return "S|JL".contains(character + "");
    }
    public boolean canGoDown(char character) {
        return "S|7F".contains(character + "");
    }
    public boolean canGoLeft(char character) {
        return "S-J7".contains(character + "");
    }
    public boolean canGoRight(char character) {
        return "S-LF".contains(character + "");
    }
    //oh my god, one sec gonna go throw and be back.
    public void bfs(int x, int y) {
        Coordinates startCoordinates = new Coordinates(x, y);
        visited.put(startCoordinates, true);
        queue.add(startCoordinates);

        int inputHeight = input.length - 1;
        int inputWidth = input[0].length - 1;

        while (queue.size() != 0) {
            Coordinates coordinates = queue.poll();
            x = coordinates.getX();
            y = coordinates.getY();
            char character = input[x][y];

            if (x > 0 && canGoUp(character)
                    && nextUp(x, y)) {
                AddtoQueue(new Coordinates(x - 1, y));
            }

            if (x < inputHeight && canGoDown(character)
                    && nextDown(x, y)) {
                AddtoQueue(new Coordinates(x + 1, y));
            }

            if (y > 0 && canGoLeft(character)
                    && nextLeft(x, y)) {
                AddtoQueue(new Coordinates(x, y - 1));
            }

            if (y < inputWidth && canGoRight(character)
                    && nextRight(x, y)) {
                AddtoQueue(new Coordinates(x, y + 1));
            }
        }
    }
    public Coordinates parse() throws Exception{
        int startX = 0;
        int startY = 0;
        try (Scanner scanner = new Scanner(new File("input.txt"))) {
            int index = 0;
            while (scanner.hasNextLine()) {
                String line = scanner.nextLine();
                String[] split = line.split("");
                for (int i = 0; i < split.length; i++) {
                    input[index][i] = split[i].charAt(0);
                    visited.put(new Coordinates(index, i), false);
                    if (input[index][i] == 'S') {
                        startX = index;
                        startY = i;
                    }
                }
                index++;
            }
            scanner.close();
            
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }
       return new Coordinates(startX, startY);
    }

    public int solve_part1(int startX, int startY){
        bfs(startX, startY);
        int total = visited.values().stream().filter(value -> value).toArray().length;
        return total / 2;
    }

    public int getCount(int x, int y) {
        char[] row = input[x];
        int count = 0;
        int i = 0;
        while (i < y) {
            i++;
            Coordinates coordinates = new Coordinates(x, i);
            if (visited.get(coordinates) == false) {
                continue;
            }
            count += "JL|".contains(row[i] + "") ? 1 : 0;
        }
        return count;
    }

    public int solve_part2(int startX, int startY){
        bfs(startX, startY);

        int count = 0;
        for (int x = 0; x < input.length-1; x++) {
            for (int y = 0; y < input[x].length-1; y++) {
                Coordinates coordinates = new Coordinates(x, y);
                if (!visited.get(coordinates)) {
                    int inside = getCount(x, y);
                    count += (inside > 0 && inside % 2 == 1) ? 1 : 0;
                }
            }
        }
    
        return count;
    }

    public static void main(String[] args) throws Exception {
        d10 day = new d10();
        Coordinates xy = day.parse();
        int p1 = day.solve_part1(xy.x,xy.y);
        int p2 = day.solve_part2(xy.x,xy.y);
        System.out.printf("part 1 = %d\npart 2 = %d\n",p1,p2);
    }

}