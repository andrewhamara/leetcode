class Solution {
public:
    bool isPathCrossing(string path) {
        int x = 0, y = 0;
        set<pair<int,int>> s;

        s.insert({0,0});

        for (auto direction : path) {
            if (direction == 'N')      y++;
            else if (direction == 'S') y--;
            else if (direction == 'W') x--;
            else if (direction == 'E') x++;

            if (s.find({x,y}) != s.end()) return true;
            else s.insert({x,y});
        }
        return false;
    }
};
