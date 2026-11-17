
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1527_0 {
    int value;
    struct Node1527_0 *next;
};

int length1527_0(const struct Node1527_0 *head) {
    int count = 0;
    const struct Node1527_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
#if defined(__STDC_VERSION__)
int variant1527_1(int value, ...) {
    return value + 75;
}
#else
int variant1527_1(int value) {
    return value + 75;
}
#endif
enum Kind1527_2 { KIND_EMPTY1527_2 = 0, KIND_FULL1527_2 = 82 };

union Cell1527_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1527_2(union Cell1527_2 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL1527_2 : KIND_EMPTY1527_2;
}
int gcd1527_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1527_3(int value) {
    return gcd1527_3(value, 56) * 89;
}
