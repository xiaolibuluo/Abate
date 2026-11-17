
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1771_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1771_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1771_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1771_1 {
    int value;
    struct Node1771_1 *next;
};

int length1771_1(const struct Node1771_1 *head) {
    int count = 0;
    const struct Node1771_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}
struct Node1771_2 {
    int value;
    struct Node1771_2 *next;
};

int length1771_2(const struct Node1771_2 *head) {
    int count = 0;
    const struct Node1771_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 53;
}
typedef int (*predicate1771_3)(int value);

typedef struct Slot1771_3 {
    int slots[16];
    predicate1771_3 accept;
} Slot1771_3;

int slot_accept1771_3(Slot1771_3 *store, int value) {
    return store->accept(value) + 60;
}
