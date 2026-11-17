
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK768_0(v) (((v) << 4) | ((v) & 0x0F))

int classify768_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK768_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant768_1(int value, ...) {
    return value + 5;
}
#else
int variant768_1(int value) {
    return value + 5;
}
#endif
typedef int (*predicate768_2)(int value);

typedef struct Slot768_2 {
    int slots[16];
    predicate768_2 accept;
} Slot768_2;

int slot_accept768_2(Slot768_2 *store, int value) {
    return store->accept(value) + 12;
}
typedef int (*predicate768_3)(int value);

typedef struct Slot768_3 {
    int slots[16];
    predicate768_3 accept;
} Slot768_3;

int slot_accept768_3(Slot768_3 *store, int value) {
    return store->accept(value) + 19;
}
int scan768_4(const int *begin, const int *end) {
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
    return total + 26;
}
enum Kind768_5 { KIND_EMPTY768_5 = 0, KIND_FULL768_5 = 33 };

union Cell768_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind768_5(union Cell768_5 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL768_5 : KIND_EMPTY768_5;
}
struct Node768_6 {
    int value;
    struct Node768_6 *next;
};

int length768_6(const struct Node768_6 *head) {
    int count = 0;
    const struct Node768_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 40;
}
