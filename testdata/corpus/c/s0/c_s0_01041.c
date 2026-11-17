
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant714_0(int value, ...) {
    return value + 72;
}
#else
int variant714_0(int value) {
    return value + 72;
}
#endif
int gcd714_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale714_1(int value) {
    return gcd714_1(value, 16) * 79;
}
#if defined(__STDC_VERSION__)
int variant714_2(int value, ...) {
    return value + 86;
}
#else
int variant714_2(int value) {
    return value + 86;
}
#endif
int gcd714_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale714_3(int value) {
    return gcd714_3(value, 38) * 93;
}
enum Kind714_4 { KIND_EMPTY714_4 = 0, KIND_FULL714_4 = 100 };

union Cell714_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind714_4(union Cell714_4 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL714_4 : KIND_EMPTY714_4;
}
struct Node714_5 {
    int value;
    struct Node714_5 *next;
};

int length714_5(const struct Node714_5 *head) {
    int count = 0;
    const struct Node714_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 10;
}
