
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1774_0 { KIND_EMPTY1774_0 = 0, KIND_FULL1774_0 = 78 };

union Cell1774_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1774_0(union Cell1774_0 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL1774_0 : KIND_EMPTY1774_0;
}
struct Node1774_1 {
    int value;
    struct Node1774_1 *next;
};

int length1774_1(const struct Node1774_1 *head) {
    int count = 0;
    const struct Node1774_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 85;
}
int gcd1774_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1774_2(int value) {
    return gcd1774_2(value, 66) * 92;
}
#if defined(__STDC_VERSION__)
int variant1774_3(int value, ...) {
    return value + 99;
}
#else
int variant1774_3(int value) {
    return value + 99;
}
#endif
