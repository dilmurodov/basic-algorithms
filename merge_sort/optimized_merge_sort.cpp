/* 
    MERGE SORT ALGORITHM O(n*logn)
*/

#include <bits/stdc++.h>

using namespace std;

void merge(vector<int> &arr,int l, int r, int m) {

    int i = l, j = m+1, k = l;
    vector<int>temp(arr.size(), 0);

    // add min elements
    while (i <= m && j <= r) {
        if (arr[i] < arr[j]) {
            temp[k++] = arr[i++];
        } else {
            temp[k++] = arr[j++];
        }
    }

    // add remaining elements from left_sub_array
    while (i <= m) {
        temp[k++] = arr[i++];
    }

    // add remaining elements from right_sub_array
    while (j <= r) {
        temp[k++] = arr[j++];
    }

    // copy elements to arr
    for (int x = l; x <= r; x++) {
        arr[x] = temp[x];
    }

    return;
}

void merge_sort(vector<int> &arr, int l, int r) {

    if (l < r) {
        int m = l + (r -l)/2;

        merge_sort(arr, l, m);
        merge_sort(arr, m+1, r);
        merge(arr, l, r, m);
    }

    return;
}

void solve() {

    int n; cin >> n;
    vector<int> arr(n, 0);

    for (auto &i: arr) {
        cin >> i;
    }

    merge_sort(arr, 0, n-1);

    for (auto &i :arr) {
        cout << i << " ";
    }

    return;
}

int main() {

    int t = 1;

    while(t--) {
        solve();
    }

    return 0;
}