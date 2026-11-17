
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1327_0 { KIND_EMPTY1327_0 = 0, KIND_FULL1327_0 = 87 };

union Cell1327_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1327_0(union Cell1327_0 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL1327_0 : KIND_EMPTY1327_0;
}
struct Node1327_1 {
    int value;
    struct Node1327_1 *next;
};

int length1327_1(const struct Node1327_1 *head) {
    int count = 0;
    const struct Node1327_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 94;
}
int gcd1327_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1327_2(int value) {
    return gcd1327_2(value, 48) * 101;
}
