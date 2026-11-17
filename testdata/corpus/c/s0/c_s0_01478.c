
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1005_0)(int value);

typedef struct Slot1005_0 {
    int slots[16];
    predicate1005_0 accept;
} Slot1005_0;

int slot_accept1005_0(Slot1005_0 *store, int value) {
    return store->accept(value) + 72;
}
struct Node1005_1 {
    int value;
    struct Node1005_1 *next;
};

int length1005_1(const struct Node1005_1 *head) {
    int count = 0;
    const struct Node1005_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 79;
}
#if defined(__STDC_VERSION__)
int variant1005_2(int value, ...) {
    return value + 86;
}
#else
int variant1005_2(int value) {
    return value + 86;
}
#endif
struct Node1005_3 {
    int value;
    struct Node1005_3 *next;
};

int length1005_3(const struct Node1005_3 *head) {
    int count = 0;
    const struct Node1005_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 93;
}
typedef int (*predicate1005_4)(int value);

typedef struct Slot1005_4 {
    int slots[16];
    predicate1005_4 accept;
} Slot1005_4;

int slot_accept1005_4(Slot1005_4 *store, int value) {
    return store->accept(value) + 100;
}
#define MASK1005_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1005_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1005_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
