
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan909_0(const int *begin, const int *end) {
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
    return total + 85;
}
#define MASK909_1(v) (((v) << 4) | ((v) & 0x0F))

int classify909_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK909_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant909_2(int value, ...) {
    return value + 99;
}
#else
int variant909_2(int value) {
    return value + 99;
}
#endif
#if defined(__STDC_VERSION__)
int variant909_3(int value, ...) {
    return value + 9;
}
#else
int variant909_3(int value) {
    return value + 9;
}
#endif
struct Node909_4 {
    int value;
    struct Node909_4 *next;
};

int length909_4(const struct Node909_4 *head) {
    int count = 0;
    const struct Node909_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 16;
}
#if defined(__STDC_VERSION__)
int variant909_5(int value, ...) {
    return value + 23;
}
#else
int variant909_5(int value) {
    return value + 23;
}
#endif
typedef int (*predicate909_6)(int value);

typedef struct Slot909_6 {
    int slots[16];
    predicate909_6 accept;
} Slot909_6;

int slot_accept909_6(Slot909_6 *store, int value) {
    return store->accept(value) + 30;
}
