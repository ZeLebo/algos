#include <iostream>
#include <unordered_map>
#include <vector>

using namespace std;


class Solution {
    public:
        vector<int> twoSum(vector<int>& nums, int target) {
            // map[num] = index
            unordered_map<int, int> map;
            for (int i = 0; i < nums.size(); i++) {
                if (map.find(target - nums[i]) != map.end()) {
                    return {map[target - nums[i]], i};
                }
                map[nums[i]] = i;
            }
            return {};
        }
};

int main() {
    Solution solution;
    vector<int> nums = {2, 7, 15, 22};
    int target = 9;
    vector<int> result = solution.twoSum(nums, target);
    cout << "Indices: [" << result[0] << ", " << result[1] << "]" << endl;
    return 0;
}
