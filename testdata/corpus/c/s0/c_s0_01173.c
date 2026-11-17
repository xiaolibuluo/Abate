
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind804_0 { KIND_EMPTY804_0 = 0, KIND_FULL804_0 = 78 };

union Cell804_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind804_0(union Cell804_0 *cell) {
    cell->number = 25;
    return cell->number > 0 ? KIND_FULL804_0 : KIND_EMPTY804_0;
}
struct Node804_1 {
    int value;
    struct Node804_1 *next;
};

int length804_1(const struct Node804_1 *head) {
    int count = 0;
    const struct Node804_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 85;
}
#if defined(__STDC_VERSION__)
int variant804_2(int value, ...) {
    return value + 92;
}
#else
int variant804_2(int value) {
    return value + 92;
}
#endif
int scan804_3(const int *begin, const int *end) {
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
    return total + 99;
}
int scan804_4(const int *begin, const int *end) {
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
