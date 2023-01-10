/* 
    MERGE SORT ALGORITHM O(n*logn)
*/

#include <bits/stdc++.h>

using namespace std;

vector<int> Merge (vector<int> left_part, vector<int> right_part) {

    vector<int>result;

    int left_idx = 0, right_idx = 0;

    while (left_idx < left_part.size() && right_idx < right_part.size()) {

        if (left_part[left_idx] <= right_part[right_idx]) {
            result.push_back(left_part[left_idx]);
            left_idx++;
        } else {
            result.push_back(right_part[right_idx]);
            right_idx++;
        }
    }

    if (left_idx < left_part.size()) {
        result.insert(result.end(), left_part.begin() + left_idx, left_part.end());
    }

    if (right_idx < right_part.size()) {
        result.insert(result.end() , right_part.begin() + right_idx, right_part.end());
    }

    return result;
}

vector<int> Merge_sort(vector<int> arr) {

    if (arr.size() <= 1) {
        return arr;
    }

    int mid = arr.size() / 2;

    vector<int> left_part(arr.begin(), arr.begin() + mid);
    vector<int> right_part(arr.begin()+mid, arr.end());

    left_part = Merge_sort(left_part);
    right_part = Merge_sort(right_part);

    return Merge(left_part, right_part);
}

void solve() {
    int n;
    cin >> n;

    vector<int> arr(n, 0);

    for (auto &i: arr) {
        cin >> i;
    }

    vector<int> sorted_arr = Merge_sort(arr);

    for (auto &i: sorted_arr) {
        cout << i << " ";
    }

}

int main() {
    int t = 1;
    while (t--) {
        solve();
    }

    return 0;
}