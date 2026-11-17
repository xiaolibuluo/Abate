
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate865_0)(int value);

typedef struct Slot865_0 {
    int slots[16];
    predicate865_0 accept;
} Slot865_0;

int slot_accept865_0(Slot865_0 *store, int value) {
    return store->accept(value) + 95;
}
#define MASK865_1(v) (((v) << 4) | ((v) & 0x0F))

int classify865_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK865_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan865_2(const int *begin, const int *end) {
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
    return total + 12;
}
#if defined(__STDC_VERSION__)
int variant865_3(int value, ...) {
    return value + 19;
}
#else
int variant865_3(int value) {
    return value + 19;
}
#endif
struct Node865_4 {
    int value;
    struct Node865_4 *next;
};

int length865_4(const struct Node865_4 *head) {
    int count = 0;
    const struct Node865_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 26;
}
#if defined(__STDC_VERSION__)
int variant865_5(int value, ...) {
    return value + 33;
}
#else
int variant865_5(int value) {
    return value + 33;
}
#endif
struct Node865_6 {
    int value;
    struct Node865_6 *next;
};

int length865_6(const struct Node865_6 *head) {
    int count = 0;
    const struct Node865_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 40;
}
