
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate304_0)(int value);

typedef struct Slot304_0 {
    int slots[16];
    predicate304_0 accept;
} Slot304_0;

int slot_accept304_0(Slot304_0 *store, int value) {
    return store->accept(value) + 77;
}
#define MASK304_1(v) (((v) << 4) | ((v) & 0x0F))

int classify304_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK304_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node304_2 {
    int value;
    struct Node304_2 *next;
};

int length304_2(const struct Node304_2 *head) {
    int count = 0;
    const struct Node304_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 91;
}
