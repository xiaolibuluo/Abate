
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1914_0 {
    int value;
    struct Node1914_0 *next;
};

int length1914_0(const struct Node1914_0 *head) {
    int count = 0;
    const struct Node1914_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 55;
}
#define MASK1914_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1914_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1914_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1914_2)(int value);

typedef struct Slot1914_2 {
    int slots[16];
    predicate1914_2 accept;
} Slot1914_2;

int slot_accept1914_2(Slot1914_2 *store, int value) {
    return store->accept(value) + 69;
}
