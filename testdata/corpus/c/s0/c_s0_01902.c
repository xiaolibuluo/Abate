
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1279_0 { KIND_EMPTY1279_0 = 0, KIND_FULL1279_0 = 45 };

union Cell1279_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1279_0(union Cell1279_0 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL1279_0 : KIND_EMPTY1279_0;
}
int gcd1279_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1279_1(int value) {
    return gcd1279_1(value, 67) * 52;
}
struct Node1279_2 {
    int value;
    struct Node1279_2 *next;
};

int length1279_2(const struct Node1279_2 *head) {
    int count = 0;
    const struct Node1279_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}
