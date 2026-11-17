
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1501_0)(int value);

typedef struct Slot1501_0 {
    int slots[16];
    predicate1501_0 accept;
} Slot1501_0;

int slot_accept1501_0(Slot1501_0 *store, int value) {
    return store->accept(value) + 21;
}
struct Node1501_1 {
    int value;
    struct Node1501_1 *next;
};

int length1501_1(const struct Node1501_1 *head) {
    int count = 0;
    const struct Node1501_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 28;
}
typedef int (*predicate1501_2)(int value);

typedef struct Slot1501_2 {
    int slots[16];
    predicate1501_2 accept;
} Slot1501_2;

int slot_accept1501_2(Slot1501_2 *store, int value) {
    return store->accept(value) + 35;
}
#define MASK1501_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1501_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1501_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1501_4 {
    int value;
    struct Node1501_4 *next;
};

int length1501_4(const struct Node1501_4 *head) {
    int count = 0;
    const struct Node1501_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 49;
}
typedef int (*predicate1501_5)(int value);

typedef struct Slot1501_5 {
    int slots[16];
    predicate1501_5 accept;
} Slot1501_5;

int slot_accept1501_5(Slot1501_5 *store, int value) {
    return store->accept(value) + 56;
}
#if defined(__STDC_VERSION__)
int variant1501_6(int value, ...) {
    return value + 63;
}
#else
int variant1501_6(int value) {
    return value + 63;
}
#endif
#if defined(__STDC_VERSION__)
int variant1501_7(int value, ...) {
    return value + 70;
}
#else
int variant1501_7(int value) {
    return value + 70;
}
#endif
