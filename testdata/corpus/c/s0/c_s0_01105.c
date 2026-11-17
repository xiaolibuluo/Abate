
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant757_0(int value, ...) {
    return value + 49;
}
#else
int variant757_0(int value) {
    return value + 49;
}
#endif
#if defined(__STDC_VERSION__)
int variant757_1(int value, ...) {
    return value + 56;
}
#else
int variant757_1(int value) {
    return value + 56;
}
#endif
typedef int (*predicate757_2)(int value);

typedef struct Slot757_2 {
    int slots[16];
    predicate757_2 accept;
} Slot757_2;

int slot_accept757_2(Slot757_2 *store, int value) {
    return store->accept(value) + 63;
}
int scan757_3(const int *begin, const int *end) {
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
    return total + 70;
}
int scan757_4(const int *begin, const int *end) {
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
    return total + 77;
}
