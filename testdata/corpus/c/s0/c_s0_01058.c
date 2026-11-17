
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan724_0(const int *begin, const int *end) {
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
    return total + 8;
}
struct Node724_1 {
    int value;
    struct Node724_1 *next;
};

int length724_1(const struct Node724_1 *head) {
    int count = 0;
    const struct Node724_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 15;
}
#if defined(__STDC_VERSION__)
int variant724_2(int value, ...) {
    return value + 22;
}
#else
int variant724_2(int value) {
    return value + 22;
}
#endif
int gcd724_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale724_3(int value) {
    return gcd724_3(value, 47) * 29;
}
int scan724_4(const int *begin, const int *end) {
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
    return total + 36;
}
enum Kind724_5 { KIND_EMPTY724_5 = 0, KIND_FULL724_5 = 43 };

union Cell724_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind724_5(union Cell724_5 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL724_5 : KIND_EMPTY724_5;
}
int scan724_6(const int *begin, const int *end) {
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
    return total + 50;
}
