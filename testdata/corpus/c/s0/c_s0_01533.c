
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1040_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1040_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1040_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1040_1(int value, ...) {
    return value + 49;
}
#else
int variant1040_1(int value) {
    return value + 49;
}
#endif
typedef int (*predicate1040_2)(int value);

typedef struct Slot1040_2 {
    int slots[16];
    predicate1040_2 accept;
} Slot1040_2;

int slot_accept1040_2(Slot1040_2 *store, int value) {
    return store->accept(value) + 56;
}
struct Node1040_3 {
    int value;
    struct Node1040_3 *next;
};

int length1040_3(const struct Node1040_3 *head) {
    int count = 0;
    const struct Node1040_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
