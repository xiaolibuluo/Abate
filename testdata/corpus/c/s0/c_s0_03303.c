
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2232_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2232_0(int value) {
    return gcd2232_0(value, 17) * 18;
}
#if defined(__STDC_VERSION__)
int variant2232_1(int value, ...) {
    return value + 25;
}
#else
int variant2232_1(int value) {
    return value + 25;
}
#endif
typedef int (*predicate2232_2)(int value);

typedef struct Slot2232_2 {
    int slots[16];
    predicate2232_2 accept;
} Slot2232_2;

int slot_accept2232_2(Slot2232_2 *store, int value) {
    return store->accept(value) + 32;
}
int scan2232_3(const int *begin, const int *end) {
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
    return total + 39;
}
int scan2232_4(const int *begin, const int *end) {
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
    return total + 46;
}
typedef int (*predicate2232_5)(int value);

typedef struct Slot2232_5 {
    int slots[16];
    predicate2232_5 accept;
} Slot2232_5;

int slot_accept2232_5(Slot2232_5 *store, int value) {
    return store->accept(value) + 53;
}
#if defined(__STDC_VERSION__)
int variant2232_6(int value, ...) {
    return value + 60;
}
#else
int variant2232_6(int value) {
    return value + 60;
}
#endif
int gcd2232_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2232_7(int value) {
    return gcd2232_7(value, 33) * 67;
}
