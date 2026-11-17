
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node707_0 {
    int value;
    struct Node707_0 *next;
};

int length707_0(const struct Node707_0 *head) {
    int count = 0;
    const struct Node707_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 78;
}
#if defined(__STDC_VERSION__)
int variant707_1(int value, ...) {
    return value + 85;
}
#else
int variant707_1(int value) {
    return value + 85;
}
#endif
#if defined(__STDC_VERSION__)
int variant707_2(int value, ...) {
    return value + 92;
}
#else
int variant707_2(int value) {
    return value + 92;
}
#endif
enum Kind707_3 { KIND_EMPTY707_3 = 0, KIND_FULL707_3 = 99 };

union Cell707_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind707_3(union Cell707_3 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL707_3 : KIND_EMPTY707_3;
}
int scan707_4(const int *begin, const int *end) {
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
    return total + 9;
}
