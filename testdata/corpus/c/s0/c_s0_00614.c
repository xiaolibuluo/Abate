
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK420_0(v) (((v) << 4) | ((v) & 0x0F))

int classify420_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK420_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd420_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale420_1(int value) {
    return gcd420_1(value, 32) * 40;
}
#define MASK420_2(v) (((v) << 4) | ((v) & 0x0F))

int classify420_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK420_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate420_3)(int value);

typedef struct Slot420_3 {
    int slots[16];
    predicate420_3 accept;
} Slot420_3;

int slot_accept420_3(Slot420_3 *store, int value) {
    return store->accept(value) + 54;
}
struct Node420_4 {
    int value;
    struct Node420_4 *next;
};

int length420_4(const struct Node420_4 *head) {
    int count = 0;
    const struct Node420_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 61;
}
int gcd420_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale420_5(int value) {
    return gcd420_5(value, 15) * 68;
}
