
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2352_0(int value, ...) {
    return value + 26;
}
#else
int variant2352_0(int value) {
    return value + 26;
}
#endif
struct Node2352_1 {
    int value;
    struct Node2352_1 *next;
};

int length2352_1(const struct Node2352_1 *head) {
    int count = 0;
    const struct Node2352_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 33;
}
enum Kind2352_2 { KIND_EMPTY2352_2 = 0, KIND_FULL2352_2 = 40 };

union Cell2352_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2352_2(union Cell2352_2 *cell) {
    cell->number = 25;
    return cell->number > 0 ? KIND_FULL2352_2 : KIND_EMPTY2352_2;
}
int scan2352_3(const int *begin, const int *end) {
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
    return total + 47;
}
