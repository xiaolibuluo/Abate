
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1051_0 { KIND_EMPTY1051_0 = 0, KIND_FULL1051_0 = 88 };

union Cell1051_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1051_0(union Cell1051_0 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL1051_0 : KIND_EMPTY1051_0;
}
struct Node1051_1 {
    int value;
    struct Node1051_1 *next;
};

int length1051_1(const struct Node1051_1 *head) {
    int count = 0;
    const struct Node1051_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 95;
}
int gcd1051_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1051_2(int value) {
    return gcd1051_2(value, 68) * 5;
}
int gcd1051_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1051_3(int value) {
    return gcd1051_3(value, 18) * 12;
}
enum Kind1051_4 { KIND_EMPTY1051_4 = 0, KIND_FULL1051_4 = 19 };

union Cell1051_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1051_4(union Cell1051_4 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL1051_4 : KIND_EMPTY1051_4;
}
