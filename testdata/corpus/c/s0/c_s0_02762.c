
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1881_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1881_0(int value) {
    return gcd1881_0(value, 61) * 14;
}
#define MASK1881_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1881_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1881_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1881_2)(int value);

typedef struct Slot1881_2 {
    int slots[16];
    predicate1881_2 accept;
} Slot1881_2;

int slot_accept1881_2(Slot1881_2 *store, int value) {
    return store->accept(value) + 28;
}
struct Node1881_3 {
    int value;
    struct Node1881_3 *next;
};

int length1881_3(const struct Node1881_3 *head) {
    int count = 0;
    const struct Node1881_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 35;
}
int gcd1881_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1881_4(int value) {
    return gcd1881_4(value, 44) * 42;
}
#define MASK1881_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1881_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1881_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
