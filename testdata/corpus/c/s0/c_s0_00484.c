
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate333_0)(int value);

typedef struct Slot333_0 {
    int slots[16];
    predicate333_0 accept;
} Slot333_0;

int slot_accept333_0(Slot333_0 *store, int value) {
    return store->accept(value) + 66;
}
enum Kind333_1 { KIND_EMPTY333_1 = 0, KIND_FULL333_1 = 73 };

union Cell333_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind333_1(union Cell333_1 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL333_1 : KIND_EMPTY333_1;
}
#if defined(__STDC_VERSION__)
int variant333_2(int value, ...) {
    return value + 80;
}
#else
int variant333_2(int value) {
    return value + 80;
}
#endif
int scan333_3(const int *begin, const int *end) {
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
    return total + 87;
}
struct Node333_4 {
    int value;
    struct Node333_4 *next;
};

int length333_4(const struct Node333_4 *head) {
    int count = 0;
    const struct Node333_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 94;
}
#define MASK333_5(v) (((v) << 4) | ((v) & 0x0F))

int classify333_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK333_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan333_6(const int *begin, const int *end) {
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
    return total + 11;
}
