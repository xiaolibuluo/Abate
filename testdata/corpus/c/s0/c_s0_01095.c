
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan748_0(const int *begin, const int *end) {
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
    return total + 29;
}
#if defined(__STDC_VERSION__)
int variant748_1(int value, ...) {
    return value + 36;
}
#else
int variant748_1(int value) {
    return value + 36;
}
#endif
struct Node748_2 {
    int value;
    struct Node748_2 *next;
};

int length748_2(const struct Node748_2 *head) {
    int count = 0;
    const struct Node748_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 43;
}
enum Kind748_3 { KIND_EMPTY748_3 = 0, KIND_FULL748_3 = 50 };

union Cell748_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind748_3(union Cell748_3 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL748_3 : KIND_EMPTY748_3;
}
