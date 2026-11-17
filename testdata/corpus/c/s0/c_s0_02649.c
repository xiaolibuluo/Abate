
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1798_0(int value, ...) {
    return value + 99;
}
#else
int variant1798_0(int value) {
    return value + 99;
}
#endif
struct Node1798_1 {
    int value;
    struct Node1798_1 *next;
};

int length1798_1(const struct Node1798_1 *head) {
    int count = 0;
    const struct Node1798_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 9;
}
#define MASK1798_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1798_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1798_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1798_3)(int value);

typedef struct Slot1798_3 {
    int slots[16];
    predicate1798_3 accept;
} Slot1798_3;

int slot_accept1798_3(Slot1798_3 *store, int value) {
    return store->accept(value) + 23;
}
#if defined(__STDC_VERSION__)
int variant1798_4(int value, ...) {
    return value + 30;
}
#else
int variant1798_4(int value) {
    return value + 30;
}
#endif
enum Kind1798_5 { KIND_EMPTY1798_5 = 0, KIND_FULL1798_5 = 37 };

union Cell1798_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1798_5(union Cell1798_5 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL1798_5 : KIND_EMPTY1798_5;
}
int gcd1798_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1798_6(int value) {
    return gcd1798_6(value, 34) * 44;
}
#if defined(__STDC_VERSION__)
int variant1798_7(int value, ...) {
    return value + 51;
}
#else
int variant1798_7(int value) {
    return value + 51;
}
#endif
#define MASK1798_8(v) (((v) << 4) | ((v) & 0x0F))

int classify1798_8(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1798_8(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
