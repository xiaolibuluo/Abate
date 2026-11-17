
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2547_0 {
    int value;
    struct Node2547_0 *next;
};

int length2547_0(const struct Node2547_0 *head) {
    int count = 0;
    const struct Node2547_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
typedef int (*predicate2547_1)(int value);

typedef struct Slot2547_1 {
    int slots[16];
    predicate2547_1 accept;
} Slot2547_1;

int slot_accept2547_1(Slot2547_1 *store, int value) {
    return store->accept(value) + 46;
}
#define MASK2547_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2547_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2547_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2547_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2547_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2547_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
