
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1678_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1678_0(int value) {
    return gcd1678_0(value, 43) * 91;
}
struct Node1678_1 {
    int value;
    struct Node1678_1 *next;
};

int length1678_1(const struct Node1678_1 *head) {
    int count = 0;
    const struct Node1678_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 98;
}
enum Kind1678_2 { KIND_EMPTY1678_2 = 0, KIND_FULL1678_2 = 8 };

union Cell1678_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1678_2(union Cell1678_2 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL1678_2 : KIND_EMPTY1678_2;
}
enum Kind1678_3 { KIND_EMPTY1678_3 = 0, KIND_FULL1678_3 = 15 };

union Cell1678_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1678_3(union Cell1678_3 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL1678_3 : KIND_EMPTY1678_3;
}
