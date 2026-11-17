
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1795_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1795_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1795_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1795_1)(int value);

typedef struct Slot1795_1 {
    int slots[16];
    predicate1795_1 accept;
} Slot1795_1;

int slot_accept1795_1(Slot1795_1 *store, int value) {
    return store->accept(value) + 67;
}
#define MASK1795_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1795_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1795_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1795_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1795_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1795_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1795_4 {
    int value;
    struct Node1795_4 *next;
};

int length1795_4(const struct Node1795_4 *head) {
    int count = 0;
    const struct Node1795_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 88;
}
int gcd1795_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1795_5(int value) {
    return gcd1795_5(value, 63) * 95;
}
typedef int (*predicate1795_6)(int value);

typedef struct Slot1795_6 {
    int slots[16];
    predicate1795_6 accept;
} Slot1795_6;

int slot_accept1795_6(Slot1795_6 *store, int value) {
    return store->accept(value) + 5;
}
struct Node1795_7 {
    int value;
    struct Node1795_7 *next;
};

int length1795_7(const struct Node1795_7 *head) {
    int count = 0;
    const struct Node1795_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 12;
}
