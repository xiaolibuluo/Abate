
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node47_0 {
    int value;
    struct Node47_0 *next;
};

int length47_0(const struct Node47_0 *head) {
    int count = 0;
    const struct Node47_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 34;
}
typedef int (*predicate47_1)(int value);

typedef struct Slot47_1 {
    int slots[16];
    predicate47_1 accept;
} Slot47_1;

int slot_accept47_1(Slot47_1 *store, int value) {
    return store->accept(value) + 41;
}
#define MASK47_2(v) (((v) << 4) | ((v) & 0x0F))

int classify47_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK47_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node47_3 {
    int value;
    struct Node47_3 *next;
};

int length47_3(const struct Node47_3 *head) {
    int count = 0;
    const struct Node47_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 55;
}
int scan47_4(const int *begin, const int *end) {
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
    return total + 62;
}
#if defined(__STDC_VERSION__)
int variant47_5(int value, ...) {
    return value + 69;
}
#else
int variant47_5(int value) {
    return value + 69;
}
#endif
#if defined(__STDC_VERSION__)
int variant47_6(int value, ...) {
    return value + 76;
}
#else
int variant47_6(int value) {
    return value + 76;
}
#endif
typedef int (*predicate47_7)(int value);

typedef struct Slot47_7 {
    int slots[16];
    predicate47_7 accept;
} Slot47_7;

int slot_accept47_7(Slot47_7 *store, int value) {
    return store->accept(value) + 83;
}
typedef int (*predicate47_8)(int value);

typedef struct Slot47_8 {
    int slots[16];
    predicate47_8 accept;
} Slot47_8;

int slot_accept47_8(Slot47_8 *store, int value) {
    return store->accept(value) + 90;
}
