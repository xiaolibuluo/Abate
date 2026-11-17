
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1724_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1724_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1724_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1724_1 {
    int value;
    struct Node1724_1 *next;
};

int length1724_1(const struct Node1724_1 *head) {
    int count = 0;
    const struct Node1724_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 17;
}
typedef int (*predicate1724_2)(int value);

typedef struct Slot1724_2 {
    int slots[16];
    predicate1724_2 accept;
} Slot1724_2;

int slot_accept1724_2(Slot1724_2 *store, int value) {
    return store->accept(value) + 24;
}
#if defined(__STDC_VERSION__)
int variant1724_3(int value, ...) {
    return value + 31;
}
#else
int variant1724_3(int value) {
    return value + 31;
}
#endif
