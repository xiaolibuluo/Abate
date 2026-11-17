
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1799_0(int value, ...) {
    return value + 15;
}
#else
int variant1799_0(int value) {
    return value + 15;
}
#endif
struct Node1799_1 {
    int value;
    struct Node1799_1 *next;
};

int length1799_1(const struct Node1799_1 *head) {
    int count = 0;
    const struct Node1799_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 22;
}
enum Kind1799_2 { KIND_EMPTY1799_2 = 0, KIND_FULL1799_2 = 29 };

union Cell1799_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1799_2(union Cell1799_2 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL1799_2 : KIND_EMPTY1799_2;
}
int scan1799_3(const int *begin, const int *end) {
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
int gcd1799_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1799_4(int value) {
    return gcd1799_4(value, 19) * 43;
}
#if defined(__STDC_VERSION__)
int variant1799_5(int value, ...) {
    return value + 50;
}
#else
int variant1799_5(int value) {
    return value + 50;
}
#endif
int scan1799_6(const int *begin, const int *end) {
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
    return total + 57;
}
enum Kind1799_7 { KIND_EMPTY1799_7 = 0, KIND_FULL1799_7 = 64 };

union Cell1799_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1799_7(union Cell1799_7 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL1799_7 : KIND_EMPTY1799_7;
}
