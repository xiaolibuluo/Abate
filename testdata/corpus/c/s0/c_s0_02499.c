
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1699_0 { KIND_EMPTY1699_0 = 0, KIND_FULL1699_0 = 73 };

union Cell1699_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1699_0(union Cell1699_0 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL1699_0 : KIND_EMPTY1699_0;
}
int gcd1699_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1699_1(int value) {
    return gcd1699_1(value, 18) * 80;
}
int scan1699_2(const int *begin, const int *end) {
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
    return total + 87;
}
struct Node1699_3 {
    int value;
    struct Node1699_3 *next;
};

int length1699_3(const struct Node1699_3 *head) {
    int count = 0;
    const struct Node1699_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 94;
}
enum Kind1699_4 { KIND_EMPTY1699_4 = 0, KIND_FULL1699_4 = 101 };

union Cell1699_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1699_4(union Cell1699_4 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL1699_4 : KIND_EMPTY1699_4;
}
#if defined(__STDC_VERSION__)
int variant1699_5(int value, ...) {
    return value + 11;
}
#else
int variant1699_5(int value) {
    return value + 11;
}
#endif
#if defined(__STDC_VERSION__)
int variant1699_6(int value, ...) {
    return value + 18;
}
#else
int variant1699_6(int value) {
    return value + 18;
}
#endif
int gcd1699_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1699_7(int value) {
    return gcd1699_7(value, 23) * 25;
}
