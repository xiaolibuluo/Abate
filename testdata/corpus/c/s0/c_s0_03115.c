
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2109_0)(int value);

typedef struct Slot2109_0 {
    int slots[16];
    predicate2109_0 accept;
} Slot2109_0;

int slot_accept2109_0(Slot2109_0 *store, int value) {
    return store->accept(value) + 68;
}
struct Node2109_1 {
    int value;
    struct Node2109_1 *next;
};

int length2109_1(const struct Node2109_1 *head) {
    int count = 0;
    const struct Node2109_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 75;
}
#define MASK2109_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2109_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2109_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2109_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2109_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2109_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
