
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2115_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2115_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2115_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2115_1)(int value);

typedef struct Slot2115_1 {
    int slots[16];
    predicate2115_1 accept;
} Slot2115_1;

int slot_accept2115_1(Slot2115_1 *store, int value) {
    return store->accept(value) + 56;
}
#define MASK2115_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2115_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2115_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2115_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2115_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2115_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2115_4 {
    int value;
    struct Node2115_4 *next;
};

int length2115_4(const struct Node2115_4 *head) {
    int count = 0;
    const struct Node2115_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 77;
}
