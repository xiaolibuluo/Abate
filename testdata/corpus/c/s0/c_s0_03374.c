
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2275_0(const int *begin, const int *end) {
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
    return total + 92;
}
struct Node2275_1 {
    int value;
    struct Node2275_1 *next;
};

int length2275_1(const struct Node2275_1 *head) {
    int count = 0;
    const struct Node2275_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 99;
}
struct Node2275_2 {
    int value;
    struct Node2275_2 *next;
};

int length2275_2(const struct Node2275_2 *head) {
    int count = 0;
    const struct Node2275_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 9;
}
enum Kind2275_3 { KIND_EMPTY2275_3 = 0, KIND_FULL2275_3 = 16 };

union Cell2275_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2275_3(union Cell2275_3 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL2275_3 : KIND_EMPTY2275_3;
}
#if defined(__STDC_VERSION__)
int variant2275_4(int value, ...) {
    return value + 23;
}
#else
int variant2275_4(int value) {
    return value + 23;
}
#endif
int scan2275_5(const int *begin, const int *end) {
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
    return total + 30;
}
