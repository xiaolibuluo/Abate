
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1715_0 {
    int value;
    struct Node1715_0 *next;
};

int length1715_0(const struct Node1715_0 *head) {
    int count = 0;
    const struct Node1715_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 87;
}
typedef int (*predicate1715_1)(int value);

typedef struct Slot1715_1 {
    int slots[16];
    predicate1715_1 accept;
} Slot1715_1;

int slot_accept1715_1(Slot1715_1 *store, int value) {
    return store->accept(value) + 94;
}
#define MASK1715_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1715_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1715_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
