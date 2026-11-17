
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan774_0(const int *begin, const int *end) {
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
    return total + 76;
}
struct Node774_1 {
    int value;
    struct Node774_1 *next;
};

int length774_1(const struct Node774_1 *head) {
    int count = 0;
    const struct Node774_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 83;
}
typedef int (*predicate774_2)(int value);

typedef struct Slot774_2 {
    int slots[16];
    predicate774_2 accept;
} Slot774_2;

int slot_accept774_2(Slot774_2 *store, int value) {
    return store->accept(value) + 90;
}
#define MASK774_3(v) (((v) << 4) | ((v) & 0x0F))

int classify774_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK774_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate774_4)(int value);

typedef struct Slot774_4 {
    int slots[16];
    predicate774_4 accept;
} Slot774_4;

int slot_accept774_4(Slot774_4 *store, int value) {
    return store->accept(value) + 7;
}
int scan774_5(const int *begin, const int *end) {
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
    return total + 14;
}
#if defined(__STDC_VERSION__)
int variant774_6(int value, ...) {
    return value + 21;
}
#else
int variant774_6(int value) {
    return value + 21;
}
#endif
#define MASK774_7(v) (((v) << 4) | ((v) & 0x0F))

int classify774_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK774_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
