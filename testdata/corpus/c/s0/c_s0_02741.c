
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1864_0 {
    int value;
    struct Node1864_0 *next;
};

int length1864_0(const struct Node1864_0 *head) {
    int count = 0;
    const struct Node1864_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 84;
}
#define MASK1864_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1864_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1864_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1864_2)(int value);

typedef struct Slot1864_2 {
    int slots[16];
    predicate1864_2 accept;
} Slot1864_2;

int slot_accept1864_2(Slot1864_2 *store, int value) {
    return store->accept(value) + 98;
}
