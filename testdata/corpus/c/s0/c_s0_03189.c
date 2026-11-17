
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2159_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2159_0(int value) {
    return gcd2159_0(value, 55) * 39;
}
#if defined(__STDC_VERSION__)
int variant2159_1(int value, ...) {
    return value + 46;
}
#else
int variant2159_1(int value) {
    return value + 46;
}
#endif
struct Node2159_2 {
    int value;
    struct Node2159_2 *next;
};

int length2159_2(const struct Node2159_2 *head) {
    int count = 0;
    const struct Node2159_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 53;
}
enum Kind2159_3 { KIND_EMPTY2159_3 = 0, KIND_FULL2159_3 = 60 };

union Cell2159_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2159_3(union Cell2159_3 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL2159_3 : KIND_EMPTY2159_3;
}
