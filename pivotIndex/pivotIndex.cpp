#include <stdio.h>
#include <vector>
#include <iostream>

using namespace std;

class Solution {
    public:
        int pivotIndex(vector<int>& nums) {
            int rightSum = 0;
            for (int i = 0; i < nums.size(); i++) {
                rightSum += nums[i];
            }
            int leftSum = 0;
            for (int i = 0; i < nums.size(); i++) {
                rightSum -= nums[i];
                if (leftSum == rightSum) {
                    return i;
                }
                leftSum += nums[i];
            }

            return -1;
        }
};

int main() {
    Solution s;
    vector<int> nums = {1,7,3,6,5,6};
    int result = s.pivotIndex(nums);
    cout << result << endl;
    return 0;
}