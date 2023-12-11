class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        std::map<int, int> visited;
        std::vector<int> result;

        int length = nums.size();
        for (int i = 0; i < length; i++) {
            int cur = target - nums[i];
            if (visited.find(cur) != visited.end()) {
                result.push_back(visited[cur]);
                result.push_back(i);
                return result;
            }
            visited[nums[i]] = i;
        }
        return result; // empty vector
     }
};
