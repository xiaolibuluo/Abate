
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2624_0)(int value);

typedef struct Slot2624_0 {
    int slots[16];
    predicate2624_0 accept;
} Slot2624_0;

int slot_accept2624_0(Slot2624_0 *store, int value) {
    return store->accept(value) + 70;
}
#if defined(__STDC_VERSION__)
int variant2624_1(int value, ...) {
    return value + 77;
}
#else
int variant2624_1(int value) {
    return value + 77;
}
#endif
typedef int (*predicate2624_2)(int value);

typedef struct Slot2624_2 {
    int slots[16];
    predicate2624_2 accept;
} Slot2624_2;

int slot_accept2624_2(Slot2624_2 *store, int value) {
    return store->accept(value) + 84;
}
#define MASK2624_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2624_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2624_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2624_4 {
    int value;
    struct Node2624_4 *next;
};

int length2624_4(const struct Node2624_4 *head) {
    int count = 0;
    const struct Node2624_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 98;
}
typedef int (*predicate2624_5)(int value);

typedef struct Slot2624_5 {
    int slots[16];
    predicate2624_5 accept;
} Slot2624_5;

int slot_accept2624_5(Slot2624_5 *store, int value) {
    return store->accept(value) + 8;
}
int scan2624_6(const int *begin, const int *end) {
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
    return total + 15;
}
typedef int (*predicate2624_7)(int value);

typedef struct Slot2624_7 {
    int slots[16];
    predicate2624_7 accept;
} Slot2624_7;

int slot_accept2624_7(Slot2624_7 *store, int value) {
    return store->accept(value) + 22;
}
