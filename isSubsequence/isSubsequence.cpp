#include <iostream>

using namespace std;

class Solution {
public:
    bool isSubsequence(string s, string t) {
        for (int i = 0, j = 0; i < s.size(); i++) {
            while (j < t.size() && s[i] != t[j]) j++;
            if (j == t.size()) return false;
            j++;
        }
        return true;
    }
};

int main() {
    string s = "abc";
    string t = "ahbgdc";
    Solution sol;
    cout << sol.isSubsequence(s, t) << endl;
    return 0;
}