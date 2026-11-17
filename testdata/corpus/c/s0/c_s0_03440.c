
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2327_0)(int value);

typedef struct Slot2327_0 {
    int slots[16];
    predicate2327_0 accept;
} Slot2327_0;

int slot_accept2327_0(Slot2327_0 *store, int value) {
    return store->accept(value) + 89;
}
typedef int (*predicate2327_1)(int value);

typedef struct Slot2327_1 {
    int slots[16];
    predicate2327_1 accept;
} Slot2327_1;

int slot_accept2327_1(Slot2327_1 *store, int value) {
    return store->accept(value) + 96;
}
typedef int (*predicate2327_2)(int value);

typedef struct Slot2327_2 {
    int slots[16];
    predicate2327_2 accept;
} Slot2327_2;

int slot_accept2327_2(Slot2327_2 *store, int value) {
    return store->accept(value) + 6;
}
int gcd2327_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2327_3(int value) {
    return gcd2327_3(value, 44) * 13;
}
#if defined(__STDC_VERSION__)
int variant2327_4(int value, ...) {
    return value + 20;
}
#else
int variant2327_4(int value) {
    return value + 20;
}
#endif
int gcd2327_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2327_5(int value) {
    return gcd2327_5(value, 66) * 27;
}
int scan2327_6(const int *begin, const int *end) {
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
    return total + 34;
}
