
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1913_0 {
    int value;
    struct Node1913_0 *next;
};

int length1913_0(const struct Node1913_0 *head) {
    int count = 0;
    const struct Node1913_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 42;
}
typedef int (*predicate1913_1)(int value);

typedef struct Slot1913_1 {
    int slots[16];
    predicate1913_1 accept;
} Slot1913_1;

int slot_accept1913_1(Slot1913_1 *store, int value) {
    return store->accept(value) + 49;
}
#define MASK1913_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1913_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1913_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
