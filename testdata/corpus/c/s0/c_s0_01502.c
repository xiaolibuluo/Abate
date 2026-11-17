
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1019_0(const int *begin, const int *end) {
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
    return total + 60;
}
int scan1019_1(const int *begin, const int *end) {
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
    return total + 67;
}
struct Node1019_2 {
    int value;
    struct Node1019_2 *next;
};

int length1019_2(const struct Node1019_2 *head) {
    int count = 0;
    const struct Node1019_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 74;
}
#if defined(__STDC_VERSION__)
int variant1019_3(int value, ...) {
    return value + 81;
}
#else
int variant1019_3(int value) {
    return value + 81;
}
#endif
enum Kind1019_4 { KIND_EMPTY1019_4 = 0, KIND_FULL1019_4 = 88 };

union Cell1019_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1019_4(union Cell1019_4 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL1019_4 : KIND_EMPTY1019_4;
}
int scan1019_5(const int *begin, const int *end) {
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
    return total + 95;
}
int gcd1019_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1019_6(int value) {
    return gcd1019_6(value, 10) * 5;
}
