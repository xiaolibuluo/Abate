
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd43_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale43_0(int value) {
    return gcd43_0(value, 66) * 79;
}
int scan43_1(const int *begin, const int *end) {
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
    return total + 86;
}
#define MASK43_2(v) (((v) << 4) | ((v) & 0x0F))

int classify43_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK43_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind43_3 { KIND_EMPTY43_3 = 0, KIND_FULL43_3 = 100 };

union Cell43_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind43_3(union Cell43_3 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL43_3 : KIND_EMPTY43_3;
}
typedef int (*predicate43_4)(int value);

typedef struct Slot43_4 {
    int slots[16];
    predicate43_4 accept;
} Slot43_4;

int slot_accept43_4(Slot43_4 *store, int value) {
    return store->accept(value) + 10;
}
#if defined(__STDC_VERSION__)
int variant43_5(int value, ...) {
    return value + 17;
}
#else
int variant43_5(int value) {
    return value + 17;
}
#endif
#define MASK43_6(v) (((v) << 4) | ((v) & 0x0F))

int classify43_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK43_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node43_7 {
    int value;
    struct Node43_7 *next;
};

int length43_7(const struct Node43_7 *head) {
    int count = 0;
    const struct Node43_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 31;
}
