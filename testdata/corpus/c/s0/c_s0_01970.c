
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1330_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1330_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1330_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1330_1 {
    int value;
    struct Node1330_1 *next;
};

int length1330_1(const struct Node1330_1 *head) {
    int count = 0;
    const struct Node1330_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 36;
}
typedef int (*predicate1330_2)(int value);

typedef struct Slot1330_2 {
    int slots[16];
    predicate1330_2 accept;
} Slot1330_2;

int slot_accept1330_2(Slot1330_2 *store, int value) {
    return store->accept(value) + 43;
}
