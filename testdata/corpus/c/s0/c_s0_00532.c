
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan365_0(const int *begin, const int *end) {
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
    return total + 94;
}
typedef int (*predicate365_1)(int value);

typedef struct Slot365_1 {
    int slots[16];
    predicate365_1 accept;
} Slot365_1;

int slot_accept365_1(Slot365_1 *store, int value) {
    return store->accept(value) + 101;
}
typedef int (*predicate365_2)(int value);

typedef struct Slot365_2 {
    int slots[16];
    predicate365_2 accept;
} Slot365_2;

int slot_accept365_2(Slot365_2 *store, int value) {
    return store->accept(value) + 11;
}
typedef int (*predicate365_3)(int value);

typedef struct Slot365_3 {
    int slots[16];
    predicate365_3 accept;
} Slot365_3;

int slot_accept365_3(Slot365_3 *store, int value) {
    return store->accept(value) + 18;
}
#if defined(__STDC_VERSION__)
int variant365_4(int value, ...) {
    return value + 25;
}
#else
int variant365_4(int value) {
    return value + 25;
}
#endif
