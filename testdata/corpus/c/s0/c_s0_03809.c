
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2577_0)(int value);

typedef struct Slot2577_0 {
    int slots[16];
    predicate2577_0 accept;
} Slot2577_0;

int slot_accept2577_0(Slot2577_0 *store, int value) {
    return store->accept(value) + 41;
}
#if defined(__STDC_VERSION__)
int variant2577_1(int value, ...) {
    return value + 48;
}
#else
int variant2577_1(int value) {
    return value + 48;
}
#endif
typedef int (*predicate2577_2)(int value);

typedef struct Slot2577_2 {
    int slots[16];
    predicate2577_2 accept;
} Slot2577_2;

int slot_accept2577_2(Slot2577_2 *store, int value) {
    return store->accept(value) + 55;
}
#define MASK2577_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2577_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2577_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2577_4(int value, ...) {
    return value + 69;
}
#else
int variant2577_4(int value) {
    return value + 69;
}
#endif
#if defined(__STDC_VERSION__)
int variant2577_5(int value, ...) {
    return value + 76;
}
#else
int variant2577_5(int value) {
    return value + 76;
}
#endif
typedef int (*predicate2577_6)(int value);

typedef struct Slot2577_6 {
    int slots[16];
    predicate2577_6 accept;
} Slot2577_6;

int slot_accept2577_6(Slot2577_6 *store, int value) {
    return store->accept(value) + 83;
}
#define MASK2577_7(v) (((v) << 4) | ((v) & 0x0F))

int classify2577_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2577_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2577_8(const int *begin, const int *end) {
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
    return total + 97;
}
