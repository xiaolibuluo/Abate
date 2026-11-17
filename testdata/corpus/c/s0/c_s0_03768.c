
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2548_0)(int value);

typedef struct Slot2548_0 {
    int slots[16];
    predicate2548_0 accept;
} Slot2548_0;

int slot_accept2548_0(Slot2548_0 *store, int value) {
    return store->accept(value) + 52;
}
#define MASK2548_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2548_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2548_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2548_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2548_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2548_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2548_3 {
    int value;
    struct Node2548_3 *next;
};

int length2548_3(const struct Node2548_3 *head) {
    int count = 0;
    const struct Node2548_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 73;
}
