
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1891_0(int value, ...) {
    return value + 47;
}
#else
int variant1891_0(int value) {
    return value + 47;
}
#endif
struct Node1891_1 {
    int value;
    struct Node1891_1 *next;
};

int length1891_1(const struct Node1891_1 *head) {
    int count = 0;
    const struct Node1891_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 54;
}
typedef int (*predicate1891_2)(int value);

typedef struct Slot1891_2 {
    int slots[16];
    predicate1891_2 accept;
} Slot1891_2;

int slot_accept1891_2(Slot1891_2 *store, int value) {
    return store->accept(value) + 61;
}
#define MASK1891_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1891_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1891_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1891_4 {
    int value;
    struct Node1891_4 *next;
};

int length1891_4(const struct Node1891_4 *head) {
    int count = 0;
    const struct Node1891_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 75;
}
int gcd1891_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1891_5(int value) {
    return gcd1891_5(value, 64) * 82;
}
#if defined(__STDC_VERSION__)
int variant1891_6(int value, ...) {
    return value + 89;
}
#else
int variant1891_6(int value) {
    return value + 89;
}
#endif
enum Kind1891_7 { KIND_EMPTY1891_7 = 0, KIND_FULL1891_7 = 96 };

union Cell1891_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1891_7(union Cell1891_7 *cell) {
    cell->number = 25;
    return cell->number > 0 ? KIND_FULL1891_7 : KIND_EMPTY1891_7;
}
