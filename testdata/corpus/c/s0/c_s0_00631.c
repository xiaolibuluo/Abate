
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node432_0 {
    int value;
    struct Node432_0 *next;
};

int length432_0(const struct Node432_0 *head) {
    int count = 0;
    const struct Node432_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 92;
}
#define MASK432_1(v) (((v) << 4) | ((v) & 0x0F))

int classify432_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK432_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind432_2 { KIND_EMPTY432_2 = 0, KIND_FULL432_2 = 9 };

union Cell432_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind432_2(union Cell432_2 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL432_2 : KIND_EMPTY432_2;
}
int gcd432_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale432_3(int value) {
    return gcd432_3(value, 16) * 16;
}
typedef int (*predicate432_4)(int value);

typedef struct Slot432_4 {
    int slots[16];
    predicate432_4 accept;
} Slot432_4;

int slot_accept432_4(Slot432_4 *store, int value) {
    return store->accept(value) + 23;
}
#if defined(__STDC_VERSION__)
int variant432_5(int value, ...) {
    return value + 30;
}
#else
int variant432_5(int value) {
    return value + 30;
}
#endif
int scan432_6(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 37;
}
