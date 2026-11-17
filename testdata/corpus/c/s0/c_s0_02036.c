
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1375_0 { KIND_EMPTY1375_0 = 0, KIND_FULL1375_0 = 32 };

union Cell1375_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1375_0(union Cell1375_0 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL1375_0 : KIND_EMPTY1375_0;
}
struct Node1375_1 {
    int value;
    struct Node1375_1 *next;
};

int length1375_1(const struct Node1375_1 *head) {
    int count = 0;
    const struct Node1375_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
int gcd1375_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1375_2(int value) {
    return gcd1375_2(value, 18) * 46;
}
