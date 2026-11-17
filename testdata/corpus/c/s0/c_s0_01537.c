
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1042_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1042_0(int value) {
    return gcd1042_0(value, 44) * 68;
}
#define MASK1042_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1042_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1042_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1042_2)(int value);

typedef struct Slot1042_2 {
    int slots[16];
    predicate1042_2 accept;
} Slot1042_2;

int slot_accept1042_2(Slot1042_2 *store, int value) {
    return store->accept(value) + 82;
}
struct Node1042_3 {
    int value;
    struct Node1042_3 *next;
};

int length1042_3(const struct Node1042_3 *head) {
    int count = 0;
    const struct Node1042_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 89;
}
