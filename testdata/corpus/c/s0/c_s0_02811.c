
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1912_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1912_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1912_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1912_1)(int value);

typedef struct Slot1912_1 {
    int slots[16];
    predicate1912_1 accept;
} Slot1912_1;

int slot_accept1912_1(Slot1912_1 *store, int value) {
    return store->accept(value) + 36;
}
struct Node1912_2 {
    int value;
    struct Node1912_2 *next;
};

int length1912_2(const struct Node1912_2 *head) {
    int count = 0;
    const struct Node1912_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 43;
}
