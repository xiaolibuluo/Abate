
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2524_0 {
    int value;
    struct Node2524_0 *next;
};

int length2524_0(const struct Node2524_0 *head) {
    int count = 0;
    const struct Node2524_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 31;
}
int gcd2524_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2524_1(int value) {
    return gcd2524_1(value, 59) * 38;
}
enum Kind2524_2 { KIND_EMPTY2524_2 = 0, KIND_FULL2524_2 = 45 };

union Cell2524_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2524_2(union Cell2524_2 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL2524_2 : KIND_EMPTY2524_2;
}
int scan2524_3(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 52;
}
int gcd2524_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2524_4(int value) {
    return gcd2524_4(value, 31) * 59;
}
int gcd2524_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2524_5(int value) {
    return gcd2524_5(value, 42) * 66;
}
#if defined(__STDC_VERSION__)
int variant2524_6(int value, ...) {
    return value + 73;
}
#else
int variant2524_6(int value) {
    return value + 73;
}
#endif
