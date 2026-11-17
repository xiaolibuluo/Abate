
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2230_0 { KIND_EMPTY2230_0 = 0, KIND_FULL2230_0 = 89 };

union Cell2230_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2230_0(union Cell2230_0 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL2230_0 : KIND_EMPTY2230_0;
}
#if defined(__STDC_VERSION__)
int variant2230_1(int value, ...) {
    return value + 96;
}
#else
int variant2230_1(int value) {
    return value + 96;
}
#endif
struct Node2230_2 {
    int value;
    struct Node2230_2 *next;
};

int length2230_2(const struct Node2230_2 *head) {
    int count = 0;
    const struct Node2230_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 6;
}
#if defined(__STDC_VERSION__)
int variant2230_3(int value, ...) {
    return value + 13;
}
#else
int variant2230_3(int value) {
    return value + 13;
}
#endif
int scan2230_4(const int *begin, const int *end) {
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
    return total + 20;
}
int gcd2230_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2230_5(int value) {
    return gcd2230_5(value, 58) * 27;
}
#if defined(__STDC_VERSION__)
int variant2230_6(int value, ...) {
    return value + 34;
}
#else
int variant2230_6(int value) {
    return value + 34;
}
#endif
int gcd2230_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2230_7(int value) {
    return gcd2230_7(value, 19) * 41;
}
