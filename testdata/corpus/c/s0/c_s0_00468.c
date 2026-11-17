
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd320_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale320_0(int value) {
    return gcd320_0(value, 53) * 91;
}
#define MASK320_1(v) (((v) << 4) | ((v) & 0x0F))

int classify320_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK320_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node320_2 {
    int value;
    struct Node320_2 *next;
};

int length320_2(const struct Node320_2 *head) {
    int count = 0;
    const struct Node320_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 8;
}
#if defined(__STDC_VERSION__)
int variant320_3(int value, ...) {
    return value + 15;
}
#else
int variant320_3(int value) {
    return value + 15;
}
#endif
typedef int (*predicate320_4)(int value);

typedef struct Slot320_4 {
    int slots[16];
    predicate320_4 accept;
} Slot320_4;

int slot_accept320_4(Slot320_4 *store, int value) {
    return store->accept(value) + 22;
}
typedef int (*predicate320_5)(int value);

typedef struct Slot320_5 {
    int slots[16];
    predicate320_5 accept;
} Slot320_5;

int slot_accept320_5(Slot320_5 *store, int value) {
    return store->accept(value) + 29;
}
