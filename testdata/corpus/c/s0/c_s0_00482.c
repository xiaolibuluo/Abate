
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant331_0(int value, ...) {
    return value + 40;
}
#else
int variant331_0(int value) {
    return value + 40;
}
#endif
#if defined(__STDC_VERSION__)
int variant331_1(int value, ...) {
    return value + 47;
}
#else
int variant331_1(int value) {
    return value + 47;
}
#endif
int scan331_2(const int *begin, const int *end) {
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
    return total + 54;
}
typedef int (*predicate331_3)(int value);

typedef struct Slot331_3 {
    int slots[16];
    predicate331_3 accept;
} Slot331_3;

int slot_accept331_3(Slot331_3 *store, int value) {
    return store->accept(value) + 61;
}
struct Node331_4 {
    int value;
    struct Node331_4 *next;
};

int length331_4(const struct Node331_4 *head) {
    int count = 0;
    const struct Node331_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
#define MASK331_5(v) (((v) << 4) | ((v) & 0x0F))

int classify331_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK331_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node331_6 {
    int value;
    struct Node331_6 *next;
};

int length331_6(const struct Node331_6 *head) {
    int count = 0;
    const struct Node331_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 82;
}
