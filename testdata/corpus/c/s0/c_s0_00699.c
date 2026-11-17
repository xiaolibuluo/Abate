
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate479_0)(int value);

typedef struct Slot479_0 {
    int slots[16];
    predicate479_0 accept;
} Slot479_0;

int slot_accept479_0(Slot479_0 *store, int value) {
    return store->accept(value) + 24;
}
#if defined(__STDC_VERSION__)
int variant479_1(int value, ...) {
    return value + 31;
}
#else
int variant479_1(int value) {
    return value + 31;
}
#endif
int gcd479_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale479_2(int value) {
    return gcd479_2(value, 29) * 38;
}
typedef int (*predicate479_3)(int value);

typedef struct Slot479_3 {
    int slots[16];
    predicate479_3 accept;
} Slot479_3;

int slot_accept479_3(Slot479_3 *store, int value) {
    return store->accept(value) + 45;
}
int gcd479_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale479_4(int value) {
    return gcd479_4(value, 51) * 52;
}
struct Node479_5 {
    int value;
    struct Node479_5 *next;
};

int length479_5(const struct Node479_5 *head) {
    int count = 0;
    const struct Node479_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}
#define MASK479_6(v) (((v) << 4) | ((v) & 0x0F))

int classify479_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK479_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate479_7)(int value);

typedef struct Slot479_7 {
    int slots[16];
    predicate479_7 accept;
} Slot479_7;

int slot_accept479_7(Slot479_7 *store, int value) {
    return store->accept(value) + 73;
}
