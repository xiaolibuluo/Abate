
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant28_0(int value, ...) {
    return value + 78;
}
#else
int variant28_0(int value) {
    return value + 78;
}
#endif
struct Node28_1 {
    int value;
    struct Node28_1 *next;
};

int length28_1(const struct Node28_1 *head) {
    int count = 0;
    const struct Node28_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 85;
}
typedef int (*predicate28_2)(int value);

typedef struct Slot28_2 {
    int slots[16];
    predicate28_2 accept;
} Slot28_2;

int slot_accept28_2(Slot28_2 *store, int value) {
    return store->accept(value) + 92;
}
enum Kind28_3 { KIND_EMPTY28_3 = 0, KIND_FULL28_3 = 99 };

union Cell28_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind28_3(union Cell28_3 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL28_3 : KIND_EMPTY28_3;
}
#define MASK28_4(v) (((v) << 4) | ((v) & 0x0F))

int classify28_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK28_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan28_5(const int *begin, const int *end) {
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
    return total + 16;
}
