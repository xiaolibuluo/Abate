
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2646_0(int value, ...) {
    return value + 65;
}
#else
int variant2646_0(int value) {
    return value + 65;
}
#endif
enum Kind2646_1 { KIND_EMPTY2646_1 = 0, KIND_FULL2646_1 = 72 };

union Cell2646_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2646_1(union Cell2646_1 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL2646_1 : KIND_EMPTY2646_1;
}
int scan2646_2(const int *begin, const int *end) {
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
    return total + 79;
}
struct Node2646_3 {
    int value;
    struct Node2646_3 *next;
};

int length2646_3(const struct Node2646_3 *head) {
    int count = 0;
    const struct Node2646_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 86;
}
