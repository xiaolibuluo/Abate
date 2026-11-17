
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1959_0)(int value);

typedef struct Slot1959_0 {
    int slots[16];
    predicate1959_0 accept;
} Slot1959_0;

int slot_accept1959_0(Slot1959_0 *store, int value) {
    return store->accept(value) + 58;
}
struct Node1959_1 {
    int value;
    struct Node1959_1 *next;
};

int length1959_1(const struct Node1959_1 *head) {
    int count = 0;
    const struct Node1959_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 65;
}
#define MASK1959_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1959_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1959_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
