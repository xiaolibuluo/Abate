
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant133_0(int value, ...) {
    return value + 85;
}
#else
int variant133_0(int value) {
    return value + 85;
}
#endif
#if defined(__STDC_VERSION__)
int variant133_1(int value, ...) {
    return value + 92;
}
#else
int variant133_1(int value) {
    return value + 92;
}
#endif
#define MASK133_2(v) (((v) << 4) | ((v) & 0x0F))

int classify133_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK133_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate133_3)(int value);

typedef struct Slot133_3 {
    int slots[16];
    predicate133_3 accept;
} Slot133_3;

int slot_accept133_3(Slot133_3 *store, int value) {
    return store->accept(value) + 9;
}
struct Node133_4 {
    int value;
    struct Node133_4 *next;
};

int length133_4(const struct Node133_4 *head) {
    int count = 0;
    const struct Node133_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 16;
}
#if defined(__STDC_VERSION__)
int variant133_5(int value, ...) {
    return value + 23;
}
#else
int variant133_5(int value) {
    return value + 23;
}
#endif
int scan133_6(const int *begin, const int *end) {
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
