
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1537_0 {
    int value;
    struct Node1537_0 *next;
};

int length1537_0(const struct Node1537_0 *head) {
    int count = 0;
    const struct Node1537_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 101;
}
int scan1537_1(const int *begin, const int *end) {
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
    return total + 11;
}
#if defined(__STDC_VERSION__)
int variant1537_2(int value, ...) {
    return value + 18;
}
#else
int variant1537_2(int value) {
    return value + 18;
}
#endif
int scan1537_3(const int *begin, const int *end) {
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
    return total + 25;
}
enum Kind1537_4 { KIND_EMPTY1537_4 = 0, KIND_FULL1537_4 = 32 };

union Cell1537_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1537_4(union Cell1537_4 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL1537_4 : KIND_EMPTY1537_4;
}
