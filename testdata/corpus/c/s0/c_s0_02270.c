
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1532_0(int value, ...) {
    return value + 36;
}
#else
int variant1532_0(int value) {
    return value + 36;
}
#endif
#define MASK1532_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1532_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1532_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1532_2 {
    int value;
    struct Node1532_2 *next;
};

int length1532_2(const struct Node1532_2 *head) {
    int count = 0;
    const struct Node1532_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 50;
}
typedef int (*predicate1532_3)(int value);

typedef struct Slot1532_3 {
    int slots[16];
    predicate1532_3 accept;
} Slot1532_3;

int slot_accept1532_3(Slot1532_3 *store, int value) {
    return store->accept(value) + 57;
}
typedef int (*predicate1532_4)(int value);

typedef struct Slot1532_4 {
    int slots[16];
    predicate1532_4 accept;
} Slot1532_4;

int slot_accept1532_4(Slot1532_4 *store, int value) {
    return store->accept(value) + 64;
}
