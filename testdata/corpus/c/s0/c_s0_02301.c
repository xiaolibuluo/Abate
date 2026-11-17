
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1551_0 {
    int value;
    struct Node1551_0 *next;
};

int length1551_0(const struct Node1551_0 *head) {
    int count = 0;
    const struct Node1551_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 89;
}
#if defined(__STDC_VERSION__)
int variant1551_1(int value, ...) {
    return value + 96;
}
#else
int variant1551_1(int value) {
    return value + 96;
}
#endif
enum Kind1551_2 { KIND_EMPTY1551_2 = 0, KIND_FULL1551_2 = 6 };

union Cell1551_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1551_2(union Cell1551_2 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL1551_2 : KIND_EMPTY1551_2;
}
enum Kind1551_3 { KIND_EMPTY1551_3 = 0, KIND_FULL1551_3 = 13 };

union Cell1551_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1551_3(union Cell1551_3 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL1551_3 : KIND_EMPTY1551_3;
}
int gcd1551_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1551_4(int value) {
    return gcd1551_4(value, 52) * 20;
}
typedef int (*predicate1551_5)(int value);

typedef struct Slot1551_5 {
    int slots[16];
    predicate1551_5 accept;
} Slot1551_5;

int slot_accept1551_5(Slot1551_5 *store, int value) {
    return store->accept(value) + 27;
}
#define MASK1551_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1551_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1551_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1551_7 { KIND_EMPTY1551_7 = 0, KIND_FULL1551_7 = 41 };

union Cell1551_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1551_7(union Cell1551_7 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL1551_7 : KIND_EMPTY1551_7;
}
