
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2271_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2271_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2271_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2271_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2271_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2271_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2271_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2271_2(int value) {
    return gcd2271_2(value, 68) * 54;
}
typedef int (*predicate2271_3)(int value);

typedef struct Slot2271_3 {
    int slots[16];
    predicate2271_3 accept;
} Slot2271_3;

int slot_accept2271_3(Slot2271_3 *store, int value) {
    return store->accept(value) + 61;
}
#define MASK2271_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2271_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2271_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2271_5 {
    int value;
    struct Node2271_5 *next;
};

int length2271_5(const struct Node2271_5 *head) {
    int count = 0;
    const struct Node2271_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 75;
}
