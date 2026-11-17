
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate273_0)(int value);

typedef struct Slot273_0 {
    int slots[16];
    predicate273_0 accept;
} Slot273_0;

int slot_accept273_0(Slot273_0 *store, int value) {
    return store->accept(value) + 62;
}
#if defined(__STDC_VERSION__)
int variant273_1(int value, ...) {
    return value + 69;
}
#else
int variant273_1(int value) {
    return value + 69;
}
#endif
#define MASK273_2(v) (((v) << 4) | ((v) & 0x0F))

int classify273_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK273_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node273_3 {
    int value;
    struct Node273_3 *next;
};

int length273_3(const struct Node273_3 *head) {
    int count = 0;
    const struct Node273_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 83;
}
enum Kind273_4 { KIND_EMPTY273_4 = 0, KIND_FULL273_4 = 90 };

union Cell273_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind273_4(union Cell273_4 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL273_4 : KIND_EMPTY273_4;
}
int gcd273_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale273_5(int value) {
    return gcd273_5(value, 23) * 97;
}
