
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK872_0(v) (((v) << 4) | ((v) & 0x0F))

int classify872_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK872_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK872_1(v) (((v) << 4) | ((v) & 0x0F))

int classify872_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK872_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK872_2(v) (((v) << 4) | ((v) & 0x0F))

int classify872_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK872_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate872_3)(int value);

typedef struct Slot872_3 {
    int slots[16];
    predicate872_3 accept;
} Slot872_3;

int slot_accept872_3(Slot872_3 *store, int value) {
    return store->accept(value) + 13;
}
int gcd872_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale872_4(int value) {
    return gcd872_4(value, 57) * 20;
}
#if defined(__STDC_VERSION__)
int variant872_5(int value, ...) {
    return value + 27;
}
#else
int variant872_5(int value) {
    return value + 27;
}
#endif
int gcd872_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale872_6(int value) {
    return gcd872_6(value, 18) * 34;
}
typedef int (*predicate872_7)(int value);

typedef struct Slot872_7 {
    int slots[16];
    predicate872_7 accept;
} Slot872_7;

int slot_accept872_7(Slot872_7 *store, int value) {
    return store->accept(value) + 41;
}
struct Node872_8 {
    int value;
    struct Node872_8 *next;
};

int length872_8(const struct Node872_8 *head) {
    int count = 0;
    const struct Node872_8 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 48;
}
