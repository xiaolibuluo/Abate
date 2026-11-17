
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1779_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1779_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1779_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1779_1 {
    int value;
    struct Node1779_1 *next;
};

int length1779_1(const struct Node1779_1 *head) {
    int count = 0;
    const struct Node1779_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 53;
}
#define MASK1779_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1779_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1779_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1779_3)(int value);

typedef struct Slot1779_3 {
    int slots[16];
    predicate1779_3 accept;
} Slot1779_3;

int slot_accept1779_3(Slot1779_3 *store, int value) {
    return store->accept(value) + 67;
}
#if defined(__STDC_VERSION__)
int variant1779_4(int value, ...) {
    return value + 74;
}
#else
int variant1779_4(int value) {
    return value + 74;
}
#endif
typedef int (*predicate1779_5)(int value);

typedef struct Slot1779_5 {
    int slots[16];
    predicate1779_5 accept;
} Slot1779_5;

int slot_accept1779_5(Slot1779_5 *store, int value) {
    return store->accept(value) + 81;
}
