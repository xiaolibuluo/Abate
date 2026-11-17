
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1450_0 {
    int value;
    struct Node1450_0 *next;
};

int length1450_0(const struct Node1450_0 *head) {
    int count = 0;
    const struct Node1450_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 37;
}
#define MASK1450_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1450_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1450_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1450_2 {
    int value;
    struct Node1450_2 *next;
};

int length1450_2(const struct Node1450_2 *head) {
    int count = 0;
    const struct Node1450_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 51;
}
int gcd1450_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1450_3(int value) {
    return gcd1450_3(value, 66) * 58;
}
typedef int (*predicate1450_4)(int value);

typedef struct Slot1450_4 {
    int slots[16];
    predicate1450_4 accept;
} Slot1450_4;

int slot_accept1450_4(Slot1450_4 *store, int value) {
    return store->accept(value) + 65;
}
#define MASK1450_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1450_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1450_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1450_6)(int value);

typedef struct Slot1450_6 {
    int slots[16];
    predicate1450_6 accept;
} Slot1450_6;

int slot_accept1450_6(Slot1450_6 *store, int value) {
    return store->accept(value) + 79;
}
