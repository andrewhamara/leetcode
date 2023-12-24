class Solution {
public:
    int minOperations(string s) {
        auto first = s[0];
        auto second = (first == '1' ?  '0' : '1');
        int result = 0;

        for (int i = 0; i < s.length(); i++) {
            auto c = s[i];
            if (i % 2 == 0 && c != first) result++;
            else if (i % 2 == 1 && c != second) result++;
        }

        int opposite = s.length() - result;
        return (opposite > result ? result : opposite);
    }
};
