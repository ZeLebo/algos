#include <iostream>

using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(): val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(NULL) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
public:
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        ListNode* head = new ListNode();
        ListNode* tail = head;
        while (list1 != nullptr && list2 != nullptr) {
            if (list1->val < list2->val) {
                tail->next = list1;
                list1 = list1->next;
            } else {
                tail->next = list2;
                list2 = list2->next;
            }
            tail = tail->next;
        }
        if (list1 != nullptr) {
            tail->next = list1;
        }
        if (list2 != nullptr) {
            tail->next = list2;
        }
        return head->next;
    }
};

int main() {
    Solution s;
    // list1 = {}
    // list2 = {0}
    ListNode* list1 = new ListNode();
    ListNode* list2 = new ListNode(0);
    ListNode* result = s.mergeTwoLists(list1, list2);
    while (result != nullptr) {
        cout << result->val << " ";
        result = result->next;
    }
    return 0;
}