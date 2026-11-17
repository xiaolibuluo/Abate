
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2400_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2400_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2400_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2400_1)(int value);

typedef struct Slot2400_1 {
    int slots[16];
    predicate2400_1 accept;
} Slot2400_1;

int slot_accept2400_1(Slot2400_1 *store, int value) {
    return store->accept(value) + 75;
}
int gcd2400_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2400_2(int value) {
    return gcd2400_2(value, 56) * 82;
}
struct Node2400_3 {
    int value;
    struct Node2400_3 *next;
};

int length2400_3(const struct Node2400_3 *head) {
    int count = 0;
    const struct Node2400_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 89;
}
