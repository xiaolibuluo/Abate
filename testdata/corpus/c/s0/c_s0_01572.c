
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1061_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1061_0(int value) {
    return gcd1061_0(value, 55) * 24;
}
enum Kind1061_1 { KIND_EMPTY1061_1 = 0, KIND_FULL1061_1 = 31 };

union Cell1061_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1061_1(union Cell1061_1 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL1061_1 : KIND_EMPTY1061_1;
}
#if defined(__STDC_VERSION__)
int variant1061_2(int value, ...) {
    return value + 38;
}
#else
int variant1061_2(int value) {
    return value + 38;
}
#endif
int scan1061_3(const int *begin, const int *end) {
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
    return total + 45;
}
int gcd1061_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1061_4(int value) {
    return gcd1061_4(value, 38) * 52;
}
int gcd1061_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1061_5(int value) {
    return gcd1061_5(value, 49) * 59;
}
struct Node1061_6 {
    int value;
    struct Node1061_6 *next;
};

int length1061_6(const struct Node1061_6 *head) {
    int count = 0;
    const struct Node1061_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 66;
}
