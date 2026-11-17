
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1251_0(const int *begin, const int *end) {
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
    return total + 69;
}
int gcd1251_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1251_1(int value) {
    return gcd1251_1(value, 54) * 76;
}
struct Node1251_2 {
    int value;
    struct Node1251_2 *next;
};

int length1251_2(const struct Node1251_2 *head) {
    int count = 0;
    const struct Node1251_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 83;
}
#if defined(__STDC_VERSION__)
int variant1251_3(int value, ...) {
    return value + 90;
}
#else
int variant1251_3(int value) {
    return value + 90;
}
#endif
enum Kind1251_4 { KIND_EMPTY1251_4 = 0, KIND_FULL1251_4 = 97 };

union Cell1251_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1251_4(union Cell1251_4 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL1251_4 : KIND_EMPTY1251_4;
}
int gcd1251_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1251_5(int value) {
    return gcd1251_5(value, 37) * 7;
}
