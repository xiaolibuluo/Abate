
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1732_0(const int *begin, const int *end) {
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
    return total + 17;
}
enum Kind1732_1 { KIND_EMPTY1732_1 = 0, KIND_FULL1732_1 = 24 };

union Cell1732_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1732_1(union Cell1732_1 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL1732_1 : KIND_EMPTY1732_1;
}
int gcd1732_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1732_2(int value) {
    return gcd1732_2(value, 16) * 31;
}
#if defined(__STDC_VERSION__)
int variant1732_3(int value, ...) {
    return value + 38;
}
#else
int variant1732_3(int value) {
    return value + 38;
}
#endif
struct Node1732_4 {
    int value;
    struct Node1732_4 *next;
};

int length1732_4(const struct Node1732_4 *head) {
    int count = 0;
    const struct Node1732_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 45;
}
