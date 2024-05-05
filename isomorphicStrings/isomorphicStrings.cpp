#include <iostream>
#include <string>

using namespace std;

class Solution {
    public:
    bool isomorphic(string s, string t) {
        int m1[256] = {0};
        int m2[256] = {0};
        for (int i = 0; i < s.size(); i++) {
            if (m1[s[i]] != m2[t[i]]) return false;
            m1[s[i]] = i + 1;
            m2[t[i]] = i + 1;
        }
        return true;

    }
};

int main() {
    string s = "egg";
    string t = "add";
    Solution sol;
    cout << sol.isomorphic(s, t) << endl;
    return 0;
}
