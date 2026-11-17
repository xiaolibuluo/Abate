
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1928_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1928_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1928_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1928_1)(int value);

typedef struct Slot1928_1 {
    int slots[16];
    predicate1928_1 accept;
} Slot1928_1;

int slot_accept1928_1(Slot1928_1 *store, int value) {
    return store->accept(value) + 50;
}
struct Node1928_2 {
    int value;
    struct Node1928_2 *next;
};

int length1928_2(const struct Node1928_2 *head) {
    int count = 0;
    const struct Node1928_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 57;
}
typedef int (*predicate1928_3)(int value);

typedef struct Slot1928_3 {
    int slots[16];
    predicate1928_3 accept;
} Slot1928_3;

int slot_accept1928_3(Slot1928_3 *store, int value) {
    return store->accept(value) + 64;
}
#define MASK1928_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1928_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1928_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1928_5(int value, ...) {
    return value + 78;
}
#else
int variant1928_5(int value) {
    return value + 78;
}
#endif
