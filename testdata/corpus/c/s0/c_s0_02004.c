
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1352_0(const int *begin, const int *end) {
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
    return total + 24;
}
struct Node1352_1 {
    int value;
    struct Node1352_1 *next;
};

int length1352_1(const struct Node1352_1 *head) {
    int count = 0;
    const struct Node1352_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 31;
}
int scan1352_2(const int *begin, const int *end) {
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
    return total + 38;
}
#if defined(__STDC_VERSION__)
int variant1352_3(int value, ...) {
    return value + 45;
}
#else
int variant1352_3(int value) {
    return value + 45;
}
#endif
int scan1352_4(const int *begin, const int *end) {
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
enum Kind1352_5 { KIND_EMPTY1352_5 = 0, KIND_FULL1352_5 = 59 };

union Cell1352_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1352_5(union Cell1352_5 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL1352_5 : KIND_EMPTY1352_5;
}
