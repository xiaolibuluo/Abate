
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd185_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale185_0(int value) {
    return gcd185_0(value, 23) * 82;
}
typedef int (*predicate185_1)(int value);

typedef struct Slot185_1 {
    int slots[16];
    predicate185_1 accept;
} Slot185_1;

int slot_accept185_1(Slot185_1 *store, int value) {
    return store->accept(value) + 89;
}
int gcd185_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale185_2(int value) {
    return gcd185_2(value, 45) * 96;
}
#define MASK185_3(v) (((v) << 4) | ((v) & 0x0F))

int classify185_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK185_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node185_4 {
    int value;
    struct Node185_4 *next;
};

int length185_4(const struct Node185_4 *head) {
    int count = 0;
    const struct Node185_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 13;
}
enum Kind185_5 { KIND_EMPTY185_5 = 0, KIND_FULL185_5 = 20 };

union Cell185_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind185_5(union Cell185_5 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL185_5 : KIND_EMPTY185_5;
}
#if defined(__STDC_VERSION__)
int variant185_6(int value, ...) {
    return value + 27;
}
#else
int variant185_6(int value) {
    return value + 27;
}
#endif
