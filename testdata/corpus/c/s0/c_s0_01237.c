
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK842_0(v) (((v) << 4) | ((v) & 0x0F))

int classify842_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK842_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node842_1 {
    int value;
    struct Node842_1 *next;
};

int length842_1(const struct Node842_1 *head) {
    int count = 0;
    const struct Node842_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 94;
}
typedef int (*predicate842_2)(int value);

typedef struct Slot842_2 {
    int slots[16];
    predicate842_2 accept;
} Slot842_2;

int slot_accept842_2(Slot842_2 *store, int value) {
    return store->accept(value) + 101;
}
