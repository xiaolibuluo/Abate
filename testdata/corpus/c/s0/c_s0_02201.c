
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1481_0 { KIND_EMPTY1481_0 = 0, KIND_FULL1481_0 = 52 };

union Cell1481_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1481_0(union Cell1481_0 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL1481_0 : KIND_EMPTY1481_0;
}
int gcd1481_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1481_1(int value) {
    return gcd1481_1(value, 17) * 59;
}
struct Node1481_2 {
    int value;
    struct Node1481_2 *next;
};

int length1481_2(const struct Node1481_2 *head) {
    int count = 0;
    const struct Node1481_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 66;
}
#if defined(__STDC_VERSION__)
int variant1481_3(int value, ...) {
    return value + 73;
}
#else
int variant1481_3(int value) {
    return value + 73;
}
#endif
