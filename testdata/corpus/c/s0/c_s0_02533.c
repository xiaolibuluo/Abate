
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1721_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1721_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1721_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1721_1 {
    int value;
    struct Node1721_1 *next;
};

int length1721_1(const struct Node1721_1 *head) {
    int count = 0;
    const struct Node1721_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 75;
}
typedef int (*predicate1721_2)(int value);

typedef struct Slot1721_2 {
    int slots[16];
    predicate1721_2 accept;
} Slot1721_2;

int slot_accept1721_2(Slot1721_2 *store, int value) {
    return store->accept(value) + 82;
}
struct Node1721_3 {
    int value;
    struct Node1721_3 *next;
};

int length1721_3(const struct Node1721_3 *head) {
    int count = 0;
    const struct Node1721_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 89;
}
