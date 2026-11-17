
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK622_0(v) (((v) << 4) | ((v) & 0x0F))

int classify622_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK622_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node622_1 {
    int value;
    struct Node622_1 *next;
};

int length622_1(const struct Node622_1 *head) {
    int count = 0;
    const struct Node622_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 47;
}
#define MASK622_2(v) (((v) << 4) | ((v) & 0x0F))

int classify622_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK622_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate622_3)(int value);

typedef struct Slot622_3 {
    int slots[16];
    predicate622_3 accept;
} Slot622_3;

int slot_accept622_3(Slot622_3 *store, int value) {
    return store->accept(value) + 61;
}
#if defined(__STDC_VERSION__)
int variant622_4(int value, ...) {
    return value + 68;
}
#else
int variant622_4(int value) {
    return value + 68;
}
#endif
#if defined(__STDC_VERSION__)
int variant622_5(int value, ...) {
    return value + 75;
}
#else
int variant622_5(int value) {
    return value + 75;
}
#endif
int gcd622_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale622_6(int value) {
    return gcd622_6(value, 37) * 82;
}
int gcd622_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale622_7(int value) {
    return gcd622_7(value, 48) * 89;
}
