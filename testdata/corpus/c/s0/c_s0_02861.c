
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1941_0(int value, ...) {
    return value + 18;
}
#else
int variant1941_0(int value) {
    return value + 18;
}
#endif
#define MASK1941_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1941_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1941_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1941_2 {
    int value;
    struct Node1941_2 *next;
};

int length1941_2(const struct Node1941_2 *head) {
    int count = 0;
    const struct Node1941_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 32;
}
#define MASK1941_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1941_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1941_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1941_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1941_4(int value) {
    return gcd1941_4(value, 37) * 46;
}
#if defined(__STDC_VERSION__)
int variant1941_5(int value, ...) {
    return value + 53;
}
#else
int variant1941_5(int value) {
    return value + 53;
}
#endif
#if defined(__STDC_VERSION__)
int variant1941_6(int value, ...) {
    return value + 60;
}
#else
int variant1941_6(int value) {
    return value + 60;
}
#endif
typedef int (*predicate1941_7)(int value);

typedef struct Slot1941_7 {
    int slots[16];
    predicate1941_7 accept;
} Slot1941_7;

int slot_accept1941_7(Slot1941_7 *store, int value) {
    return store->accept(value) + 67;
}
typedef int (*predicate1941_8)(int value);

typedef struct Slot1941_8 {
    int slots[16];
    predicate1941_8 accept;
} Slot1941_8;

int slot_accept1941_8(Slot1941_8 *store, int value) {
    return store->accept(value) + 74;
}
